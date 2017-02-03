package alias

// TODO: Tests are failing... Make em' pass!

import (
	"testing"
)

func TestCreateAlias(t *testing.T) {
	// Test the creation of a Alias object
	y := NewAlias("Test", "/usr/local/bin/Test", "TestAlias.dev")
	if y.GetName() != "Test" {
		t.Error("Test failed")
	}
	if y.GetPath() != "/usr/local/bin/Test" {
		t.Error("Test failed")
	}
	if y.GetUri() != "TestAlias.dev" {
		t.Error("Test failed")
	}
}

func TestCreateAliasSetters(t *testing.T) {
	// Test the setter methods of an Alias object
	y := NewAlias("Test", "/usr/local/bin/Test", "TestAlias.dev")
	y.SetName("Test")
	y.SetPath("/usr/local/bin/Test")
	y.SetUri("TestAlias.dev")
	if y.GetName() != "Test" {
		t.Error("Test failed")
	}
	if y.GetPath() != "/usr/local/bin/Test" {
		t.Error("Test failed")
	}
	if y.GetUri() != "TestAlias.dev" {
		t.Error("Test failed")
	}
}
