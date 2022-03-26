package win

import (
	"syscall"
)

func sendMessage(hwnd syscall.Handle, msg uint32, wParam uintptr, lParam uintptr) (int32, error) {
	r0, _, err := syscall.Syscall6(
		user32.MustFindProc("SendMessageW").Addr(),
		4,
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam,
		0,
		0)
	len := int32(r0)
	if len == 0 {
		if err != 0 {
			return 0, error(err)
		} else {
			return len, syscall.EINVAL
		}
	}
	return len, nil
}
