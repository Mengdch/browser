package win32

import (
	"strconv"
	"unsafe"
)

// onContentsEvent
const (
	DidStopLoading     = "did-stop-loading" //mbOnLoadingFinish
	PageTitleUpdated   = "page-title-updated"
	PageFaviconUpdated = "page-favicon-updated"
)

func (v *BlinkView) wkeOnDocumentReady(wke wkeHandle, param uintptr, frame wkeFrame) uintptr {
	if !mbHandle.wkeIsMainFrame(wke, frame) {
		return 0
	}
	v.runJs(frame, v.readyJs)
	v.readyJs = ""
	v.callback(wke, frame, DidStopLoading, "")
	return 0
}
func (v *BlinkView) SetOnNewWindow(callback wkeOnCreateViewCallback) {
	mbHandle.wkeOnCreateView(v.handle, callback, 0)
}
func (v *BlinkView) callback(wke wkeHandle, frame wkeFrame, key, param string) {
	if v.call != nil {
		c := v.call[key]
		if c != nil {
			c(EventData{wke, frame, param})
		}
	}
}
func (v *BlinkView) titleChanged(wke wkeHandle, param, title uintptr) uintptr {
	v.callback(wke, mbHandle.wkeWebFrameGetMainFrame(wke), PageTitleUpdated, ptrToUtf8(title))
	return 0
}
func (v *BlinkView) wkeGetFavicon(wke wkeHandle, param, url, buf uintptr) uintptr {
	if buf == 0 {
		return 0
	}
	mb := *((*wkeMemBuf)(unsafe.Pointer(buf)))
	var icon string
	if mb.length > 0 {
		bs := ptrToByte(mb.data, mb.length)
		icon = base64.StdEncoding.EncodeToString(bs)
	}
	v.callback(wke, mbHandle.wkeWebFrameGetMainFrame(wke), PageFaviconUpdated, icon)
	return 0
}
func (v *BlinkView) wkeLoadingFinishCallback(wke wkeHandle, param uintptr, frame wkeFrame, url uintptr, result wkeLoadingResult, reason uintptr) uintptr {
	if !mbHandle.wkeIsMainFrame(wke, frame) {
		return 0
	}
	mbHandle.wkeNetGetFavicon(wke, v.wkeGetFavicon, param)
	//v.callback(wke, frame, DidStopLoading, "")
	return 0
}
func (v *BlinkView) wkeLoadUrlEndCallback(wke wkeHandle, param, url uintptr, job wkeNetJob, buf uintptr, count int32) uintptr {
	return 0
}

func (v *BlinkView) wkeLoadUrlBeginCallback(wke wkeHandle, param, utf8Url uintptr, job wkeNetJob) uintptr {
	uri := ptrToUtf8(utf8Url)
	if len(v.url) > 0 {
		if len(v.preJs) > 0 {
			frame := mbHandle.wkeWebFrameGetMainFrame(wke)
			v.runJs(frame, v.preJs)
			v.preJs = ""
		}
		go logRecord("loadUrlBegin:"+v.url, "")
		v.url = ""
		if v.parent != nil {
			v.parent.show()
		}
	}

	return operateUri(uri)
}
func (v *BlinkView) onCanGoForward(wke wkeHandle, param uintptr, state, buf int32) uintptr {
	fmt.Println("onCanGoForward", state, buf)
	if state == 0 {
		v.canGoForwardChan <- buf != 0
	} else {
		v.canGoForwardChan <- false
	}
	return 0
}
