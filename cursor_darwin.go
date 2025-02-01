//go:build darwin

package displayindex

/*
#cgo darwin LDFLAGS: -framework CoreGraphics -framework Foundation
#include <CoreGraphics/CoreGraphics.h>

int getCurrentDisplayIndex() {
    CGEventRef event = CGEventCreate(nil);
    CGPoint location = CGEventGetLocation(event);
    CFRelease(event);

    CGDirectDisplayID displays[16];
    uint32_t displayCount = 0;

    CGError err = CGGetActiveDisplayList(16, displays, &displayCount);
    if (err != kCGErrorSuccess || displayCount == 0) {
        return -1;
    }

    for (uint32_t i = 0; i < displayCount; i++) {
        CGRect bounds = CGDisplayBounds(displays[i]);
        if (CGRectContainsPoint(bounds, location)) {
            return (int)i;
        }
    }

    return -1;
}
*/
import "C"

import "fmt"

func CurrentDisplayIndex() (int, error) {
	index := C.getCurrentDisplayIndex()
	if index == -1 {
		return -1, fmt.Errorf("cursor not on any active display")
	}
	return int(index), nil
}
