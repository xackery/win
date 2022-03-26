package win

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	windowMessageGetText       = 0x000D // WM_GETTEXT
	windowMessageGetTextLength = 0x000E //WM_GETTEXTLENGTH
)

func controlText(hwnd syscall.Handle) (string, error) {
	size, err := sendMessage(hwnd, windowMessageGetTextLength, 0, 0)
	if err != nil {
		return "", fmt.Errorf("sendMessage getTextLength: %w", err)
	}
	buf := make([]uint16, size+1)
	_, err = sendMessage(hwnd, windowMessageGetText, uintptr(size+1), uintptr(unsafe.Pointer(&buf[0])))
	if err != nil {
		return "", fmt.Errorf("sendMessage getText: %w", err)
	}
	if len(buf) == 0 {
		return "", nil
	}

	//text := strings.Replace(syscall.UTF16ToString(buf), decimalSepS, ".", 1)
	return syscall.UTF16ToString(buf), nil
}

/*
func separator() (string, error) {

	win.GetLocaleInfo(win.LOCALE_USER_DEFAULT, win.LOCALE_SDECIMAL, &buf[0], int32(len(buf)))
		decimalSepB = byte(buf[0])
		decimalSepS = syscall.UTF16ToString(buf[0:1])
}*/
