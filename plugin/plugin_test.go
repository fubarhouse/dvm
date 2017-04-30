package plugin

import (
	"fmt"
	"github.com/fubarhouse/dvm/version"
	"testing"
)

var initState = false

func TestDrushInstallDefaultTestCase(t *testing.T) {
	Version := "7.0.0"
	// Ensure a version is tested to prevent failure.
	y := version.NewDrushVersion(Version)
	y.Install()
	y.SetDefault()
	if !y.Status() {
		t.Error("Test failed")
	}
}

func TestInitiate(t *testing.T) {
	y := NewDrushPackage("registry_rebuild")
	if y.Status() {
		initState = true
	}
}

func TestCreateNewPackage(t *testing.T) {
	// Test the creation of package object
	y := NewDrushPackage("registry_rebuild")
	if fmt.Sprint(y.name) != "registry_rebuild" {
		t.Error("Test failed")
	}
}

func TestInstallPackage(t *testing.T) {
	// Test the installation of a drush package
	y := NewDrushPackage("registry_rebuild")
	y.Install()
	if y.status == false {
		t.Error("Test failed")
	}
}

func TestReinstallPackage(t *testing.T) {
	// Test the reinstallation of a drush package
	y := NewDrushPackage("registry_rebuild")
	y.Reinstall()
	if y.status == false {
		t.Error("Test failed")
	}
}

func TestListPackage(t *testing.T) {
	// Test the 'list installed' feature of a drush package object
	y := NewDrushPackage("registry_rebuild")
	x := y.List()
	foundPackage := false
	for index := range x {
		if x[index] == y.name {
			foundPackage = true
		}
	}
	if foundPackage != true {
		t.Error("Test failed")
	}
}

func TestUninstallPackage(t *testing.T) {
	// Test the uninstallation of a drush package
	y := NewDrushPackage("registry_rebuild")
	y.Uninstall()
	if y.status == true {
		t.Error("Test failed")
	}
}
