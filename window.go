package win

import (
	"fmt"
	"syscall"
	"unsafe"
)

// IsActiveWindow returns true if provided hwnd is currently in foreground
func IsWindowActive(hwnd syscall.Handle) bool {
	fgWnd, err := foregroundWindow()
	if err != nil {
		return false
		//return false, fmt.Errorf("foregroundWindow: %w", err)
	}
	return hwnd == fgWnd
}

// SetActiveWindow makes a window go to foreground
func SetActiveWindow(hwnd syscall.Handle) error {
	_, _, err := syscall.Syscall(user32.MustFindProc("SetActiveWindow").Addr(), 1, uintptr(hwnd), 0, 0)
	if err != 0 {
		return error(err)
	}
	return nil
}

// WindowButtonByTitleAndCaption simplifies finding the hwnd of a button by providing with main window title, and button caption
func WindowButtonByTitleAndCaption(windowTitle string, buttonCaption string) (syscall.Handle, error) {
	hwnd, err := WindowByTitle(windowTitle)
	if err != nil {
		return 0, fmt.Errorf("windowByTitle %s: %w", windowTitle, err)
	}

	if hwnd == 0 {
		return 0, fmt.Errorf("window not found")
	}

	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		src, err := controlText(h)
		if err != nil {
			return 1
		}
		if src != buttonCaption {
			return 1
		}
		className, err := ClassName(h)
		if err != nil {
			return 1
		}
		if className != "Button" {
			return 1
		}
		hwnd = h
		return 0
	})
	enumChildWindows(hwnd, cb, 0)

	if hwnd == 0 {
		return hwnd, fmt.Errorf("child window not found")
	}

	return hwnd, nil
}

// WindowByTitle iterates all windows looking for one matching title
func WindowByTitle(title string) (syscall.Handle, error) {
	var hwnd syscall.Handle
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		b := make([]uint16, 200)
		_, err := windowText(h, &b[0], int32(len(b)))
		if err != nil {
			return 1
		}
		if syscall.UTF16ToString(b) != title {
			return 1
		}
		hwnd = h
		return 0
	})
	enumWindows(cb, 0)
	if hwnd == 0 {
		return 0, fmt.Errorf("not found")
	}
	return hwnd, nil
}

// ChildWindowByCaption enumerates windows that are children of provided hwnd, looking for caption
func ChildWindowByCaption(hwnd syscall.Handle, caption string) (syscall.Handle, error) {
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		src, err := controlText(h)
		if err != nil {
			return 1
		}
		if src != caption {
			return 1
		}
		hwnd = h
		return 0
	})
	enumChildWindows(hwnd, cb, 0)
	if hwnd == 0 {
		return 0, fmt.Errorf("not found")
	}
	return hwnd, nil
}

// ClassName returns the class name of a window, assisting in identifying what behaviors it has
func ClassName(hwnd syscall.Handle) (string, error) {
	b := make([]uint16, 200)
	r0, _, err := syscall.Syscall(user32.MustFindProc("GetClassNameW").Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(&b[0])), uintptr(200))
	len := int32(r0)
	if len == 0 {
		if err != 0 {
			return "", error(err)
		}
		return "", syscall.EINVAL
	}
	src := syscall.UTF16ToString(b)
	if src == "#32770" {
		src = "Dialog"
	}
	return src, nil
}

// Owner calls GetWindow with the GW_OWNER flag to return the parent of a hwnd
// This is useful for cases like a popup dialog
func Owner(hwnd syscall.Handle) (syscall.Handle, error) {
	resp, _, err := syscall.Syscall(user32.MustFindProc("GetOwner").Addr(), 1, uintptr(hwnd), 0, 0)
	if err != 0 {
		return 0, error(err)
	}
	return syscall.Handle(resp), nil
}

func enumWindows(enumFunc uintptr, lparam uintptr) error {
	r1, _, err := syscall.Syscall(user32.MustFindProc("EnumWindows").Addr(), 2, uintptr(enumFunc), uintptr(lparam), 0)
	if r1 == 0 {
		if err != 0 {
			return error(err)
		}
		return syscall.EINVAL
	}
	return nil
}

func windowText(hwnd syscall.Handle, str *uint16, maxCount int32) (int32, error) {
	r0, _, err := syscall.Syscall(user32.MustFindProc("GetWindowTextW").Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
	len := int32(r0)
	if len == 0 {
		if err != 0 {
			return 0, error(err)
		}
		return len, syscall.EINVAL
	}
	return len, nil
}

func enumChildWindows(hwnd syscall.Handle, enumFunc uintptr, lparam uintptr) error {
	r1, _, err := syscall.Syscall(user32.MustFindProc("EnumChildWindows").Addr(), 3, uintptr(hwnd), uintptr(enumFunc), uintptr(lparam))
	if r1 == 0 {
		if err != 0 {
			return error(err)
		}
		return syscall.EINVAL
	}
	return nil
}

func foregroundWindow() (syscall.Handle, error) {
	hwnd, _, err := syscall.Syscall(user32.MustFindProc("GetForegroundWindow").Addr(), 0, 0, 0, 0)
	if err != 0 {
		return 0, error(err)
	}
	return syscall.Handle(hwnd), nil
}
