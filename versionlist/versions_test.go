package versionlist

import (
	"testing"
	"fmt"
)

func TestCreateNewVersionList(t *testing.T) {
	// Test the creation of version list object
	y := NewDrushVersionList()
	if fmt.Sprint(y) != "{[]}" {
		t.Error("Test failed")
	}
}
