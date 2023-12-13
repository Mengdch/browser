package win32

import (
	"fmt"
	"github.com/Mengdch/goUtil/FileTools"
	"github.com/Mengdch/win"
	"golang.org/x/sys/windows"
	"path/filepath"
	"syscall"
	"unsafe"
)

type Thublink struct {
	_dll *windows.LazyDLL

	_wkeInitialize                     *windows.LazyProc
	_wkeUnInitialize                   *windows.LazyProc
	_wkeCreateWebView                  *windows.LazyProc
	_wkeSetHandle                      *windows.LazyProc
	_wkeOnPaintBitUpdated              *windows.LazyProc
	_wkeOnPaintUpdated                 *windows.LazyProc
	_wkeLoadURL                        *windows.LazyProc
	_wkeGetHostHWND                    *windows.LazyProc
	_wkeResize                         *windows.LazyProc
	_wkeNetOnResponse                  *windows.LazyProc
	_wkeOnLoadUrlBegin                 *windows.LazyProc
	_wkeFireMouseEvent                 *windows.LazyProc
	_wkeFireContextMenuEvent           *windows.LazyProc
	_wkeFireWindowsMessage             *windows.LazyProc
	_wkeCreateWebWindow                *windows.LazyProc
	_wkeShowWindow                     *windows.LazyProc
	_wkeFireMouseWheelEvent            *windows.LazyProc
	_wkeFireKeyUpEvent                 *windows.LazyProc
	_wkeFireKeyDownEvent               *windows.LazyProc
	_wkeFireKeyPressEvent              *windows.LazyProc
	_wkeSetFocus                       *windows.LazyProc
	_wkeNetGetRequestMethod            *windows.LazyProc
	_wkeNetSetData                     *windows.LazyProc
	_wkeNetCancelRequest               *windows.LazyProc
	_wkeDestroyWebView                 *windows.LazyProc
	_jsGetWebView                      *windows.LazyProc
	_wkeKillFocus                      *windows.LazyProc
	_wkeOnDidCreateScriptContext       *windows.LazyProc
	_wkeIsMainFrame                    *windows.LazyProc
	_wkeGetString                      *windows.LazyProc
	_wkeNetSetHTTPHeaderField          *windows.LazyProc
	_wkeNetChangeRequestUrl            *windows.LazyProc
	_wkeNetHookRequest                 *windows.LazyProc
	_wkeNetHoldJobToAsynCommit         *windows.LazyProc
	_wkeNetContinueJob                 *windows.LazyProc
	_wkeOnLoadUrlEnd                   *windows.LazyProc
	_wkeOnConsole                      *windows.LazyProc
	_wkeOnLoadUrlFail                  *windows.LazyProc
	_wkeOnJsQuery                      *windows.LazyProc
	_wkeOnDocumentReady2               *windows.LazyProc
	_wkeOnDownload                     *windows.LazyProc
	_wkeOnAlertBox                     *windows.LazyProc
	_wkeOnCreateView                   *windows.LazyProc
	_wkeSetContextMenuEnabled          *windows.LazyProc
	_wkeResponseQuery                  *windows.LazyProc
	_wkeNetGetMIMEType                 *windows.LazyProc
	_wkeNetSetMIMEType                 *windows.LazyProc
	_wkeNetGetRawResponseHead          *windows.LazyProc
	_wkeSetTransparent                 *windows.LazyProc
	_wkeSetViewProxy                   *windows.LazyProc
	_wkeSetNavigationToNewWindowEnable *windows.LazyProc
	_wkeSetUserAgent                   *windows.LazyProc
	_wkeSetDebugConfig                 *windows.LazyProc
	_wkePopupDialogAndDownload         *windows.LazyProc
	_wkeGetLockedViewDC                *windows.LazyProc
	_wkeRunMessageLoop                 *windows.LazyProc
	_wkeWebFrameGetMainFrame           *windows.LazyProc
	_wkeRunJs                          *windows.LazyProc
	_wkeOnLoadingFinish                *windows.LazyProc
	_wkeEnableHighDPISupport           *windows.LazyProc
	_wkeOnTitleChanged                 *windows.LazyProc
	_wkeNetGetFavicon                  *windows.LazyProc
	_wkeReload                         *windows.LazyProc
	_wkeGetUrl                         *windows.LazyProc
	_wkeStopLoading                    *windows.LazyProc
	_wkeGoBack                         *windows.LazyProc
	_wkeGoForward                      *windows.LazyProc
	_wkeCanGoForward                   *windows.LazyProc
}

