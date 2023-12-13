package win32

import (
	"github.com/Mengdch/win"
	"golang.org/x/sys/windows"
	"unsafe"
)

func (w *Window) close() {
	win.SendMessage(w.hWnd, win.WM_CLOSE, 0, 0)
}
func (w *Window) min() {
	win.ShowWindow(w.hWnd, win.SW_MINIMIZE)
}
func (w *Window) max() {
	if w.profile.noTitle() {
		if w.last.Top == w.last.Bottom {
			w.last, _ = win.GetWindowRect(w.hWnd)
			var rect win.RECT
			win.SystemParametersInfo(win.SPI_GETWORKAREA, 0, unsafe.Pointer(&rect), 0)
			win.MoveWindow(w.hWnd, 0, 0, rect.Width(), rect.Height(), true)
		}
	} else {
		win.ShowWindow(w.hWnd, win.SW_MAXIMIZE)
	}
}
func (w *Window) restore() {
	if w.profile.noTitle() {
		if w.last.Bottom > w.last.Top {
			win.MoveWindow(w.hWnd, w.last.Left, w.last.Top, w.last.Width(), w.last.Height(), false)
			w.last.Bottom = w.last.Top
		}
	} else {
		win.ShowWindow(w.hWnd, win.SW_RESTORE)
	}
}

//setFullScreen\setPosition\setResizable\setTitle(任务栏）、winHide、winFocus、winIsMinimized、winMaximize、winMinimize、winVisible、winRestore、winUnmaximize

//getAllDisplaysFn getScreenWorkAreaSize

// dialogOpen
func (w *Window) ShowMessageBox(title, msg string) {
	content := windows.StringToUTF16Ptr(msg)
	win.MessageBox(0, content, windows.StringToUTF16Ptr(title), win.MB_ICONINFORMATION|win.MB_OK)
	return
}
