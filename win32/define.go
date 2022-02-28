package win32

type (
	wkeString   uintptr
	wkeFrame    uintptr
	wkeHandle   uintptr
	jsExecState uintptr
)

const (
	WKE_LBUTTON  = 0x01
	WKE_RBUTTON  = 0x02
	WKE_SHIFT    = 0x04
	WKE_CONTROL  = 0x08
	WKE_MBUTTON  = 0x10
	WKE_EXTENDED = 0x0100
	WKE_REPEAT   = 0x4000
)

type debugType string

const (
	showDevTools           debugType = "showDevTools"           // 开启开发者工具，此时param要填写开发者工具的资源路径，如file:///c:/miniblink-release/front_end/inspector.html，注意 debugType = ""//路径中不能有中文，并且必须为完整路径。
	wakeMinInterval        debugType = "wakeMinInterval"        // 设置帧率，默认值是10，值越大帧率越低。
	drawMinInterval        debugType = "drawMinInterval"        // 设置帧率，默认值是3，值越大帧率越低。
	antiAlias              debugType = "antiAlias"              // 设置抗锯齿渲染，param必须设置为“1”。
	minimumFontSize        debugType = "minimumFontSize"        // 最小字体。
	minimumLogicalFontSize debugType = "minimumLogicalFontSize" // 最小逻辑字体。
	defaultFontSize        debugType = "defaultFontSize"        // 默认字号。
	defaultFixedFontSize   debugType = "defaultFixedFontSize"   // 默认Fixed字号。
	imageEnable            debugType = "imageEnable"            // 是否打开无图模式，param为“0”表示开启无图模式。
	jsEnable               debugType = "jsEnable"               // 是否禁用js，param为“0”表示禁用。

)

type WindowType int

const (
	WKE_WINDOW_TYPE_POPUP WindowType = iota
	WKE_WINDOW_TYPE_TRANSPARENT
	WKE_WINDOW_TYPE_CONTROL
)

type ProxyType int

const (
	ProxyType_NONE ProxyType = iota
	ProxyType_HTTP
	ProxyType_SOCKS4
	ProxyType_SOCKS4A
	ProxyType_SOCKS5
	ProxyType_SOCKS5HOSTNAME
)

type ProxyInfo struct {
	Type     ProxyType
	HostName string
	Port     int
	UserName string
	Password string
}
type wkeProxy struct {
	Type     int32
	HostName [100]byte
	Port     uint16
	UserName [50]byte
	Password [50]byte
}

type wkeSlist struct {
	str  uintptr
	next uintptr
}

type mbSettingMask uint32

const (
	MB_SETTING_PROXY            mbSettingMask = 1
	MB_ENABLE_NODEJS            mbSettingMask = 1 << 3
	MB_ENABLE_DISABLE_H5VIDEO   mbSettingMask = 1 << 4
	MB_ENABLE_DISABLE_PDFVIEW   mbSettingMask = 1 << 5
	MB_ENABLE_DISABLE_CC        mbSettingMask = 1 << 6
	MB_ENABLE_ENABLE_EGLGLES2   mbSettingMask = 1 << 7
	MB_ENABLE_ENABLE_SWIFTSHAER mbSettingMask = 1 << 8
)

type mbSettings struct {
	mbProxy                      ProxyInfo
	mask                         mbSettingMask
	blinkThreadInitCallback      mbOnBlinkThreadInitCallback
	blinkThreadInitCallbackParam uintptr
	version                      uintptr
	mainDllPath                  uintptr
	mainDllHandle                uintptr
}
type jsType uint32

const (
	jsType_NUMBER jsType = iota
	jsType_STRING
	jsType_BOOLEAN
	jsType_OBJECT
	jsType_FUNCTION
	jsType_UNDEFINED
	jsType_ARRAY
	jsType_NULL
)

type jsData struct {
	name [100]byte
	propertyGet,
	propertySet,
	finalize,
	callAsFunction uintptr
}

type jsKeys struct {
	length uint32
	first  uintptr
}

type wkeRequestType int

const (
	wkeRequestType_Unknow = 1
	wkeRequestType_Get    = 2
	wkeRequestType_Post   = 3
	wkeRequestType_Put    = 4
)

type wkeKeyFlags int

const (
	wkeKeyFlags_Extend wkeKeyFlags = 0x0100
	wkeKeyFlags_Repeat wkeKeyFlags = 0x4000
)

type wkeRect struct {
	x, y, w, h int32
}

type wkeNetJob uintptr

type wkeMouseFlags int

