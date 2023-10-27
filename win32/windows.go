package win32

import (
	"errors"
	"fmt"
	thuOS "github.com/Mengdch/goUtil/OS"
	"github.com/Mengdch/win"
	"golang.org/x/sys/windows"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"unsafe"
)

const (
	className      = "thublink_class"
	windowName     = "thublink_window"
	classViewName  = "thublink_view_class"
	windowViewName = "thublink_view_window"
)

var (
	classNamePtr      *uint16
	windowNamePtr     *uint16
	classViewNamePtr  *uint16
	windowViewNamePtr *uint16
	hInst             win.HINSTANCE
	procMap           map[win.HWND]uintptr
	mbHandle          *Thublink
	iconHandle        win.HANDLE
	urls              []string
	main              *Window
)

type SaveCallback func(url, path string)
type FinishCallback func(url string, success bool)

func init() {
	mbHandle = new(Thublink).Init()
	if mbHandle == nil {
		return
	}
	mbHandle.wkeEnableHighDPISupport()
	var err error
	classNamePtr, err = syscall.UTF16PtrFromString(className)
	if err != nil {
		fmt.Println(err)
		return
	}
	windowNamePtr, err = syscall.UTF16PtrFromString(windowName)
	if err != nil {
		fmt.Println(err)
		return
	}
	classViewNamePtr, err = syscall.UTF16PtrFromString(classViewName)
	if err != nil {
		fmt.Println(err)
		return
	}
	windowViewNamePtr, err = syscall.UTF16PtrFromString(windowViewName)
	if err != nil {
		fmt.Println(err)
		return
	}
	hInst = win.GetModuleHandle(nil)
	wndClass := win.WNDCLASSEX{
		Style:         win.CS_HREDRAW | win.CS_VREDRAW,
		LpfnWndProc:   syscall.NewCallbackCDecl(classMsgProc),
		HInstance:     hInst,
		LpszClassName: classNamePtr,
		HCursor:       win.LoadCursor(0, win.MAKEINTRESOURCE(win.IDC_ARROW)),
		HbrBackground: win.GetSysColorBrush(win.COLOR_WINDOW + 1),
	}
	wndClass.CbSize = uint32(unsafe.Sizeof(wndClass))
	win.RegisterClassEx(&wndClass)
	wndClass = win.WNDCLASSEX{
		Style:         win.CS_DBLCLKS,
		LpfnWndProc:   syscall.NewCallbackCDecl(classMsgProc),
		HInstance:     hInst,
		LpszClassName: classViewNamePtr,
		HbrBackground: win.GetSysColorBrush(win.COLOR_WINDOW + 1),
		HCursor:       win.LoadCursor(0, win.MAKEINTRESOURCE(win.IDC_ARROW)),
	}
	wndClass.CbSize = uint32(unsafe.Sizeof(wndClass))
	win.RegisterClassEx(&wndClass)
	procMap = make(map[win.HWND]uintptr)
}
func classMsgProc(hWnd win.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	if v, e := procMap[hWnd]; e {
		return win.CallWindowProc(v, hWnd, msg, wParam, lParam)
	}
	return win.DefWindowProc(hWnd, msg, wParam, lParam)
}
func newWindow(exStyle, style uint32, pos int, parent win.HWND, width, height int32, proc func(hWnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr) win.HWND {
	return newClassWindow(exStyle, style, parent, pos, width, height, classNamePtr, windowNamePtr, proc)
}
func newClassWindow(exStyle, style uint32, parent win.HWND, pos int, width, height int32, className, windowName *uint16,
	proc func(hWnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr) win.HWND {
	var x, y, sw, sh int32
	noMax := parent == 0 && style&win.WS_MAXIMIZE == 0
	if noMax { // 居中
		sw, sh = getRect()
		x, y, width, height = getMiddlePos(width, height, sw, sh)
		switch pos {
		case thuOS.LeftTop:
			x, y = 0, 0
		case thuOS.RightBottom:
			x = sw - width
			y = sh - height
		}
	}
	wnd := win.CreateWindowEx(exStyle, className, windowName, style, x, y, width, height,
		parent, 0, hInst, unsafe.Pointer(nil))
	if wnd != 0 {
		if noMax {
			dpi := GetDPI(wnd)
			if dpi != 1.0 {
				width := int32(float64(width) * dpi)
				height := int32(float64(height) * dpi)
				x, y, width, height = getMiddlePos(width, height, sw, sh)
				win.SetWindowPos(wnd, win.HWND_NOTOPMOST, x, y, width, height, win.SWP_NOREDRAW)
			}
		}
		procMap[wnd] = syscall.NewCallbackCDecl(proc)
	}
	return wnd
}
func getRect() (int32, int32) {
	m := win.MonitorFromPoint(0, 0, win.MONITOR_DEFAULTTOPRIMARY)
	var sw, sh int32
	if m > 0 {
		var i win.MONITORINFO
		i.CbSize = uint32(unsafe.Sizeof(win.MONITORINFO{}))
		win.GetMonitorInfo(m, &i)
		sw = i.RcWork.Width()
		sh = i.RcWork.Height()
	} else {
		sw = win.GetSystemMetrics(win.SM_CXFULLSCREEN)
		sh = win.GetSystemMetrics(win.SM_CYFULLSCREEN)
	}
	return sw, sh
}
func getMiddlePos(width, height, sw, sh int32) (int32, int32, int32, int32) {
	var top, left int32
	if width < sw {
		left = (sw - width) / 2
	}
	if height < sh {
		top = (sh - height) / 2
	}
	return left, top, width, height
}

type FormProfile struct {
	Title      string
	UserAgent  string
	Width      int
	Height     int
	Max        bool
	Mb         bool
	Ib         bool
	index      string
	devPath    string
	jsFunction map[int32]func(string) string
	subs       map[string]FormProfile
	main       bool
	save       SaveCallback
	finish     FinishCallback
	domain     []string
	background bool
	size       bool
	close      func(uintptr)
	pos        int
}

func ShowMainWindow(url, script string, x, y int32) {
	if len(url) == 0 {
		return
	}
	if main != nil {
		main.profile.index = url
		main.view.LoadUrlScript(url, script)
		win.SetWindowPos(main.hWnd, win.HWND_TOPMOST, x, y, 0, 0, win.SWP_NOSIZE|win.SWP_NOREDRAW|win.SWP_NOACTIVATE|win.SWP_SHOWWINDOW)
	}
}
func StartBlinkMain(url, title, ico, ua, devPath string, max, mb, ib, show, size bool, width, height, pos int,
	jsFunc map[int32]func(string) string, forms map[string]FormProfile, set, close func(uintptr),
	s SaveCallback, f FinishCallback, domains []string) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	for i, v := range forms {
		if len(v.Title) == 0 {
			v.Title = title
		}
		if len(v.UserAgent) == 0 {
			v.UserAgent = ua
		}
		if len(devPath) > 0 {
			v.devPath = devPath
		}
		if v.Width <= 0 {
			v.Width = width
		}
		if v.Height <= 0 {
			v.Height = height
		}
		if v.jsFunction == nil {
			v.jsFunction = jsFunc
		}
		v.finish = f
		v.save = s
		v.index = i
		v.domain = domains
		forms[i] = v
	}
	main := FormProfile{Title: title, UserAgent: ua, index: url, devPath: devPath, Max: max, Mb: mb, Ib: ib, background: !show, main: true,
		jsFunction: jsFunc, subs: forms, Width: width, Height: height, close: close, finish: f, save: s, domain: domains, size: size, pos: pos}
	loadIcon(ico)
	if !main.newBlinkWindow(set) {
		return errors.New("not start")
	}
	// 3. 主消息循环
	MainLoop()
	return nil
}

func MainLoop() {
	msg := (*win.MSG)(unsafe.Pointer(win.GlobalAlloc(0, unsafe.Sizeof(win.MSG{}))))
	defer win.GlobalFree(win.HGLOBAL(unsafe.Pointer(msg)))
	for win.GetMessage(msg, 0, 0, 0) > 0 {
		// fmt.Println(msg.Message, msg.HWnd, msg.LParam, msg.WParam)
		if msg.Message == win.WM_QUIT {
			mbHandle.wkeUnInit()
			break
		}
		win.TranslateMessage(msg)
		win.DispatchMessage(msg)
	}
}

func (fp FormProfile) noTitle() bool {
	if len(fp.index) == 0 { // TODO 临时
		return true
	}
	for i := range fp.domain {
		if strings.Contains(fp.index, fp.domain[i]) {
			return true
		}
	}
	return false
}

func (fp FormProfile) newBlinkWindow(set func(uintptr)) bool {
	w := window{profile: fp}
	w.init()
	if set != nil {
		set(uintptr(w.hWnd))
	}
	v := BlinkView{}
	var r win.RECT
	win.GetClientRect(w.hWnd, &r)
	if fp.main {
		main = &w
	}
	if !v.init(fp.UserAgent, fp.devPath, fp.jsFunction) {
		return false
	}
	v.SetOnNewWindow(w.onCreateView)
	v.SetDownloadCallback(w.wkeOnDownloadCallback)
	w.child = newClassWindow(0, win.WS_CHILD|win.WS_VISIBLE|win.WS_CLIPSIBLINGS|win.WS_CLIPCHILDREN, w.hWnd, thuOS.Center, r.Width(), r.Height(), classViewNamePtr, windowViewNamePtr, v.OnWndProc)
	v.setHWnd(w.child)
	v.resize(r.Width(), r.Height(), true)
	v.LoadUrl(fp.index)
	mbHandle.wkeOnLoadUrlBegin(v.handle, v.wkeLoadUrlBeginCallback, 0)
	urls = append(urls, fp.index)
	w.view = &v
	return true
}

func loadIcon(ico string) {
	if len(ico) == 0 {
		return
	}
	if !fileFunc.CheckFileExist(ico) {
		ico = filepath.Join(fileFunc.NowPath(), ico)
	}
	if fileFunc.CheckFileExist(ico) {
		hInst := win.GetModuleHandle(nil)
		fromString, err := syscall.UTF16PtrFromString(ico)
		if err != nil {
			fmt.Println(err)
			return
		}
		iconHandle = win.LoadImage(hInst, fromString, win.IMAGE_ICON, 0, 0, win.LR_LOADFROMFILE)
	}
}

type window struct {
	hWnd    win.HWND
	child   win.HWND
	profile FormProfile
	view    *BlinkView
	down    map[string]*downInfo
	bind    map[string]*wkeDownloadBind
	mux     sync.Mutex
}

func (w *window) init() {
	w.down = make(map[string]*downInfo)
	w.bind = make(map[string]*wkeDownloadBind)
	w.hWnd = newWindow(0, w.style(), w.profile.pos, 0, int32(w.profile.Width), int32(w.profile.Height), w.windowMsgProc)
	if w.hWnd == 0 {
		return
	}
	if iconHandle != 0 {
		win.SendMessage(w.hWnd, win.WM_SETICON, 1, uintptr(iconHandle))
	}
	if len(w.profile.Title) > 0 {
		win.SetWindowText(w.hWnd, w.profile.Title)
	}
	//w.show()
}
func (w *window) show() {
	if win.IsWindowVisible(w.hWnd) {
		return
	}
	nCmdShow := win.SW_SHOW
	if w.profile.background {
		nCmdShow = win.SW_HIDE
	}
	win.ShowWindow(w.hWnd, int32(nCmdShow))
	win.UpdateWindow(w.hWnd)
	if !w.profile.background {
		win.SetForegroundWindow(w.hWnd)
		win.SetActiveWindow(w.hWnd)
		win.SetFocus(w.hWnd)

	}
}
func (w *Window) style() uint32 {
	var style uint32 = win.CS_DROPSHADOW
	if w.profile.noTitle() {
		style |= win.WS_POPUP | win.WS_THICKFRAME
	} else {
		style |= win.WS_OVERLAPPEDWINDOW
	}
	if !w.profile.Ib {
		style ^= win.WS_MINIMIZEBOX
	}
	if !w.profile.Mb {
		style ^= win.WS_MAXIMIZEBOX
	}
	if w.profile.Max {
		style |= win.WS_MAXIMIZE
	}
	return style
}
func (w *window) roundRect() { // 有效果，但是很丑，还有bug
	/*
		您可以创建一个没有任何框架的窗口，使用WS_EX_LAYERED获取透明度，然后在WM_PAINT中“正常”绘制包含自定义框架的窗口，或者组成一个离屏位图，并使用
		UpdateLayeredWindow
		（后一种方法更有效） 。 当然，您必须将绘制的内容调整为窗口的当前大小。通常，您可以从不同的元素组成 - 例如使用四个“角落”位图（或椭圆函数）绘制角点，然后绘制边框等。 此外，您可以处理
		WM_NCHITTEST
		将“标题”/“边框”/“角落”功能（即移动和调整窗口大小）分配到窗口的任意区域。
	*/
	r, e := win.GetWindowRect(w.hWnd)
	if e {
		rgn := win.CreateRoundRectRgn(r.Left, r.Top, r.Right, r.Bottom, 20, 20)
		win.SetWindowRgn(w.hWnd, rgn, true)
	}
}
func (w *window) windowMsgProc(hWnd win.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	switch msg {
	case win.WM_SIZE:
		if w.child > 0 && w.view != nil {
			var r win.RECT
			win.GetClientRect(hWnd, &r)
			win.MoveWindow(w.child, r.Left, r.Top, r.Width(), r.Height(), true)
			w.view.resize(r.Width(), r.Height(), false)
		}
	case win.WM_CLOSE:
		if w.view == nil {
			break
		}
		w.view.close()
		if w.profile.main {
			if w.profile.close != nil {
				w.profile.close(uintptr(w.hWnd))
			}
			win.PostQuitMessage(0)
		}
	}
	return win.DefWindowProc(hWnd, msg, wParam, lParam)
}
func (w *window) wkeOnDownloadCallback(wke wkeHandle, param uintptr, length uint32, url, mime, disposition uintptr, job wkeNetJob, dataBind uintptr) wkeDownloadOpt {
	info := downInfo{}
	urlStr := ptrToUtf8(url)
	info.url = strToCharPtr(urlStr)
	info.recvSize = 0
	info.allSize = length
	bind := wkeDownloadBind{param: uintptr(unsafe.Pointer(&info))}
	if w.profile.finish != nil {
		bind.finishCallback = syscall.NewCallback(func(param uintptr, job wkeNetJob, result wkeLoadingResult) uintptr {
			info := (*downInfo)(unsafe.Pointer(param))
			if info != nil {
				w.profile.finish(ptrToUtf8(info.url), result == WKE_LOADING_SUCCEEDED)
			}
			w.mux.Lock()
			defer w.mux.Unlock()
			delete(w.down, urlStr)
			delete(w.bind, urlStr)
			return 0
		})
	}
	if w.profile.save != nil {
		bind.saveNameCallback = syscall.NewCallback(func(ptr, filePath uintptr) uintptr {
			info := (*downInfo)(unsafe.Pointer(ptr))
			if info != nil {
				w.profile.save(ptrToUtf8(info.url), windows.UTF16PtrToString((*uint16)(unsafe.Pointer(filePath))))
			}
			return 0
		})
	}
	w.mux.Lock()
	defer w.mux.Unlock()
	w.down[urlStr] = &info
	w.bind[urlStr] = &bind
	return w.view.wkePopupDialogAndDownload(param, length, url, mime, disposition, job, dataBind, &bind)
}
func (w *window) onCreateView(wke wkeHandle, param uintptr, naviType wkeNavigationType, url, feature uintptr) uintptr {
	a := ptrToUtf8(url)
	if Debug() {
		fmt.Println("onCreateView", a)
	}
	urls = append(urls, a)
	if v, e := w.profile.subs[a]; e {
		v.newBlinkWindow(nil)
	} else {
		o := operateUri(a)
		if o == 1 {
			return 0
		}
		n := w.profile
		n.index = a
		n.main = false
		n.close = nil
		n.newBlinkWindow(nil)
	}
	return 0
}
func Debug() bool {
	return true
}
