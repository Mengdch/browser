package win32

import (
	"github.com/Mengdch/browser/log"
	"github.com/Mengdch/goUtil/OS"
	"github.com/Mengdch/goUtil/TypeTools"
	"github.com/Mengdch/win"
)

func init() {
	thuOS.SetLog(logRecord, nil)
}
func logRecord(value, error string) {
	defer log.CatchPanic("logRecord")
	log.Log(value+":"+TypeTools.OutJson(urls), error)
}
func GetDPI(wndHandle win.HWND) float64 {
	m := win.MonitorFromWindow(wndHandle, win.MONITOR_DEFAULTTONEAREST)
	if m != 0 {
		hdc := win.GetDC(wndHandle)
		var realDpi float64
		dpiA := float64(win.GetDeviceCaps(hdc, win.DESKTOPHORZRES)) / float64(win.GetDeviceCaps(hdc, win.HORZRES))
		dpiB := float64(win.GetDeviceCaps(hdc, win.LOGPIXELSX)) / float64(96)
		if dpiA == dpiB {
			realDpi = dpiA
		} else {
			if dpiA == 1 {
				realDpi = dpiB
			} else {
				realDpi = dpiA
			}
		}
		win.ReleaseDC(wndHandle, hdc)
		return realDpi
	}
	return 1.0
}
