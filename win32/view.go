package win32

import (
	"fmt"
	"github.com/Mengdch/win"
	"golang.org/x/sys/windows"
	"net/url"
	"os/exec"
	"strings"
	"unsafe"
)

type BlinkView struct {
	mWnd          win.HWND
	handle        wkeHandle
	proc          uintptr
	fnMap         map[int32]func(string) string
	mDC           win.HDC
	width, height int32
	pixels        unsafe.Pointer
	mBitmap       win.HBITMAP
	url           string
	inJs          string
}

func (v *BlinkView) createBitmap() {
	var bi win.BITMAPINFOHEADER
	bi.BiSize = 40 // (win.BITMAPINFOHEADER)
	bi.BiWidth = v.width
	bi.BiHeight = v.height
	bi.BiPlanes = 1
	bi.BiBitCount = 32
	bi.BiCompression = win.BI_RGB

	hBmp := win.CreateDIBSection(0, &bi, win.DIB_RGB_COLORS, &v.pixels, 0, 0)
	win.SelectObject(v.mDC, win.HGDIOBJ(hBmp))
	if v.mBitmap != 0 {
		win.DeleteObject(win.HGDIOBJ(v.mBitmap))
	}

	v.mBitmap = hBmp
}

func (v *BlinkView) init(ua, dev string, jsFunc map[int32]func(string) string) bool {
	v.fnMap = jsFunc
	if mbHandle != nil {
		v.handle = mbHandle.wkeCreateWebView()
		mbHandle.wkeSetTransparent(v.handle, false)
		mbHandle.wkeOnJsQuery(v.handle, v.onJsQuery, 0)
		mbHandle.wkeSetNavigationToNewWindowEnable(v.handle, true)
		mbHandle.wkeOnAlertBox(v.handle, v.onAlert, 0)
		mbHandle.wkeOnPaintUpdated(v.handle, v.paintUpdatedCallback, 0)
		// mbHandle.wkeOnLoadUrlEnd(v.handle, v.wkeLoadUrlEndCallback, 0)
		mbHandle.wkeOnDocumentReady(v.handle, v.wkeOnDocumentReady, 0)
		if len(ua) > 0 {
			mbHandle.wkeSetUserAgent(v.handle, ua)
		}
		if len(dev) > 0 {
			dev = "file:///" + strings.ReplaceAll(dev, "\\", "/") + "/inspector.html"
			mbHandle.wkeSetDebugConfig(v.handle, showDevTools, dev)
		}
		return true
	}
	return false
}

func (v *BlinkView) setHWnd(parent win.HWND) {
	v.mWnd = parent
	v.mDC = win.CreateCompatibleDC(0)
	mbHandle.wkeSetHandle(v.handle, uintptr(v.mWnd))
}
func (v *BlinkView) close() {
	if v.mDC != 0 {
		win.DeleteDC(v.mDC)
	}
	if v.mBitmap != 0 {
		win.DeleteObject(win.HGDIOBJ(v.mBitmap))
	}
	mbHandle.wkeOnPaintUpdated(v.handle, nil, uintptr(v.mWnd))
	mbHandle.wkeSetHandle(v.handle, 0)
	mbHandle.wkeDestroyWebView(v.handle)
}
func (v *BlinkView) setDownloadCallback(callback func(wke wkeHandle, param uintptr, length uint32, url, mime, disposition uintptr, job wkeNetJob, dataBind uintptr) wkeDownloadOpt) {
	mbHandle.wkeOnDownload(v.handle, callback, 0)
	return
}
func (v *BlinkView) wkePopupDialogAndDownload(param uintptr, contentLength uint32, url, mime,
	disposition uintptr, job wkeNetJob, data uintptr, callback *wkeDownloadBind) wkeDownloadOpt {
	r, _, _ := mbHandle._wkePopupDialogAndDownload.Call(uintptr(v.handle), param, uintptr(contentLength), url, mime,
		disposition, uintptr(job), data, uintptr(unsafe.Pointer(callback)))
	return wkeDownloadOpt(r)
}
func (v *BlinkView) wkeOnDocumentReady(wke wkeHandle, param uintptr, frame wkeFrame) uintptr {
	v.runJs(frame)
	return 0
}

