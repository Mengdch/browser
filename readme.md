# 概述 

基于win32API和miniblink封装的GOLANG使用的WebUI，使用原生窗体加内嵌视图的方式实现。目前已完成常用接口封装，后续会逐步完善。
直接调用DLL，未使用CGO。

----
miniblink是一个开源的、单文件、且目前已知的最小的基于chromium的，浏览器控件。
通过其导出的纯C接口，几行代码即可创建一个浏览器控件。
您可以通过官网http://miniblink.net 来获取更多的关于miniblink的信息。

# 使用前准备
普通版本可以自行到 https://github.com/weolar/miniblink49/releases 下载最新编译后的 dll 替换就行。
VIP版本需进行购买https://miniblink.net/views/features.html
## 注意
1. 下载的名称默认是 node.dll，需跟代码中function.go中dll_name值一致）。
2. 默认代码版本为VIP接口，普通接口可以自行修改function.go中加载函数名将mb前缀改为wke。

----

# 使用
创建一个加载指定链接的窗口：

```go
func main() {
	defer log.CatchPanic("main")
	url := flag.String("url", "https://www.baidu.com", "链接")
	title := flag.String("title", "aa", "标题")
	noTitles := flag.String("noHead", "[]", "无标题域名")
	ico := flag.String("icon", "", "图标")
	dev := flag.String("dev", "", "调试目录")
	ua := flag.String("ua", "", "UserAgent")
	max := flag.Bool("max", true, "初始最大化")
	width := flag.Int("width", 1600, "宽")
	height := flag.Int("height", 900, "高")
	flag.Parse()
	var domains []string
	err := json.Unmarshal([]byte(*noTitles), &domains)
	if err != nil {
		return
	}
	userAgent = *ua
	if len(*url) == 0 {
		return
	}
	jsFunc := map[int32]func(string) string{
		1: func(sha string) string {
			s := string(findOne(sha, keyTable))
			return s
		},
	}
	err = browser.StartFull(*url, *title, *ico, userAgent, *dev, *max, true, true, true, *width, *height, thuOS.Center, finish, save, jsFunc, nil, nil, nil, domains)
	if err != nil {
		log.Log(*title+":"+*url, err.Error())
	}
}
```
具体可参考demo文件夹下文件
## miniblink文档

- [miniblink的介绍](https://zhuanlan.zhihu.com/p/22611497?group_id=764036386641707008)
- [普通版本API文档](https://miniblink.net/views/doc/index.html)
- [VIP版本API文档](https://miniblink.net/views/doc/api-doc-vip.html)
----


