package displayindex

import (
	"github.com/kbinani/screenshot"
	"testing"
)

// If someone has a better idea for testing this, please let me know :) - json

func TestCurrentDisplayIndexBasic(t *testing.T) {
	// This test works with the real system state
	index, err := CurrentDisplayIndex()

	// Handle different environment possibilities
	activeDisplays := screenshot.NumActiveDisplays()

	switch {
	case activeDisplays == 0:
		if err == nil {
			t.Error("Expected error when no displays are active")
		}
	case index < 0 && activeDisplays > 0:
		t.Errorf("Unexpected negative index with %d active displays", activeDisplays)
	case index >= activeDisplays:
		t.Errorf("Index %d exceeds active displays count %d", index, activeDisplays)
	}
}

func TestConsistency(t *testing.T) {
	// Test that subsequent calls return consistent results
	index1, err1 := CurrentDisplayIndex()
	index2, err2 := CurrentDisplayIndex()

	if (err1 != nil) != (err2 != nil) {
		t.Errorf("Inconsistent error results: %v vs %v", err1, err2)
	}

	if index1 != index2 {
		t.Errorf("Inconsistent indices: %d vs %d", index1, index2)
	}
}

func TestErrorConditions(t *testing.T) {
	// Test error message format
	_, err := CurrentDisplayIndex()
	activeDisplays := screenshot.NumActiveDisplays()

	if activeDisplays > 0 && err != nil {
		t.Errorf("Unexpected error with active displays: %v", err)
	}

	if activeDisplays == 0 && err == nil {
		t.Error("Expected error when no displays are active")
	}
}