func (v *BlinkView) runJs(frame wkeFrame) {
	if len(v.inJs) == 0 {
		return
	}
	fmt.Println("run script")
	mbHandle.wkeRunJs(v.handle, frame, strToCharPtr(v.inJs), false, nil, 0, 0)
	v.inJs = ""
	return
}
func (v *BlinkView) wkeLoadingFinishCallback(wke wkeHandle, param uintptr, frame wkeFrame, url uintptr, result wkeLoadingResult, reason uintptr) uintptr {
	uri := ptrToUtf8(url)
	fmt.Println("load finish", result, v.url, uri)
	v.runJs(frame)
	return 0
}
func (v *BlinkView) wkeLoadUrlEndCallback(wke wkeHandle, param, url uintptr, job wkeNetJob, buf uintptr, count int32) uintptr {
	frame := mbHandle.wkeWebFrameGetMainFrame(v.handle)
	v.runJs(frame)
	return 0
}
func (v *BlinkView) wkeLoadUrlBeginCallback(wke wkeHandle, param, utf8Url uintptr, job wkeNetJob) uintptr {
	uri := ptrToUtf8(utf8Url)
	if len(v.url) > 0 {
		if len(v.inJs) > 0 {
			frame := mbHandle.wkeWebFrameGetMainFrame(wke)
			v.runJs(frame)
		}
		go logRecord("loadUrlBegin:"+v.url, "")
		v.url = ""
	}
	return operateUri(uri)
}

func operateUri(uri string) uintptr {
	u, err := url.Parse(uri)
	if err != nil {
		go logRecord("operateUri.Parse:"+uri, err.Error())
		return 0
	}
	switch u.Scheme {
	case "http", "https", "ws", "wss":
		return 0
	default:
		if exist, err := checkProtocol(u.Scheme); exist {
			go exec.Command("start", uri).Run()
			return 1
		} else if err != nil {
			go logRecord("operateUri.checkProtocol:"+uri+"("+u.Scheme+"):", err.Error())
		}
	}
	return 0
}
func (v *BlinkView) consoleCallback(wke wkeHandle, param uintptr, level int32, msg, name, line, stack uintptr) uintptr {
	// v.runJs()
	fmt.Println("console")
	return 0
}
func (v *BlinkView) paintUpdatedCallback(wke wkeHandle, param, hdc uintptr, x, y, cx, cy int32) uintptr {
	if v.pixels == nil {
		if v.width != cx || v.height != cy {
			return 0
		}
		v.createBitmap()
	}

	hScreenDC := win.GetDC(v.mWnd)
	win.BitBlt(v.mDC, x, y, v.width, v.height, win.HDC(hdc), x, y, win.SRCCOPY)
	win.BitBlt(hScreenDC, x, y, v.width, v.height, v.mDC, x, y, win.SRCCOPY)
	win.ReleaseDC(v.mWnd, hScreenDC)
	return 0
}

func (v *BlinkView) LoadUrl(url string) {
	v.LoadUrlScript(url, "")
}
func (v *BlinkView) LoadUrlScript(url, script string) {
	v.url = url
	v.inJs = script
	mbHandle.wkeLoadURL(v.handle, url)
	mbHandle.wkeOnLoadUrlBegin(v.handle, v.wkeLoadUrlBeginCallback, 0) // 这里没找到为什么必须加载后
}
func (v *BlinkView) SetOnNewWindow(callback wkeOnCreateViewCallback) {
	mbHandle.wkeOnCreateView(v.handle, callback, 0)
}

