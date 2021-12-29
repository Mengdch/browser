package win32

import (
	"golang.org/x/sys/windows/registry"
	"syscall"
)

const (
	// Registry key security and access rights.
	// See https://msdn.microsoft.com/en-us/library/windows/desktop/ms724878.aspx
	// for details.
	ALL_ACCESS         = 0xf003f
	CREATE_LINK        = 0x00020
	CREATE_SUB_KEY     = 0x00004
	ENUMERATE_SUB_KEYS = 0x00008
	EXECUTE            = 0x20019
	NOTIFY             = 0x00010
	QUERY_VALUE        = 0x00001
	READ               = 0x20019
	SET_VALUE          = 0x00002
	WOW64_32KEY        = 0x00200
	WOW64_64KEY        = 0x00100
	WRITE              = 0x20006

	_ERROR_NO_MORE_ITEMS syscall.Errno = 259
)

// Key is a handle to an open Windows registry key.
// Keys can be obtained by calling OpenKey; there are
// also some predefined root keys such as CURRENT_USER.
// Keys can be used directly in the Windows API.
type Key syscall.Handle

const (
	// Windows defines some predefined root keys that are always open.
	// An application can use these keys as entry points to the registry.
	// Normally these keys are used in OpenKey to open new keys,
	// but they can also be used anywhere a Key is required.
	CLASSES_ROOT     = Key(syscall.HKEY_CLASSES_ROOT)
	CURRENT_USER     = Key(syscall.HKEY_CURRENT_USER)
	LOCAL_MACHINE    = Key(syscall.HKEY_LOCAL_MACHINE)
	USERS            = Key(syscall.HKEY_USERS)
	CURRENT_CONFIG   = Key(syscall.HKEY_CURRENT_CONFIG)
	PERFORMANCE_DATA = Key(syscall.HKEY_PERFORMANCE_DATA)
)

// Close closes open key k.
func (k Key) Close() error {
	return syscall.RegCloseKey(syscall.Handle(k))
}

// OpenKey opens a new key with path name relative to key k.
// It accepts any open key, including CURRENT_USER and others,
// and returns the new key and an error.
// The access parameter specifies desired access rights to the
// key to be opened.
func OpenKey(k Key, path string, access uint32) (Key, error) {
	p, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return 0, err
	}
	var subKey syscall.Handle
	err = syscall.RegOpenKeyEx(syscall.Handle(k), p, 0, access, &subKey)
	if err != nil {
		return 0, err
	}
	return Key(subKey), nil
}

func (k Key) HaveSubKey(name string) (bool, error) {
	p, err := syscall.UTF16FromString(name)
	if err != nil {
		return false, err
	}
	err = syscall.RegQueryValueEx(syscall.Handle(k), &p[0], nil, nil, nil, nil)
	return err == nil, err

}
func checkProtocol(scheme string) (bool, error) {
	key, err := OpenKey(CLASSES_ROOT, scheme, registry.READ)
	if err != nil {
		return false, err
	}
	defer key.Close()
	return key.HaveSubKey(scheme)
}
