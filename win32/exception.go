package win32

import (
	"github.com/Mengdch/browser/log"
	"github.com/Mengdch/win"
	"golang.org/x/sys/windows"
	"github.com/Mengdch/goUtil/TypeTools"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

const (
	PROCESS_CREATE_PROCESS            = 0x0080
	PROCESS_CREATE_THREAD             = 0x0002
	PROCESS_DUP_HANDLE                = 0x0040
	PROCESS_QUERY_INFORMATION         = 0x0400
	PROCESS_QUERY_LIMITED_INFORMATION = 0x1000
	PROCESS_SET_INFORMATION           = 0x0200
	PROCESS_SET_QUOTA                 = 0x0100
	PROCESS_SUSPEND_RESUME            = 0x0800
	PROCESS_TERMINATE                 = 0x0001
	PROCESS_VM_OPERATION              = 0x0008
	PROCESS_VM_READ                   = 0x0010
	PROCESS_VM_WRITE                  = 0x0020

	PROCESS_ALL_ACCESS = (PROCESS_CREATE_PROCESS | PROCESS_CREATE_THREAD | PROCESS_DUP_HANDLE | PROCESS_QUERY_INFORMATION | PROCESS_QUERY_LIMITED_INFORMATION | PROCESS_SET_INFORMATION | PROCESS_SET_QUOTA | PROCESS_SUSPEND_RESUME | PROCESS_TERMINATE | PROCESS_VM_OPERATION | PROCESS_VM_WRITE | PROCESS_VM_READ)

	GENERIC_WRITE         = 0x40000000
	FILE_SHARE_WRITE      = 0x00000002
	CREATE_ALWAYS         = 0x2
	FILE_ATTRIBUTE_NORMAL = 0x80

	MiniDumpNormal                         = 0x00000000
	MiniDumpWithDataSegs                   = 0x00000001
	MiniDumpWithFullMemory                 = 0x00000002
	MiniDumpWithHandleData                 = 0x00000004
	MiniDumpFilterMemory                   = 0x00000008
	MiniDumpScanMemory                     = 0x00000010
	MiniDumpWithUnloadedModules            = 0x00000020
	MiniDumpWithIndirectlyReferencedMemory = 0x00000040
	MiniDumpFilterModulePaths              = 0x00000080
	MiniDumpWithProcessThreadData          = 0x00000100
	MiniDumpWithPrivateReadWriteMemory     = 0x00000200
	MiniDumpWithoutOptionalData            = 0x00000400
	MiniDumpWithFullMemoryInfo             = 0x00000800
	MiniDumpWithThreadInfo                 = 0x00001000
	MiniDumpWithCodeSegs                   = 0x00002000
	MiniDumpWithoutAuxiliaryState          = 0x00004000
	MiniDumpWithFullAuxiliaryState         = 0x00008000
	MiniDumpWithPrivateWriteCopyMemory     = 0x00010000
	MiniDumpIgnoreInaccessibleMemory       = 0x00020000
	MiniDumpWithTokenInformation           = 0x00040000
	MiniDumpWithModuleHeaders              = 0x00080000
	MiniDumpFilterTriage                   = 0x00100000
	MiniDumpWithAvxXStateContext           = 0x00200000
	MiniDumpWithIptTrace                   = 0x00400000
	MiniDumpValidTypeFlags                 = 0x007fffff
)

var pid uint32
var miniDumpWriteDump *windows.LazyProc

func init() {
	pid = windows.GetCurrentProcessId()
	dbghelp := windows.NewLazySystemDLL("Dbghelp.dll")
	miniDumpWriteDump = dbghelp.NewProc("MiniDumpWriteDump")
	win.SetUnhandledExceptionFilter(onException)
}

type exceptionInfo struct {
	ThreadId          uint32
	ExceptionPointers uintptr
	ClientPointers    uint32
}

func onException(param uintptr) uintptr {
	if miniDumpWriteDump == nil {
		logRecord("onException", "miniDumpWriteDump nil")
		return 0
	}
	pHandle := windows.CurrentProcess()
	name := "thublink" + strconv.FormatInt(time.Now().Unix(), 32) + ".dmp"
	fromString, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		logRecord("onException.UTF16PtrFromString:"+name, err.Error())
		return 0
	}
	var sa windows.SecurityAttributes
	fHandle, err := windows.CreateFile(fromString, GENERIC_WRITE, FILE_SHARE_WRITE, &sa, CREATE_ALWAYS, FILE_ATTRIBUTE_NORMAL, 0)
	if err != nil {
		logRecord("onException.CreateFile:"+name, err.Error())
		return 0
	}
	var info exceptionInfo
	info.ExceptionPointers = param
	info.ThreadId = windows.GetCurrentThreadId()
	success, _, err := miniDumpWriteDump.Call(uintptr(pHandle), uintptr(pid), uintptr(fHandle), MiniDumpNormal, uintptr(unsafe.Pointer(&info)), 0, 0)
	windows.CloseHandle(fHandle)
	if success != 1 {
		logRecord("onException.miniDumpWriteDump", err.Error())
		return 0
	}
	url := log.Upload(name)
	if len(url) > 0 {
		logRecord(TypeTools.OutJson(map[string]interface{}{"dmp": url, "url": urls}), "dmpHandle")
	}
	return 0
}

func GetFocus() win.HWND {
	hForeWnd := win.GetForegroundWindow()
	dwForeID := win.GetWindowThreadProcessId(hForeWnd, nil)
	dwCurID := win.GetCurrentThreadId()
	win.AttachThreadInput(int32(dwForeID), int32(dwCurID), true)
	wnd := win.GetFocus()
	win.AttachThreadInput(int32(dwForeID), int32(dwCurID), false)
	return wnd
}
func SetTop(hWnd win.HWND) {
	hForeWnd := win.GetForegroundWindow()
	dwForeID := win.GetWindowThreadProcessId(hForeWnd, nil)
	dwCurID := win.GetCurrentThreadId()
	win.AttachThreadInput(int32(dwCurID), int32(dwForeID), true)
	win.ShowWindow(hWnd, win.SW_SHOWNORMAL)
	win.SetWindowPos(hWnd, win.HWND_TOPMOST, 0, 0, 0, 0, win.SWP_NOSIZE|win.SWP_NOMOVE)
	win.SetWindowPos(hWnd, win.HWND_NOTOPMOST, 0, 0, 0, 0, win.SWP_NOSIZE|win.SWP_NOMOVE)
	win.SetForegroundWindow(hWnd)
	win.AttachThreadInput(int32(dwCurID), int32(dwForeID), false)
}
func logRecord(value, error string) {
	defer log.CatchPanic("logRecord")
	log.Log(value, error)
}
