package aliases

import (
	"fmt"
	"testing"
)

func TestCreateAliasList(t *testing.T) {
	// Test the creation of a Alias List object
	y := NewAliasList()
	if fmt.Sprint(y.Values()) != "[]" {
		t.Error("Test failed")
	}
}

func TestGenerateAliasList(t *testing.T) {
	// Test the generation of a Alias List object
	y := NewAliasList()
	y.Generate("loc")
	if len(y.Values()) == 0 {
		t.Error("Test failed")
	}
}

func TestValueOfAliasList(t *testing.T) {
	// Test the values of a Alias List object
	const testVal = "@ssc.loc"
	y := NewAliasList()
	y.Generate("loc")
	if y.value[0] != testVal {
		t.Error("Test failed")
	}
}