func (v *BlinkView) OnWndProc(hWnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_ERASEBKGND:
		return 1
	case win.WM_SIZE:
		w, h := int32(win.LOWORD(uint32(lParam))), int32(win.HIWORD(uint32(lParam)))
		v.resize(w, h, true)
	case win.WM_SYSKEYDOWN, win.WM_KEYDOWN:
		if v.keyDown(msg, wParam, lParam, mbHandle.wkeFireKeyDownEvent) {
			return win.DefWindowProc(hWnd, msg, wParam, lParam)
		}
	case win.WM_SYSKEYUP, win.WM_KEYUP:
		if v.keyDown(msg, wParam, lParam, mbHandle.wkeFireKeyUpEvent) {
			return win.DefWindowProc(hWnd, msg, wParam, lParam)
		}
	case win.WM_CHAR:
		if v.keyDown(msg, wParam, lParam, mbHandle.wkeFireKeyPressEvent) {
			return 0
		}
	case win.WM_LBUTTONUP, win.WM_LBUTTONDOWN, win.WM_LBUTTONDBLCLK, win.WM_RBUTTONUP, win.WM_RBUTTONDOWN,
		win.WM_RBUTTONDBLCLK, win.WM_MBUTTONUP, win.WM_MBUTTONDOWN, win.WM_MBUTTONDBLCLK, win.WM_MOUSEMOVE:
		if v.mouse(hWnd, msg, wParam, lParam) {
			return 0
		}
	case win.WM_CONTEXTMENU:
		if v.menu(hWnd, wParam, lParam) {
			return 0
		}
	case win.WM_MOUSEWHEEL:
		if v.mouseWheel(hWnd, wParam, lParam) {
			return 0
		}
	case win.WM_SETFOCUS:
		mbHandle.wkeSetFocus(v.handle)
		return 0
	case win.WM_KILLFOCUS:
		mbHandle.wkeKillFocus(v.handle)
		return 0
	case win.WM_PAINT:
		var paintInfo win.PAINTSTRUCT
		win.BeginPaint(hWnd, &paintInfo)
		win.BitBlt(paintInfo.Hdc, 0, 0, v.width, v.height, v.mDC, 0, 0, win.SRCCOPY)
		win.EndPaint(hWnd, &paintInfo)
		return 0
	case win.WM_SETCURSOR, win.WM_IME_STARTCOMPOSITION:
		if mbHandle.wkeFireWindowsMessage(v.handle, hWnd, int32(msg), int32(0), int32(0)) {
			return 0
		}
	case win.WM_INPUTLANGCHANGE:
		return win.DefWindowProc(hWnd, msg, wParam, lParam)
	}
	return win.DefWindowProc(hWnd, msg, wParam, lParam)
}
func (v *BlinkView) menu(hWnd win.HWND, wParam, lParam uintptr) bool {
	pt := getPoint(hWnd, lParam)

	flags := getFlags(wParam)
	return mbHandle.wkeFireContextMenuEvent(v.handle, pt.X, pt.Y, flags)
}
func (v *BlinkView) mouseWheel(hWnd win.HWND, wParam, lParam uintptr) bool {
	pt := getPoint(hWnd, lParam)
	flags := getFlags(wParam)
	return mbHandle.wkeFireMouseWheelEvent(v.handle, pt.X, pt.Y, int32(win.HIWORD(uint32(wParam))), flags)
}

func getPoint(hWnd win.HWND, lParam uintptr) win.POINT {
	var pt win.POINT
	pt.X = int32(win.LOWORD(uint32(lParam)))
	pt.Y = int32(win.HIWORD(uint32(lParam)))
	win.ScreenToClient(hWnd, &pt)
	return pt
}
func (v *BlinkView) mouse(hWnd win.HWND, msg uint32, wParam, lParam uintptr) bool {
	if msg == win.WM_LBUTTONDOWN || msg == win.WM_MBUTTONDOWN || msg == win.WM_RBUTTONDOWN {
		if win.GetFocus() != hWnd {
			win.SetFocus(hWnd)
		}

		win.SetCapture(hWnd)
	} else if msg == win.WM_LBUTTONUP || msg == win.WM_MBUTTONUP || msg == win.WM_RBUTTONUP {
		win.ReleaseCapture()
	}

	x := win.LOWORD(uint32(lParam))
	y := win.HIWORD(uint32(lParam))

	flags := getFlags(wParam)
	return mbHandle.wkeFireMouseEvent(v.handle, int32(msg), int32(x), int32(y), flags)
}

func getFlags(wParam uintptr) int32 {
	var flags int32 = 0
	if (wParam & win.MK_CONTROL) != 0 {
		flags |= WKE_CONTROL
	}

	if (wParam & win.MK_SHIFT) != 0 {
		flags |= WKE_SHIFT
	}

	if (wParam & win.MK_LBUTTON) != 0 {
		flags |= WKE_LBUTTON
	}

	if (wParam & win.MK_MBUTTON) != 0 {
		flags |= WKE_MBUTTON
	}

	if (wParam & win.MK_RBUTTON) != 0 {
		flags |= WKE_RBUTTON
	}
	return flags
}
func (v *BlinkView) resize(w, h int32, set bool) {
	if v.handle > 0 {
		mbHandle.wkeResize(v.handle, uint32(w), uint32(h))
	}
	if !set {
		return
	}
	if v.width == w && v.height == h {
		return
	}

	v.width = w
	v.height = h
	v.pixels = nil
}
func (v *BlinkView) keyDown(msg uint32, wParam, lParam uintptr, fun func(wkeHandle, uint32, uint32, bool) bool) bool {
	var flags uint32 = 0
	lp := int32(lParam)
	if lp>>16&win.KF_REPEAT != 0 {
		flags |= WKE_REPEAT
	}
	if lp>>16&win.KF_EXTENDED != 0 {
		flags |= WKE_EXTENDED
	}
	isSys := false
	if msg == win.WM_SYSKEYDOWN {
		isSys = true
	}
	return fun(v.handle, uint32(wParam), flags, isSys)
}
func (v *BlinkView) onAlert(wke wkeHandle, param uintptr, msg uintptr) uintptr {
	content := ptrToUtf8(msg)
	win.MessageBox(v.mWnd, windows.StringToUTF16Ptr(content), windows.StringToUTF16Ptr("警告"), 0)
	return 0
}
