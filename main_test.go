package main

import (
	"fmt"
	"testing"
)

const TESTVERSION = "7.0.0"
const LEGACYVERSION = "4.0.0"

func TestCreateNewVersion(t *testing.T) {
	// Test the creation of version object
	y := newDrushVersion(TESTVERSION)
	if fmt.Sprint(y.version) != TESTVERSION {
		t.Error("Test failed")
	}
}

func TestCreateNewVersionList(t *testing.T) {
	// Test the creation of version list object
	y := newDrushVersionList()
	if fmt.Sprint(y) != "{[]}" {
		t.Error("Test failed")
	}
}

func TestCreateNewPackage(t *testing.T) {
	// Test the creation of package object
	y := newDrushPackage("registry_rebuild")
	if fmt.Sprint(y.name) != "registry_rebuild" {
		t.Error("Test failed")
	}
}

func TestUninstallPackage(t *testing.T) {
	// Test the uninstallation of a drush package
	y := newDrushPackage("registry_rebuild")
	y.Uninstall()
	if y.status == true {
		t.Error("Test failed")
	}
}

func TestInstallPackage(t *testing.T) {
	// Test the installation of a drush package
	y := newDrushPackage("registry_rebuild")
	y.Install()
	if y.status == false {
		t.Error("Test failed")
	}
}

func TestReinstallPackage(t *testing.T) {
	// Test the reinstallation of a drush package
	y := newDrushPackage("registry_rebuild")
	y.Reinstall()
	if y.status == false {
		t.Error("Test failed")
	}
}

func TestListPackage(t *testing.T) {
	// Test the 'list installed' feature of a drush package object
	y := newDrushPackage("registry_rebuild")
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

func TestDrushExists(t *testing.T) {
	// Test if a Drush version is available for installation
	y := newDrushVersion(TESTVERSION)
	if y.Exists() == false {
		t.Error("Test failed")
	}
}
func TestDrushStatus(t *testing.T) {
	// Test if a Drush version is installed
	y := newDrushVersion(TESTVERSION)
	if y.Status() == false {
		t.Error("Test failed")
	}
}

func TestDrushInstall(t *testing.T) {
	// Test if a version of Drush can be installed.
	y := newDrushVersion(TESTVERSION)
	y.Install()
	if y.Status() == false {
		t.Error("Test failed")
	}
	y.Uninstall()
}

func TestDrushUninstall(t *testing.T) {
	// Test if a version of Drush can be uninstalled.
	y := newDrushVersion(TESTVERSION)
	y.Uninstall()
	if y.Status() == true {
		t.Error("Test failed")
	}
}

func TestDrushReinstall(t *testing.T) {
	// Test if a version of Drush can be reinstalled.
	y := newDrushVersion(TESTVERSION)
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
	// Test if a legacy (non-composer) Drush install can execute.
	y := newDrushVersion(LEGACYVERSION)
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
	// Test if a default Drush version can be used.

	// Set-up constants for use, they cannot be
	// constants in case we need to reassign them.
	ACTIVEVERSION := getActiveVersion()
	NEWVERSION := TESTVERSION
	VERSIONCHANGED := false

	// Create objects
	x := newDrushVersion(NEWVERSION)
	y := newDrushVersion(ACTIVEVERSION)

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

	if getActiveVersion() == NEWVERSION {
		t.Error("Test failed")
	} else {
		if VERSIONCHANGED != false {
			y.SetDefault()
			x.Uninstall()
		}
	}
}