const (
	dll_name = "thublink"
)

func (t *Thublink) Init() *Thublink {
	lib := windows.NewLazyDLL(filepath.Join(fileFunc.NowPath(), getName()))
	if !fileFunc.CheckFileExist(lib.Name) {
		fmt.Println(lib.Name)
		return nil
	}
	t._wkeSetViewProxy = lib.NewProc("mbSetViewProxy")
	t._wkeSetTransparent = lib.NewProc("mbSetTransparent")
	t._wkeOnDocumentReady2 = lib.NewProc("mbOnDocumentReady")
	t._wkeNetGetRawResponseHead = lib.NewProc("mbNetGetRawResponseHeadInBlinkThread")
	t._wkeNetSetMIMEType = lib.NewProc("mbNetSetMIMEType")
	t._wkeNetGetMIMEType = lib.NewProc("mbNetGetMIMEType")
	t._wkeOnLoadUrlFail = lib.NewProc("mbOnLoadUrlFail")
	t._wkeOnLoadUrlEnd = lib.NewProc("mbOnLoadUrlEnd")
	t._wkeNetContinueJob = lib.NewProc("mbNetContinueJob")
	t._wkeNetHoldJobToAsynCommit = lib.NewProc("mbNetHoldJobToAsynCommit")
	t._wkeNetHookRequest = lib.NewProc("mbNetHookRequest")
	t._wkeNetChangeRequestUrl = lib.NewProc("mbNetChangeRequestUrl")
	t._wkeNetSetHTTPHeaderField = lib.NewProc("mbNetSetHTTPHeaderField")
	t._wkeGetString = lib.NewProc("mbGetString")
	t._wkeOnConsole = lib.NewProc("mbOnConsole")
	t._wkeIsMainFrame = lib.NewProc("mbIsMainFrame")
	t._wkeOnDidCreateScriptContext = lib.NewProc("mbOnDidCreateScriptContext")
	t._wkeKillFocus = lib.NewProc("mbKillFocus")
	t._wkeNetCancelRequest = lib.NewProc("mbNetCancelRequest")
	t._wkeNetSetData = lib.NewProc("mbNetSetData")
	t._wkeNetGetRequestMethod = lib.NewProc("mbNetGetRequestMethod")
	t._wkeFireKeyPressEvent = lib.NewProc("mbFireKeyPressEvent")
	t._wkeFireKeyUpEvent = lib.NewProc("mbFireKeyUpEvent")
	t._wkeFireKeyDownEvent = lib.NewProc("mbFireKeyDownEvent")
	t._wkeFireMouseWheelEvent = lib.NewProc("mbFireMouseWheelEvent")
	t._wkeFireContextMenuEvent = lib.NewProc("mbFireContextMenuEvent")
	t._wkeFireWindowsMessage = lib.NewProc("mbFireWindowsMessage")
	t._wkeCreateWebWindow = lib.NewProc("mbCreateWebWindow")
	t._wkeShowWindow = lib.NewProc("mbShowWindow")
	t._wkeFireMouseEvent = lib.NewProc("mbFireMouseEvent")
	t._wkeOnLoadUrlBegin = lib.NewProc("mbOnLoadUrlBegin")
	t._wkeNetOnResponse = lib.NewProc("mbNetOnResponse")
	t._wkeLoadURL = lib.NewProc("mbLoadURL")
	t._wkeGetHostHWND = lib.NewProc("mbGetHostHWND")
	t._wkeResize = lib.NewProc("mbResize")
	t._wkeOnPaintBitUpdated = lib.NewProc("mbOnPaintBitUpdated")
	t._wkeOnPaintUpdated = lib.NewProc("mbOnPaintUpdated")
	t._wkeSetHandle = lib.NewProc("mbSetHandle")
	t._wkeCreateWebView = lib.NewProc("mbCreateWebView")
	t._wkeInitialize = lib.NewProc("mbInit")
	t._wkeUnInitialize = lib.NewProc("mbUninit")
	t._wkeSetFocus = lib.NewProc("mbSetFocus")
	t._wkeDestroyWebView = lib.NewProc("mbDestroyWebView")
	t._jsGetWebView = lib.NewProc("jsGetWebView")
	t._wkeOnDownload = lib.NewProc("mbOnDownloadInBlinkThread")
	t._wkeOnAlertBox = lib.NewProc("mbOnAlertBox")
	t._wkeOnCreateView = lib.NewProc("mbOnCreateView")
	t._wkeSetContextMenuEnabled = lib.NewProc("mbSetContextMenuEnabled")
	t._wkeSetNavigationToNewWindowEnable = lib.NewProc("mbSetNavigationToNewWindowEnable")
	t._wkeSetUserAgent = lib.NewProc("mbSetUserAgent")
	t._wkePopupDialogAndDownload = lib.NewProc("mbPopupDialogAndDownload")
	t._wkeSetDebugConfig = lib.NewProc("mbSetDebugConfig")
	t._wkeOnJsQuery = lib.NewProc("mbOnJsQuery")
	t._wkeResponseQuery = lib.NewProc("mbResponseQuery")
	t._wkeGetLockedViewDC = lib.NewProc("mbGetLockedViewDC")
	t._wkeRunMessageLoop = lib.NewProc("mbRunMessageLoop")
	t._wkeWebFrameGetMainFrame = lib.NewProc("mbWebFrameGetMainFrame")
	t._wkeRunJs = lib.NewProc("mbRunJs")
	t._wkeOnLoadingFinish = lib.NewProc("mbOnLoadingFinish")
	t._wkeEnableHighDPISupport = lib.NewProc("mbEnableHighDPISupport")
	t._wkeOnTitleChanged = lib.NewProc("mbOnTitleChanged")
	t._wkeNetGetFavicon = lib.NewProc("mbOnNetGetFavicon")
	t._wkeReload = lib.NewProc("mbReload")
	t._wkeGetUrl = lib.NewProc("mbGetUrl")
	t._wkeStopLoading = lib.NewProc("mbStopLoading")
	t._wkeGoBack = lib.NewProc("mbGoBack")
	t._wkeGoForward = lib.NewProc("mbGoForward")
	t._wkeCanGoForward = lib.NewProc("mbCanGoForward")
	var set mbSettings
	set.mask = MB_ENABLE_NODEJS
	r, _, err := t._wkeInitialize.Call(uintptr(unsafe.Pointer(&set)))
	if r != 0 {
		fmt.Println(err.Error())
	}
	return t
}

