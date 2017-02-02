package versionlist

import (
	"io/ioutil"
	"strings"
	"fmt"
	"os/user"
	"os/exec"
)

const PATH_COMPOSER = "/usr/local/bin/composer"

type DrushVersionList struct {
	// A struct to store associated versions in a simple []string.
	// This is used by methods to store and use multiple version data.
	list []string
}

func NewDrushVersionList() DrushVersionList {
	// An API to create/store a Command version list object.
	retVal := DrushVersionList{}
	return retVal
}

func (drushVersionList *DrushVersionList) ListContents() []string {
	// Return a list of all local versions of Command.
	// This is a manually updated array (for performance sake)
	// which stores all valid Command versions for testing.
	return drushVersionList.list
}

func (drushVersionList *DrushVersionList) ListLocal() {
	// Return a list of all local versions of Command.
	// This is a manually updated array (for performance sake)
	// which stores all valid Command versions for testing.
	drushVersionList.list = []string{"1.0.0+drupal5", "1.0.0+drupal6", "1.0.0-beta1+drupal5", "1.0.0-beta2+drupal5", "1.0.0-beta3+drupal5", "1.0.0-beta4+drupal5", "1.0.0-rc1+drupal5", "1.0.0-rc1+drupal6", "1.0.0-rc11+drupal6", "1.0.0-rc2+drupal5", "1.0.0-rc2+drupal6", "1.0.0-rc2+drupal7", "1.0.0-rc3+drupal5", "1.1.0+drupal5", "1.1.0+drupal6", "1.2.0+drupal5", "1.2.0+drupal6", "1.3.0+drupal5", "1.4.0+drupal5", "2.0.0", "2.0.0-alpha1+drupal5", "2.0.0-alpha1+drupal6", "2.0.0-alpha1+drupal7", "2.0.0-alpha2+drupal5", "2.0.0-alpha2+drupal6", "2.0.0-alpha2+drupal7", "2.0.0-rc1", "2.1.0", "3.0.0", "3.0.0-alpha1", "3.0.0-beta1", "3.0.0-rc1", "3.0.0-rc2", "3.0.0-rc3", "3.0.0-rc4", "3.1.0", "3.2.0", "3.3.0", "4.0.0", "4.0.0-rc1", "4.0.0-rc10", "4.0.0-rc3", "4.0.0-rc4", "4.0.0-rc5", "4.0.0-rc6", "4.0.0-rc7", "4.0.0-rc8", "4.0.0-rc9", "4.1.0", "4.2.0", "4.3.0", "4.4.0", "4.5.0", "4.5.0-rc1", "4.6.0", "5.0.0", "5.0.0-rc1", "5.0.0-rc2", "5.0.0-rc3", "5.0.0-rc4", "5.0.0-rc5", "5.1.0", "5.2.0", "5.3.0", "5.4.0", "5.5.0", "5.6.0", "5.7.0", "5.8.0", "5.9.0", "6.0.0-rc1", "6.0.0-rc2", "6.0.0-rc3", "6.0.0-rc4", "6.1.0-rc1", "6.0.0", "6.1.0", "6.2.0", "6.3.0", "6.4.0", "6.5.0", "6.6.0", "7.0.0-alpha1", "7.0.0-alpha2", "7.0.0-alpha3", "7.0.0-alpha4", "7.0.0-alpha5", "7.0.0-alpha6", "7.0.0-alpha7", "7.0.0-alpha8", "7.0.0-alpha9", "7.0.0-rc1", "7.0.0-rc2", "7.0.0", "7.1.0", "7.2.0", "8.0.0-beta11", "8.0.0-beta12", "8.0.0-beta14", "8.0.0-rc1", "8.0.0-rc2", "8.0.0-rc3", "8.0.0-rc4", "8.0.0", "8.0.1", "8.0.2", "8.0.3", "8.0.5"}
}

func (drushVersionList *DrushVersionList) PrintLocal() {
	// Print a list of all local versions, see ListLocal().
	drushVersionList.ListLocal()
	for _, value := range drushVersionList.list {
		fmt.Println(value)
	}
}

func (drushVersionList *DrushVersionList) ListRemote() {
	// Fetches a list of all available versions from composer.
	// Versions must start with integers 6,7,8 or 9 to be returned.
	drushVersionsObj := NewDrushVersionList()
	drushVersionsCommand, _ := exec.Command("sh", "-c", PATH_COMPOSER+" show drush/drush -a | grep versions | sort | uniq").Output()
	drushVersions := strings.Split(string(drushVersionsCommand), ", ")
	drushVersionsObj.list = drushVersions
	acceptableVersions := []string{}
	for x := range drushVersions {
		if strings.HasPrefix(drushVersions[x], "6") || strings.HasPrefix(drushVersions[x], "7") || strings.HasPrefix(drushVersions[x], "8") || strings.HasPrefix(drushVersions[x], "9") {
			acceptableVersions = append(acceptableVersions, drushVersions[x])
		}
	}
	drushVersionList.list = acceptableVersions
}

func (drushVersionList *DrushVersionList) PrintRemote() {
	// Print all available remote versions on composer.
	// See ListRemote() for more information.
	drushVersionList.ListRemote()
	for _, value := range drushVersionList.list {
		fmt.Println(value)
	}
}

func (drushVersionList *DrushVersionList) ListInstalled() DrushVersionList {
	// Returns a list of all available installed versions and includes
	// an identifier for the currently used version.
	usr, _ := user.Current()
	workingDir := usr.HomeDir + "/.dvm/versions"
	//thisDrush := getActiveVersion()
	thisDrush := "7.2.0"
	files, _ := ioutil.ReadDir(workingDir)
	installedVersions := NewDrushVersionList()
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "drush-") {
			thisVersion := strings.Replace(file.Name(), "drush-", "", -1)
			if thisDrush == thisVersion {
				fmt.Sprintf("%v*\n", thisVersion)
				installedVersions.list = append(installedVersions.list, thisVersion+"*")
			} else {
				fmt.Sprintln(thisVersion)
				installedVersions.list = append(installedVersions.list, thisVersion)
			}
		}
	}
	return installedVersions
}

func (drushVersionList *DrushVersionList) PrintInstalled() {
	// Prints a list of all installed Command versions.
	// See ListInstalled() for more information.
	drushVersionList.ListInstalled()
	for _, value := range drushVersionList.list {
		fmt.Println(value)
	}
}
