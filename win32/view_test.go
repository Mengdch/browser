package win32

import (
	"net/url"
	"testing"
)

func TestStart(t *testing.T) {
	parse, err := url.Parse("http://www.baidu.com")
	t.Log(parse.Scheme, err)
	parse, err = url.Parse("https://www.baidu.com")
	t.Log(parse.Scheme, err)
	parse, err = url.Parse("ftp://www.baidu.com")
	t.Log(parse.Scheme, err)
}
