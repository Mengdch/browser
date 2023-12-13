package win32

import (
	"fmt"
	"github.com/Mengdch/win"
)

type menuItem struct {
	hm   win.HMENU
	x, y int32
	cmd  func(int)
}

func newMenu(hWnd win.HWND, items []string, base int, left bool, x, y int32) {
	hm := win.CreatePopupMenu()
	for x, i := range items {
		var s bool
		var e error
		if len(i) == 0 {
			s, e = win.AppendMenu(hm, win.UINT(win.MF_SEPARATOR), 0, 0)
		} else {
			s, e = win.AppendMenu(hm, win.UINT(win.MF_STRING), win.UINT(x+base), win.String2UIntPtr(i))
		}
		if !s {
			fmt.Println(s, e)
		}
	}
	p := win.POINT{x, y}
	//win.ClientToScreen(hWnd, &p)
	var ff uint32
	if left {
		ff = win.TPM_LEFTALIGN
	} else {
		ff = win.TPM_RIGHTALIGN
	}
	win.TrackPopupMenuEx(hm, ff, p.X, p.Y, hWnd, nil)
	win.DestroyMenu(hm)
}

//func newMenu(items []string, base int) win.HMENU {
//	hm, _, err := procCreatePopupMenu.Call()
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	hmenu := win.HMENU(hm)
//	u := uintptr(hmenu)
//	fmt.Println(hm, hmenu, u)
//	for x, i := range items {
//		s, _, e := procAppendMenu.Call(uintptr(hmenu), uintptr(win.MF_STRING), uintptr(x+base), win.String2UIntPtr(i))
//		if e != nil {
//			fmt.Println(s, e.Error())
//		}
//	}
//	return hmenu
//}
//func popMenu(hWnd win.HWND, hm win.HMENU, x, y int32) {
//	p := win.POINT{x, y}
//	//win.ClientToScreen(hWnd, &p)
//	fmt.Println(hm, x, y, p, hWnd)
//	ret, _, err := procTrackPopupMenu.Call(uintptr(hm), win.TPM_LEFTALIGN|win.TPM_BOTTOMALIGN|win.TPM_RIGHTBUTTON,
//		uintptr(p.X), uintptr(p.Y), uintptr(hWnd), uintptr(0))
//	//ret := win.TrackPopupMenu(hm, win.TPM_LEFTALIGN|win.TPM_BOTTOMALIGN|win.TPM_RIGHTBUTTON, p.X, p.Y, 0, hWnd, nil)
//	if ret == 0 && err != nil {
//		fmt.Println(err.Error())
//	}
//}
