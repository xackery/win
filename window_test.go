package win

import (
	"fmt"
	"testing"
)

func TestWindowByTitle(t *testing.T) {
	hwnd, err := WindowByTitle("helloworld-cpp")
	if err != nil {
		t.Fatalf("WindowByTitle: %s", err)
	}
	if hwnd == 0 {
		t.Fatalf("hwnd zero")
	}
	fmt.Printf("hwnd is 0x%x\n", hwnd)
}

func TestChildWindowByCaption(t *testing.T) {
	hwnd, err := WindowByTitle("About helloworld-cpp")
	if err != nil {
		t.Fatalf("WindowByTitle: %s", err)
	}
	if hwnd == 0 {
		t.Fatalf("hwnd zero")
	}
	hwnd, err = ChildWindowByCaption(hwnd, "OK")
	if err != nil {
		t.Fatalf("ChildWindowByCaption: %s", err)
	}
	t.Fatalf("got hwnd: 0x%x", hwnd)
}

func TestClassName(t *testing.T) {
	hwnd, err := WindowByTitle("helloworld-cpp")
	if err != nil {
		t.Fatalf("WindowByTitle: %s", err)
	}
	if hwnd == 0 {
		t.Fatalf("hwnd zero")
	}
	name, err := ClassName(hwnd)
	if err != nil {
		t.Fatalf("ClassName: %s", err)
	}
	t.Fatalf("got className: %s", name)
}

func TestClassNameChild(t *testing.T) {
	hwnd, err := WindowByTitle("About helloworld-cpp")
	if err != nil {
		t.Fatalf("WindowByTitle: %s", err)
	}
	if hwnd == 0 {
		t.Fatalf("hwnd zero")
	}
	_, err = ClassName(hwnd)
	if err != nil {
		t.Fatalf("ClassName: %s", err)
	}
	hwnd, err = ChildWindowByCaption(hwnd, "OK")
	if err != nil {
		t.Fatalf("ChildWindowByCaption: %s", err)
	}

	name, err := ClassName(hwnd)
	if err != nil {
		t.Fatalf("ClassName: %s", err)
	}
	t.Fatalf("got className: %s", name)
}
