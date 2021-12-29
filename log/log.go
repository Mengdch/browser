package log

import (
	"io/ioutil"
	"net/http"
)

var userAgent string

func Log(value, ua, error string) {
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
