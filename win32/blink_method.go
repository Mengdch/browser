package win32

import "github.com/Mengdch/win"

func (v *BlinkView) Show(s bool) {
	if !s && win.IsWindowVisible(v.mWnd) {
		win.ShowWindow(v.mWnd, int32(win.SW_HIDE))
	} else if s && !win.IsWindowVisible(v.mWnd) {
		win.ShowWindow(v.mWnd, int32(win.SW_SHOW))
	}
}
