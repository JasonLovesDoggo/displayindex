//go:build windows || linux

package displayindex

import (
	"fmt"
	"github.com/kbinani/screenshot"
)

// CurrentDisplayIndex returns the display index containing the cursor (0-based)
// Returns -1 and error if cursor not found on any display
func CurrentDisplayIndex() (int, error) {
	x, y, err := getCursorPosition()
	if err != nil {
		return -1, fmt.Errorf("cursor position error: %w", err)
	}

	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		if x >= bounds.Min.X && x < bounds.Max.X &&
			y >= bounds.Min.Y && y < bounds.Max.Y {
			return i, nil
		}
	}

	return -1, fmt.Errorf("cursor not on any active display")
}
