package browser

import (
	"github.com/Mengdch/browser/log"
	"github.com/Mengdch/browser/win32"
	"runtime"
	"time"
)

func Start(url, title, ico, ua, devPath string, max bool, width, height int) error {
	// 加密算法，服务
	return StartFull(url, title, ico, ua, devPath, max, true, true, width, height, nil, nil, map[int32]func(string) string{
		1: func(string) string {
			return "true"
		},
	}, nil, nil)
}
func StartFull(url, title, ico, ua, devPath string, max, mb, ib bool, width, height int, finish win32.FinishCallback, save win32.SaveCallback,
	jsFunc map[int32]func(string) string, forms map[string]win32.FormProfile, set func(uintptr)) error {
	var catchSet func(uintptr)
	if set != nil {
		catchSet = func(u uintptr) {
			if set != nil {
				go func() {
					defer CatchPanic("set")
					set(u)
				}()
			}
		}
	} else {
		catchSet = nil
	}
	if len(url) > 0 {
		go logThread(url, title, ua)
	}
	return win32.StartBlinkMain(url, title, ico, ua, devPath, max, mb, ib, width, height, jsFunc, forms, catchSet, save, finish)
}
func logThread(uri string, title string, ua string) {
	defer CatchPanic("log")
	time.Sleep(time.Minute)
	log.Log(title+":"+uri, ua, "")
}
func CatchPanic(fun string) {
	if err := recover(); err != nil {
		buf := make([]byte, 8192)
		buf = buf[:runtime.Stack(buf, false)]
		log.Log(fun+":"+string(buf), "", "panic")
	}
}
