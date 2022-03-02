// +build 386

package win32

import (
	"github.com/Mengdch/goUtil/TypeTools"
	"github.com/Mengdch/win"
	"golang.org/x/sys/windows"
	"syscall"
)

func load() *windows.LazyDLL {
	return windows.NewLazyDLL(dll_name + "_x86.dll")
}

type wkeJsQueryCallback func(wke wkeHandle, param uintptr, es jsExecState, queryId uintptr, queryId2 uintptr, customMsg int32, request uintptr) uintptr
type mbRunJsCallback func(wke wkeHandle, param uintptr, es jsExecState, val uintptr)

func (v *BlinkView) setProc(init bool) {
	var value int32
	if init {
		value = int32(syscall.NewCallback(v.OnWndProc))
	} else {
		value = 0
	}
	v.proc = uintptr(win.SetWindowLong(v.mWnd, win.GWL_WNDPROC, value))
}
func (v *BlinkView) onJsQuery(wke wkeHandle, param uintptr, es jsExecState, queryId uintptr, queryId2 uintptr, customMsg int32, request uintptr) uintptr {
	if f, e := v.fnMap[customMsg]; e {
		response := f(ptrToUtf8(request))
		mbHandle.wkeResponseQuery(wke, queryId, queryId2, customMsg, TypeTools.OutJson(response))
	}
	return 0
}
func (t *Thublink) wkeResponseQuery(wke wkeHandle, queryId uintptr, queryId2 uintptr, customMsg int32, response string) {
	t._wkeResponseQuery.Call(uintptr(wke), queryId, queryId2, uintptr(customMsg), strToCharPtr(response))
}
func (v *BlinkView) onRunJs(wke wkeHandle, param uintptr, es jsExecState, val uintptr) {

}
