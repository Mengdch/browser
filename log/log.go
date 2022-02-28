package log

import (
	"io/ioutil"
	"net/http"
	"runtime"
)

var userAgent string

func SetUA(ua string) {
	userAgent = ua
}
func Log(value, error string) {
}
func CatchPanic(fun string) {
	if err := recover(); err != nil {
		buf := make([]byte, 8192)
		buf = buf[:runtime.Stack(buf, false)]
		Log(fun+":"+string(buf), "panic")
	}
}

func GetBody(req *http.Request) []byte {
	client := &http.Client{}
	response, err := client.Do(req)
	if nil != err {
		return nil
	}
	if response.Body != nil {
		var body []byte
		body, err = ioutil.ReadAll(response.Body)
		response.Body.Close()
		if nil != err {
			return nil
		}
		return body
	}
	return nil
}