func GetBound(h win.HWND) win.RECT {
	rect, e := win.GetWindowRect(h)
	if !e {
		return win.RECT{}
	}
	bn := win.RECT{
		Left: rect.Left,
		Top:  rect.Top,
	}
	win.GetClientRect(h, &rect)

	bn.Right = rect.Width() + bn.Left
	bn.Bottom = rect.Height() + bn.Top
	return bn
}

func (t *Thublink) wkeUnInit() {
	t._wkeUnInitialize.Call()
}
func (t *Thublink) wkeOnDownload(wke wkeHandle, callback wkeOnDownloadCallback, param uintptr) {
	t._wkeOnDownload.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (t *Thublink) wkeOnAlertBox(wke wkeHandle, callback wkeOnAlertBoxCallback, param uintptr) {
	t._wkeOnAlertBox.Call(uintptr(wke), syscall.NewCallback(callback), param)
}
func (t *Thublink) wkeOnCreateView(wke wkeHandle, callback wkeOnCreateViewCallback, param uintptr) {
	t._wkeOnCreateView.Call(uintptr(wke), syscall.NewCallback(callback), param)
}
func (t *Thublink) wkeSetContextMenuEnabled(wke wkeHandle, show bool) {
	t._wkeSetContextMenuEnabled.Call(uintptr(wke), uintptr(toBool(show)))
}
func (t *Thublink) wkeSetNavigationToNewWindowEnable(wke wkeHandle, b bool) {
	t._wkeSetNavigationToNewWindowEnable.Call(uintptr(wke), uintptr(toBool(b)))
}
func (t *Thublink) wkeSetUserAgent(wke wkeHandle, ua string) {
	p := strToCharPtr(ua)
	t._wkeSetUserAgent.Call(uintptr(wke), p)
}

func (t *Thublink) wkeSetViewProxy(wke wkeHandle, proxy ProxyInfo) {
	px := wkeProxy{
		Type: int32(proxy.Type),
		Port: uint16(proxy.Port),
	}
	for i, c := range proxy.HostName {
		px.HostName[i] = byte(c)
	}
	if proxy.UserName != "" {
		for i, c := range proxy.UserName {
			px.UserName[i] = byte(c)
		}
	}
	if proxy.Password != "" {
		for i, c := range proxy.Password {
			px.Password[i] = byte(c)
		}
	}
	t._wkeSetViewProxy.Call(uintptr(wke), uintptr(unsafe.Pointer(&px)))
}

func (t *Thublink) wkeSetTransparent(wke wkeHandle, enable bool) {
	t._wkeSetTransparent.Call(uintptr(wke), uintptr(toBool(enable)))
}

func (t *Thublink) wkeOnDocumentReady(wke wkeHandle, callback wkeDocumentReady2Callback, param uintptr) {
	t._wkeOnDocumentReady2.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (t *Thublink) wkeNetGetRawResponseHead(job wkeNetJob) map[string]string {
	r, _, _ := t._wkeNetGetRawResponseHead.Call(uintptr(job))
	var list []string
	slist := *((*wkeSlist)(unsafe.Pointer(r)))
	for slist.str != 0 {
		list = append(list, ptrToUtf8(slist.str))
		if slist.next == 0 {
			break
		} else {
			slist = *((*wkeSlist)(unsafe.Pointer(slist.next)))
		}
	}
	hMap := make(map[string]string)
	for i := 0; i < len(list); i += 2 {
		hMap[list[i]] = list[i+1]
	}
	return hMap
}

func (t *Thublink) wkeNetSetMIMEType(job wkeNetJob, mime string) {
	p := strToCharPtr(mime)
	t._wkeNetSetMIMEType.Call(uintptr(job), p)
}

func (t *Thublink) wkeNetGetMIMEType(job wkeNetJob) string {
	r, _, _ := t._wkeNetGetMIMEType.Call(uintptr(job))
	return ptrToUtf8(r)
}

func (t *Thublink) wkeOnLoadUrlFail(wke wkeHandle, callback wkeLoadUrlFailCallback, param uintptr) {
	t._wkeOnLoadUrlFail.Call(uintptr(wke), syscall.NewCallback(callback), param)
}
func (t *Thublink) wkeOnJsQuery(wke wkeHandle, callback wkeJsQueryCallback, param uintptr) {
	t._wkeOnJsQuery.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (t *Thublink) wkeOnLoadUrlEnd(wke wkeHandle, callback wkeLoadUrlEndCallback, param uintptr) {
	t._wkeOnLoadUrlEnd.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (t *Thublink) wkeNetContinueJob(job wkeNetJob) {
	t._wkeNetContinueJob.Call(uintptr(job))
}

func (t *Thublink) wkeNetHoldJobToAsynCommit(job wkeNetJob) {
	t._wkeNetHoldJobToAsynCommit.Call(uintptr(job))
}

func (t *Thublink) wkeNetHookRequest(job wkeNetJob) {
	t._wkeNetHookRequest.Call(uintptr(job))
}

func (t *Thublink) wkeNetChangeRequestUrl(job wkeNetJob, url string) {
	p := strToCharPtr(url)
	t._wkeNetChangeRequestUrl.Call(uintptr(job), p)
}

func (t *Thublink) wkeNetSetHTTPHeaderField(job wkeNetJob, name, value string) {
	np := strToCharPtr(name)
	vp := strToCharPtr(value)
	t._wkeNetSetHTTPHeaderField.Call(uintptr(job), np, vp)
}

func (t *Thublink) wkeGetString(str wkeString) string {
	r, _, _ := t._wkeGetString.Call(uintptr(str))
	return ptrToUtf8(r)
}

func (t *Thublink) wkeOnConsole(wke wkeHandle, callback wkeConsoleCallback, param uintptr) {
	t._wkeOnConsole.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (t *Thublink) wkeIsMainFrame(wke wkeHandle, frame wkeFrame) bool {
	r, _, _ := t._wkeIsMainFrame.Call(uintptr(wke), uintptr(frame))
	return r != 0
}

func (t *Thublink) wkeOnDidCreateScriptContext(wke wkeHandle, callback wkeDidCreateScriptContextCallback, param uintptr) {
	t._wkeOnDidCreateScriptContext.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (t *Thublink) wkeKillFocus(wke wkeHandle) {
	t._wkeKillFocus.Call(uintptr(wke))
}

func (t *Thublink) jsGetWebView(es jsExecState) wkeHandle {
	r, _, _ := t._jsGetWebView.Call(uintptr(es))
	return wkeHandle(r)
}

func (t *Thublink) wkeDestroyWebView(wke wkeHandle) {
	t._wkeDestroyWebView.Call(uintptr(wke))
}

func (t *Thublink) wkeNetCancelRequest(job wkeNetJob) {
	t._wkeNetCancelRequest.Call(uintptr(job))
}

func (t *Thublink) wkeNetOnResponse(wke wkeHandle, callback wkeNetResponseCallback, param uintptr) {
	t._wkeNetOnResponse.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (t *Thublink) wkeOnLoadUrlBegin(wke wkeHandle, callback wkeLoadUrlBeginCallback, param uintptr) {
	t._wkeOnLoadUrlBegin.Call(uintptr(wke), syscall.NewCallback(callback), param)
}

func (t *Thublink) wkeNetGetRequestMethod(job wkeNetJob) wkeRequestType {
	r, _, _ := t._wkeNetGetRequestMethod.Call(uintptr(job))
	return wkeRequestType(r)
}

func (t *Thublink) wkeNetSetData(job wkeNetJob, buf []byte) {
	if len(buf) == 0 {
		buf = []byte{0}
	}
	t._wkeNetSetData.Call(uintptr(job), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
}

func (t *Thublink) wkeSetFocus(wke wkeHandle) {
	t._wkeSetFocus.Call(uintptr(wke))
}

func (t *Thublink) wkeFireKeyPressEvent(wke wkeHandle, code, flags uint32, isSysKey bool) bool {
	ret, _, _ := t._wkeFireKeyPressEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func (t *Thublink) wkeFireKeyDownEvent(wke wkeHandle, code, flags uint32, isSysKey bool) bool {
	ret, _, _ := t._wkeFireKeyDownEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func (t *Thublink) wkeFireKeyUpEvent(wke wkeHandle, code, flags uint32, isSysKey bool) bool {
	ret, _, _ := t._wkeFireKeyUpEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func (t *Thublink) wkeFireMouseWheelEvent(wke wkeHandle, x, y, delta, flags int32) bool {
	r, _, _ := t._wkeFireMouseWheelEvent.Call(
		uintptr(wke),
		uintptr(x),
		uintptr(y),
		uintptr(delta),
		uintptr(flags))
	return byte(r) != 0
}
func (t *Thublink) wkeFireContextMenuEvent(wke wkeHandle, x, y, flags int32) bool {
	r, _, _ := t._wkeFireContextMenuEvent.Call(
		uintptr(wke),
		uintptr(x),
		uintptr(y),
		uintptr(flags))
	return byte(r) != 0
}
func (t *Thublink) wkeFireWindowsMessage(wke wkeHandle, hWnd win.HWND, message, wParam, lParam int32) bool {
	r, _, _ := t._wkeFireWindowsMessage.Call(
		uintptr(wke),
		uintptr(hWnd),
		uintptr(message),
		uintptr(wParam),
		uintptr(lParam),
		uintptr(0))
	return byte(r) != 0
}

func (t *Thublink) wkeCreateWebWindow(wt WindowType, parent win.HWND, x, y, width, height int32) wkeHandle {
	r, _, _ := t._wkeCreateWebWindow.Call(
		uintptr(wt),
		uintptr(parent),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
	return wkeHandle(r)
}

func (t *Thublink) wkeShowWindow(wke wkeHandle, show bool) {
	t._wkeShowWindow.Call(uintptr(wke), uintptr(toBool(show)))
}

func (t *Thublink) wkeFireMouseEvent(wke wkeHandle, message, x, y, flags int32) bool {
	r, _, _ := t._wkeFireMouseEvent.Call(
		uintptr(wke),
		uintptr(message),
		uintptr(x),
		uintptr(y),
		uintptr(flags))
	return byte(r) != 0
}

func (t *Thublink) wkeResize(wke wkeHandle, w, h uint32) {
	t._wkeResize.Call(uintptr(wke), uintptr(w), uintptr(h))
}

func (t *Thublink) wkeLoadURL(wke wkeHandle, url string) {
	ptr := strToCharPtr(url)
	t._wkeLoadURL.Call(uintptr(wke), ptr)
}

/*
设置一些实验性选项。debugString可用参数有：
*/
func (t *Thublink) wkeSetDebugConfig(wke wkeHandle, debug debugType, param string) {
	dp := strToCharPtr(string(debug))
	pp := strToCharPtr(param)
	t._wkeSetDebugConfig.Call(uintptr(wke), dp, pp)
}

func (t *Thublink) wkeOnPaintBitUpdated(wke wkeHandle, callback wkePaintBitUpdatedCallback, param uintptr) {
	t._wkeOnPaintBitUpdated.Call(uintptr(wke), syscall.NewCallback(callback), param)
}
func (t *Thublink) wkeOnPaintUpdated(wke wkeHandle, callback wkePaintUpdatedCallback, param uintptr) {
	t._wkeOnPaintUpdated.Call(uintptr(wke), syscall.NewCallback(callback), param)
}
func (t *Thublink) wkeOnLoadingFinish(wke wkeHandle, callback wkeLoadingFinishCallback, param uintptr) {
	t._wkeOnLoadingFinish.Call(uintptr(wke), syscall.NewCallback(callback), param)
}
func (t *Thublink) wkeOnTitleChanged(wke wkeHandle, callback wkeTitleChangedCallback, param uintptr) {
	t._wkeOnTitleChanged.Call(uintptr(wke), syscall.NewCallback(callback), param)
}
func (t *Thublink) wkeEnableHighDPISupport() {
	t._wkeEnableHighDPISupport.Call()
}
func (t *Thublink) wkeNetGetFavicon(wke wkeHandle, callback wkeNetGetFaviconCallback, param uintptr) {
	t._wkeNetGetFavicon.Call(uintptr(wke), syscall.NewCallback(callback), param)
}
func (t *Thublink) wkeRunJs(handle wkeHandle, frame wkeFrame, script uintptr, isInClosure bool, callback wkeRunJsCallback, param, unUse uintptr) {
	t._wkeRunJs.Call(uintptr(handle), uintptr(frame), script, uintptr(toBool(isInClosure)), 0, param, unUse)
}

func (t *Thublink) wkeSetHandle(wke wkeHandle, handle uintptr) {
	t._wkeSetHandle.Call(uintptr(wke), handle)
}
func (t *Thublink) wkeOnShowDevtoolsCallback(wke uintptr, param uintptr) uintptr {
	return 0
}
func (t *Thublink) wkeCreateWebView() wkeHandle {
	r, _, _ := t._wkeCreateWebView.Call()
	return wkeHandle(r)
}
func (t *Thublink) wkeGetHostHWND() win.HWND {
	r, _, _ := t._wkeGetHostHWND.Call()
	return win.HWND(r)
}

func (t *Thublink) wkeGetLockedViewDC(wke wkeHandle) win.HDC {
	r, _, _ := t._wkeGetLockedViewDC.Call(uintptr(wke))
	return win.HDC(r)
}
func (t *Thublink) wkeRunMessageLoop() {
	t._wkeRunMessageLoop.Call()
}
func (t *Thublink) wkeWebFrameGetMainFrame(wke wkeHandle) wkeFrame {
	r, _, _ := t._wkeWebFrameGetMainFrame.Call(uintptr(wke))
	return wkeFrame(r)
}
func (t *Thublink) wkeReload(wke wkeHandle) {
	t._wkeReload.Call(uintptr(wke))
}

func (t *Thublink) wkeGetUrl(wke wkeHandle) uintptr {
	r, _, _ := t._wkeGetUrl.Call(uintptr(wke))
	return r
}

func (t *Thublink) wkeStopLoading(wke wkeHandle) {
	t._wkeStopLoading.Call(uintptr(wke))
	return
}
func (t *Thublink) wkeGoBack(wke wkeHandle) {
	t._wkeGoBack.Call(uintptr(wke))
	return
}

func (t *Thublink) wkeGoForward(wke wkeHandle) {
	t._wkeGoForward.Call(uintptr(wke))
	return
}
func (t *Thublink) wkeCanGoForward(wke wkeHandle, callback wkeCanGoBackForwardCallback, param uintptr) {
	t._wkeGoForward.Call(uintptr(wke), syscall.NewCallback(callback), param)
	return
}
func strToCharPtr(str string) uintptr {
	buf := []byte(str)
	rs := make([]byte, len(str)+1)
	for i, v := range buf {
		rs[i] = v
	}
	return uintptr(unsafe.Pointer(&rs[0]))
}
func StrPtr(s string) (uintptr, error) {
	fromString, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		return 0, err
	}
	return uintptr(unsafe.Pointer(fromString)), nil
}

func ptrToByte(ptr uintptr, len uint32) []byte {
	seq := make([]byte, len)
	if ptr > 0 {
		for i := uint32(0); i < len; i++ {
			seq[i] = *((*byte)(unsafe.Pointer(ptr)))
			ptr++
		}
	}
	return seq
}
func ptrToUtf8(ptr uintptr) string {
	var seq []byte
	for ptr > 0 {
		b := *((*byte)(unsafe.Pointer(ptr)))
		if b != 0 {
			seq = append(seq, b)
			ptr++
		} else {
			break
		}
	}
	return string(seq)
}

func toBool(b bool) byte {
	if b {
		return 1
	} else {
		return 0
	}
}
