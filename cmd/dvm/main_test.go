// Test file attempts to identify configuration
// issues ahead of package tests and execution.
package main

import (
	"github.com/fubarhouse/dvm/version"
	"os"
	"os/exec"
	"os/user"
	"testing"
)

const DRUSHPATH = "/usr/local/bin/drush"

type testResult struct {
	name   string
	result bool
}

func (testResult *testResult) Pass() {
	testResult.result = true
}

func (testResult *testResult) Fail() {
	testResult.result = false
}

func (testResult *testResult) Finalize(t *testing.T) {
	if !testResult.result {
		t.FailNow()
	}
}

func TestDrushHomePathExists(t *testing.T) {
	test := &testResult{"Test drush home directory exists", true}
	User, userErr := user.Current()
	_, statErr := os.Stat(User.HomeDir + "/.drush")
	if userErr != nil || statErr != nil {
		t.Fail()
	}
	test.Finalize(t)
}

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

//x

//func TestComposerCommand(t *testing.T) {
//	test := &testResult{"Test composer command execution", true}
//	_, err := exec.Command("composer", "--help").Output()
//	if err != nil {
//		test.Fail()
//	}
//	test.Finalize(t)
//}

func TestUnzipCommand(t *testing.T) {
	test := &testResult{"Test unzip command execution", true}
	_, err := exec.Command("unzip", "--help").Output()
	if err != nil {
		test.Fail()
	}
	test.Finalize(t)
}

func TestWgetCommand(t *testing.T) {
	test := &testResult{"Test wget command execution", true}
	_, err := exec.Command("wget", "--help").Output()
	if err != nil {
		test.Fail()
	}
	test.Finalize(t)
}
