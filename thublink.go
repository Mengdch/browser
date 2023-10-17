package browser

import (
	"github.com/Mengdch/browser/log"
	"github.com/Mengdch/browser/win32"
	thuOS "github.com/Mengdch/goUtil/OS"
)

func Start(url, title, ico, ua, devPath string, max bool, width, height int) error {
	// 加密算法，服务
	return StartFull(url, title, ico, ua, devPath, max, true, true, true, width, height, thuOS.Center, nil, nil, map[int32]func(string) string{
		1: func(string) string {
			return "true"
		},
	}, nil, nil, nil)
}
func StartFull(url, title, ico, ua, devPath string, max, mb, ib, show bool, width, height, pos int, finish win32.FinishCallback, save win32.SaveCallback,
	jsFunc map[int32]func(string) string, forms map[string]win32.FormProfile, set func(uintptr), domains []string) error {
	catchSet := getSet(set)
	log.SetUA(ua)
	return win32.StartBlinkMain(url, title, ico, ua, devPath, max, mb, ib, show, max, width, height, pos, jsFunc, forms, catchSet, save, finish, domains)
}
func getSet(set func(uintptr)) func(uintptr) {
	var catchSet func(uintptr)
	if set != nil {
		catchSet = func(u uintptr) {
			if set != nil {
				go func() {
					defer log.CatchPanic("set")
					set(u)
				}()
			}
		}
	} else {
		catchSet = nil
	}

	return catchSet
}
func RunNoWindow(ico, title, ua, devPath string, max, mb, ib bool, width, height int,
	jsFunc map[int32]func(string) string, set func(uintptr)) error {
	log.SetUA(ua)
	return win32.StartBlinkMain("", title, ico, ua, devPath, max, mb, ib, false, true, width, height, thuOS.Center, jsFunc,
		nil, getSet(set), nil, nil, nil)
}
func Show(url, script string, x, y int32) {
	win32.ShowMainWindow(url, script, x, y)
}
