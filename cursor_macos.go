//go:build darwin

package displayindex

/*
#cgo darwin LDFLAGS: -framework CoreGraphics
#include <CoreGraphics/CoreGraphics.h>
*/
import "C"
import "fmt"

func getCursorPosition() (x, y int, err error) {
	// Create event source first
	source := C.CGEventSourceCreate(C.kCGEventSourceStateHIDSystemState)
	if source == nil {
		return 0, 0, fmt.Errorf("CGEventSourceCreate failed")
	}
	defer C.CFRelease(C.CFTypeRef(source))

	event := C.CGEventCreate(source)
	if event == nil {
		return 0, 0, fmt.Errorf("CGEventCreate failed")
	}
	defer C.CFRelease(C.CFTypeRef(event))

	point := C.CGEventGetLocation(event)
	return int(point.x), int(point.y), nil
}
