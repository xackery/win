package win

import (
	"fmt"
	"testing"
	"time"
)

func TestClickButton(t *testing.T) {
	hwnd, err := WindowButtonByTitleAndCaption("About helloworld-cpp", "OK")
	if err != nil {
		t.Fatalf("button: %s", err)
	}
	fmt.Printf("found 0x%x\n", hwnd)
	if !IsWindowActive(hwnd) {
		parentHwnd, err := Owner(hwnd)
		if err != nil {
			t.Fatalf("owner: %s", err)
		}

		SetActiveWindow(parentHwnd)
	}
	ClickButton(hwnd)
	time.Sleep(1 * time.Second)
	hwnd, err = WindowButtonByTitleAndCaption("About helloworld-cpp", "OK")
	if err != nil {
		t.Fatalf("button2: %s", err)
	}
	fmt.Printf("found 0x%x\n", hwnd)
	if hwnd != 0 {
		t.Fatalf("button not clicked")
	}
}