const (
	wkeMouseFlags_None    wkeMouseFlags = 0
	wkeMouseFlags_LBUTTON wkeMouseFlags = 0x01
	wkeMouseFlags_RBUTTON wkeMouseFlags = 0x02
	wkeMouseFlags_SHIFT   wkeMouseFlags = 0x04
	wkeMouseFlags_CONTROL wkeMouseFlags = 0x08
	wkeMouseFlags_MBUTTON wkeMouseFlags = 0x10
)

type wkeConsoleLevel int

const (
	wkeConsoleLevel_Log          wkeConsoleLevel = 1
	wkeConsoleLevel_Warning      wkeConsoleLevel = 2
	wkeConsoleLevel_Error        wkeConsoleLevel = 3
	wkeConsoleLevel_Debug        wkeConsoleLevel = 4
	wkeConsoleLevel_Info         wkeConsoleLevel = 5
	wkeConsoleLevel_RevokedError wkeConsoleLevel = 6
)

type wkeLoadingResult int

const (
	WKE_LOADING_SUCCEEDED wkeLoadingResult = 0
	WKE_LOADING_FAILED    wkeLoadingResult = 1
	WKE_LOADING_CANCELED  wkeLoadingResult = 2
)

type wkeDownloadOpt int

const (
	kWkeDownloadOptCancel    wkeDownloadOpt = 0
	kWkeDownloadOptCacheData wkeDownloadOpt = 1
)

type wkeNavigationType int

const (
	WKE_NAVIGATION_TYPE_LINKCLICK     wkeNavigationType = 0
	WKE_NAVIGATION_TYPE_FORMSUBMITTE  wkeNavigationType = 1
	WKE_NAVIGATION_TYPE_BACKFORWARD   wkeNavigationType = 2
	WKE_NAVIGATION_TYPE_RELOAD        wkeNavigationType = 3
	WKE_NAVIGATION_TYPE_FORMRESUBMITT wkeNavigationType = 4
	WKE_NAVIGATION_TYPE_OTHER         wkeNavigationType = 5
)

type downInfo struct {
	url      uintptr
	recvSize uint32
	allSize  uint32
}
type wkeNetJobDataBind struct {
	param          uintptr
	recvCallback   wkeNetJobDataRecvCallback
	finishCallback wkeNetJobDataFinishCallback
}
type wkeDownloadBind struct {
	param            uintptr
	recvCallback     uintptr
	finishCallback   uintptr
	saveNameCallback uintptr
}
type wkePopupDialogSaveNameCallback func(ptr, filePath uintptr) uintptr
type wkePaintBitUpdatedCallback func(wke wkeHandle, param, buf uintptr, rect *wkeRect, width, height int32) uintptr
type wkePaintUpdatedCallback func(wke wkeHandle, param, hdc uintptr, left, top, width, height int32) uintptr
type wkeNetResponseCallback func(wke wkeHandle, param, utf8Url uintptr, job wkeNetJob) uintptr
type wkeLoadUrlBeginCallback func(wke wkeHandle, param, utf8Url uintptr, job wkeNetJob) uintptr
type wkeDidCreateScriptContextCallback func(wke wkeHandle, param uintptr, frame wkeFrame, context uintptr, exGroup, worldId int) uintptr
type wkeConsoleCallback func(wke wkeHandle, param uintptr, level int32, msg, name uintptr, line uint32, stack uintptr) uintptr
type wkeLoadUrlEndCallback func(wke wkeHandle, param, url uintptr, job wkeNetJob, buf uintptr, len int32) uintptr
type wkeLoadUrlFailCallback func(wke wkeHandle, param, url uintptr, job wkeNetJob) uintptr
type wkeDocumentReady2Callback func(wke wkeHandle, param uintptr, frame wkeFrame) uintptr
type wkeOnDownloadCallback func(wke wkeHandle, param uintptr, length uint32, url, mime, disposition uintptr, job wkeNetJob, dataBind uintptr) wkeDownloadOpt
type wkeOnAlertBoxCallback func(wke wkeHandle, param uintptr, msg uintptr) uintptr
type wkeOnCreateViewCallback func(wke wkeHandle, param uintptr, naviType wkeNavigationType, url uintptr, feature uintptr) uintptr
type wkeNetJobDataFinishCallback func(param uintptr, job wkeNetJob, result wkeLoadingResult) uintptr
type wkeNetJobDataRecvCallback func(param uintptr, job wkeNetJob, data uintptr, length int32) uintptr
type mbOnBlinkThreadInitCallback func(param uintptr) uintptr
