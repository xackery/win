package win

import (
	"fmt"
	"syscall"
)

// ClickButton sends a BM_CLICK to a target hwnd
func ClickButton(hwnd syscall.Handle) error {
	_, err := sendMessage(hwnd, 0x00F5, 0, 0)
	if err != nil {
		return fmt.Errorf("SendMessage: %w", err)
	}
	return nil
}
