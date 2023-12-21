package win32

//onceEvent  onEvent

//ready-to-show\show\restore\minimize\move\focus\blur\resize\close\hide

//will-quit before-quit

const (
	SYSTEM_RESIZE = "resize"
)

func (w *Window) callback(key, param string) bool {
	if w.events != nil {
		c := w.events[key]
		if c != nil {
			c(param)
			return true
		}
	}
	return false
}
