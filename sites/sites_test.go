package sites

import (
	"testing"
)

func TestCreateSiteList(t *testing.T) {
	// Test the creation of package object
	y := NewSiteList()
	y.SetKey("test value")
	if y.key != "test value" {
		t.Error("Test failed")
	}
}

func TestCreateSiteListValue(t *testing.T) {
	// Test the creation of package object
	y := NewSiteList()
	y.SetKey(".prod")
	if y.Count() == 0 {
		t.Error("Test failed")
	}
}
