//go:build darwin

package main

/*
#cgo darwin LDFLAGS: -framework CoreGraphics
#include <CoreGraphics/CoreGraphics.h>
*/
import "C"

import "fmt"

func GetCursorPosition() (x, y int, err error) {
	event := C.CGEventCreate(nil)
	if event == nil {
		return 0, 0, fmt.Errorf("CGEventCreate failed")
	}
	defer C.CFRelease(C.CFTypeRef(event))

	point := C.CGEventGetLocation(event)
	return int(point.x), int(point.y), nil
}
