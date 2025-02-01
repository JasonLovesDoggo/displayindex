//go:build windows

package displayindex

import (
	"fmt"
	"syscall"
	"unsafe"
)

type POINT struct {
	X, Y int32
}

func getCursorPosition() (x, y int, err error) {
	user32 := syscall.NewLazyDLL("user32.dll")
	getCursorPos := user32.NewProc("GetCursorPos")
	pt := POINT{}
	ret, _, _ := getCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return 0, 0, fmt.Errorf("GetCursorPos failed")
	}
	return int(pt.X), int(pt.Y), nil
}
