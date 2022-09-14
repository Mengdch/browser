package win32

import (
	"github.com/Mengdch/browser/log"
	"github.com/Mengdch/goUtil/OS"
	"github.com/Mengdch/goUtil/TypeTools"
)

func init() {
	thuOS.SetLog(logRecord, nil)
}
func logRecord(value, error string) {
	defer log.CatchPanic("logRecord")
	log.Log(value+":"+TypeTools.OutJson(urls), error)
}
