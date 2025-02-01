//go:build linux

package displayindex

/*
#cgo linux LDFLAGS: -lX11
#include <X11/Xlib.h>
*/
import "C"
import (
	"fmt"
)

func getCursorPosition() (x, y int, err error) {
	display := C.XOpenDisplay(nil)
	if display == nil {
		return 0, 0, fmt.Errorf("XOpenDisplay failed")
	}
	defer C.XCloseDisplay(display)

	var rootWindow, childWindow C.Window
	var rootX, rootY, winX, winY C.int
	var mask C.uint

	ret := C.XQueryPointer(
		display,
		C.XDefaultRootWindow(display),
		&rootWindow,
		&childWindow,
		&rootX,
		&rootY,
		&winX,
		&winY,
		&mask,
	)
	if ret == 0 {
		return 0, 0, fmt.Errorf("XQueryPointer failed")
	}

	return int(rootX), int(rootY), nil
}
