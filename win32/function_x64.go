//go:build amd64
// +build amd64

package win32

import (
	"github.com/Mengdch/win"
	"syscall"
)

func getName() string {
	return dll_name + "_x64.dll"
}

type wkeJsQueryCallback func(wke wkeHandle, param uintptr, es jsExecState, queryId uintptr, customMsg int32, request uintptr) uintptr
type wkeRunJsCallback func(wke wkeHandle, param uintptr, es jsExecState, val1, val2 uintptr)

func (v *BlinkView) setProc(init bool) {
	var value uintptr
	if init {
		value = syscall.NewCallback(v.OnWndProc)
	} else {
		value = 0
	}
	v.proc = win.SetWindowLongPtr(v.mWnd, win.GWL_WNDPROC, value)
}
func (v *BlinkView) onJsQuery(wke wkeHandle, param uintptr, es jsExecState, queryId uintptr, customMsg int32, request uintptr) uintptr {
	if f, e := v.fnMap[customMsg]; e {
		response := f(ptrToUtf8(request))
		mbHandle.wkeResponseQuery(wke, queryId, customMsg, response)
	}
	return 0
}
func (t *Thublink) wkeResponseQuery(wke wkeHandle, queryId uintptr, customMsg int32, response string) {
	t._wkeResponseQuery.Call(uintptr(wke), queryId, uintptr(customMsg), strToCharPtr(response))
}
func (v *BlinkView) onRunJs(wke wkeHandle, param uintptr, es jsExecState, val1, val2 uintptr) {

}
