package version

import (
	"fmt"
	"github.com/fubarhouse/dvm/version"
	"testing"
)

const TESTVERSION = "7.0.0"
const LEGACYVERSION = "4.0.0"

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

func TestCreateNewVersion(t *testing.T) {
	// Test the creation of version object
	y := NewDrushVersion(TESTVERSION)
	if fmt.Sprint(y.version) != TESTVERSION {
		t.Error("Test failed")
	}
}

func TestDrushInstall(t *testing.T) {
	// Test if a version of Command can be installed.
	y := NewDrushVersion(TESTVERSION)
	y.Install()
	y.SetDefault()
	if y.Status() == false {
		t.Error("Test failed")
	}
}

func TestDrushExists(t *testing.T) {
	// Test if a Command version is available for installation
	y := NewDrushVersion(TESTVERSION)
	if !y.Exists() {
		t.Error("Test failed")
	}
}

func TestDrushStatus(t *testing.T) {
	// Test if a Command version is installed
	y := NewDrushVersion(TESTVERSION)
	if !y.Status() {
		t.Error("Test failed")
	}
}

func TestDrushUninstall(t *testing.T) {
	// Test if a version of Command can be uninstalled.
	y := NewDrushVersion(TESTVERSION)
	y.Uninstall()
	if y.Status() {
		t.Error("Test failed")
	}
}

func TestDrushReinstall(t *testing.T) {
	// Test if a version of Command can be reinstalled.
	y := NewDrushVersion(TESTVERSION)
	y.Uninstall()
	if y.Status() {
		t.Error("Test failed")
	}
	y.Install()
	if !y.Status() {
		t.Error("Test failed")
	}
}

func TestDrushLegacyInstall(t *testing.T) {
	// Test if a legacy (non-composer) Command install can execute.
	y := NewDrushVersion(LEGACYVERSION)
	if y.Exists() {
		y.Install()
		if y.Status() {
			y.Uninstall()
		} else {
			t.Error("Test failed")
		}
	} else {
		t.Error("Test failed")
	}
}

func TestDrushSpecifyDefault(t *testing.T) {
	// Test if a default Command version can be used.

	// Set-up constants for use, they cannot be
	// constants in case we need to reassign them.
	ACTIVEVERSION := GetActiveVersion()
	NEWVERSION := TESTVERSION
	VERSIONCHANGED := false

	if NEWVERSION == ACTIVEVERSION {
		if NEWVERSION == "7.0.0" {
			NEWVERSION = "7.2.0"
		} else {
			NEWVERSION = "7.0.0"
		}
		VERSIONCHANGED = true
	}

	// Create objects
	x := NewDrushVersion(NEWVERSION)
	y := NewDrushVersion(ACTIVEVERSION)

	x.Install()
	y.Install()
	x.SetDefault()

	if GetActiveVersion() != NEWVERSION {
		t.Error("Test failed")
	} else {
		if VERSIONCHANGED != false {
			y.SetDefault()
			x.Uninstall()
		}
	}
}
