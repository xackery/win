package win

import "syscall"

var (
	user32 = syscall.MustLoadDLL("user32.dll")
)
