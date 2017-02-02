package version

import (
	"fmt"
	"testing"
)

const TESTVERSION = "7.0.0"
const LEGACYVERSION = "4.0.0"

func TestCreateNewVersion(t *testing.T) {
	// Test the creation of version object
	y := NewDrushVersion(TESTVERSION)
	if fmt.Sprint(y.version) != TESTVERSION {
		t.Error("Test failed")
	}
}

func TestDrushExists(t *testing.T) {
	// Test if a Command version is available for installation
	y := NewDrushVersion(TESTVERSION)
	if y.Exists() == false {
		t.Error("Test failed")
	}
}
func TestDrushStatus(t *testing.T) {
	// Test if a Command version is installed
	y := NewDrushVersion(TESTVERSION)
	if y.Status() == false {
		t.Error("Test failed")
	}
}

func TestDrushInstall(t *testing.T) {
	// Test if a version of Command can be installed.
	y := NewDrushVersion(TESTVERSION)
	y.Install()
	if y.Status() == false {
		t.Error("Test failed")
	}
	y.Uninstall()
}

func TestDrushUninstall(t *testing.T) {
	// Test if a version of Command can be uninstalled.
	y := NewDrushVersion(TESTVERSION)
	y.Uninstall()
	if y.Status() == true {
		t.Error("Test failed")
	}
}

func TestDrushReinstall(t *testing.T) {
	// Test if a version of Command can be reinstalled.
	y := NewDrushVersion(TESTVERSION)
	y.Uninstall()
	if y.Status() == true {
		t.Error("Test failed")
	}
	y.Install()
	if y.Status() == false {
		t.Error("Test failed")
	}
}

func TestDrushLegacyInstall(t *testing.T) {
	// Test if a legacy (non-composer) Command install can execute.
	y := NewDrushVersion(LEGACYVERSION)
	if y.Exists() == true {
		y.Install()
		if y.Status() == true {
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

	// Create objects
	x := NewDrushVersion(NEWVERSION)
	y := NewDrushVersion(ACTIVEVERSION)

	if NEWVERSION == ACTIVEVERSION {
		if NEWVERSION == "7.0.0" {
			NEWVERSION = "7.2.0"
		} else {
			NEWVERSION = "7.0.0"
		}
		VERSIONCHANGED = true
	}

	x.Install()
	y.Install()
	x.SetDefault()

	if GetActiveVersion() == NEWVERSION {
		t.Error("Test failed")
	} else {
		if VERSIONCHANGED != false {
			y.SetDefault()
			x.Uninstall()
		}
	}
}
