//go:build darwin

package displayindex

/*
#cgo darwin LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

int getDisplayIndex() {
    NSPoint mouseLoc = [NSEvent mouseLocation];
    NSArray *screens = [NSScreen screens];

    for (NSUInteger i = 0; i < [screens count]; i++) {
        NSRect frame = [[screens objectAtIndex:i] frame];
        if (NSMouseInRect(mouseLoc, frame, NO)) {
            return (int)i;
        }
    }
    return -1;
}
*/
import "C"
import "fmt"

func CurrentDisplayIndex() (int, error) {
	index := C.getDisplayIndex()
	if index == -1 {
		return -1, fmt.Errorf("cursor not on any active display")
	}
	return int(index), nil
}
