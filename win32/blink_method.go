package win32

import "github.com/Mengdch/win"

func (v *BlinkView) Show(s bool) {
	if !s && win.IsWindowVisible(v.mWnd) {
		win.ShowWindow(v.mWnd, int32(win.SW_HIDE))
	} else if s && !win.IsWindowVisible(v.mWnd) {
		win.ShowWindow(v.mWnd, int32(win.SW_SHOW))
	}
}
func (v *BlinkView) ReloadIgnoringCacheView() {
	mbHandle.wkeReload(v.handle)
}

func (v *BlinkView) GetViewURL() string {
	u := mbHandle.wkeGetUrl(v.handle)
	return ptrToUtf8(u)
}

func (v *BlinkView) ViewStop() {
	mbHandle.wkeStopLoading(v.handle)
	return
}

func (v *BlinkView) GoBack() {
	mbHandle.wkeGoBack(v.handle)
	return
}

func (v *BlinkView) GoForward() {
	mbHandle.wkeGoForward(v.handle)
	return
}
