package versionlist

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fubarhouse/dvm/commands/composer"
	"github.com/fubarhouse/dvm/commands/drush"
)

const (
	sep = string(os.PathSeparator)
)

var (
	supportedLegacyVersions = []string{"1.0.0+drupal5", "1.0.0+drupal6", "1.0.0-beta1+drupal5", "1.0.0-beta2+drupal5", "1.0.0-beta3+drupal5", "1.0.0-beta4+drupal5", "1.0.0-rc1+drupal5", "1.0.0-rc1+drupal6", "1.0.0-rc11+drupal6", "1.0.0-rc2+drupal5", "1.0.0-rc2+drupal6", "1.0.0-rc2+drupal7", "1.0.0-rc3+drupal5", "1.1.0+drupal5", "1.1.0+drupal6", "1.2.0+drupal5", "1.2.0+drupal6", "1.3.0+drupal5", "1.4.0+drupal5", "2.0.0", "2.0.0-alpha1+drupal5", "2.0.0-alpha1+drupal6", "2.0.0-alpha1+drupal7", "2.0.0-alpha2+drupal5", "2.0.0-alpha2+drupal6", "2.0.0-alpha2+drupal7", "2.0.0-rc1", "2.1.0", "3.0.0", "3.0.0-alpha1", "3.0.0-beta1", "3.0.0-rc1", "3.0.0-rc2", "3.0.0-rc3", "3.0.0-rc4", "3.1.0", "3.2.0", "3.3.0", "4.0.0", "4.0.0-rc1", "4.0.0-rc10", "4.0.0-rc3", "4.0.0-rc4", "4.0.0-rc5", "4.0.0-rc6", "4.0.0-rc7", "4.0.0-rc8", "4.0.0-rc9", "4.1.0", "4.2.0", "4.3.0", "4.4.0", "4.5.0", "4.5.0-rc1", "4.6.0", "5.0.0", "5.0.0-rc1", "5.0.0-rc2", "5.0.0-rc3", "5.0.0-rc4", "5.0.0-rc5", "5.1.0", "5.2.0", "5.3.0", "5.4.0", "5.5.0", "5.6.0", "5.7.0", "5.8.0", "5.9.0"}
)

// DrushVersionList is a struct to store associated versions in a simple []string.
// This is used by methods to store and use multiple version data.
type DrushVersionList struct {
	list []string
}

// NewDrushVersionList will return a newly instantiated DrushVersionList.
func NewDrushVersionList() DrushVersionList {
	retVal := DrushVersionList{}
	return retVal
}

func GetVersion() (Version []string) {

	Versions := NewDrushVersionList()
	Versions.ListAll()

	for _, v := range Versions.list {
		Version = append(Version, v)
	}

	sort.Strings(Version)

	return
}

func FindVersion(substring string) {
	for _, v := range GetVersion() {
		if ok, _ := regexp.MatchString(substring, v); ok {
			fmt.Println(v)
		}
	}
}

// ListContents returns a list of all local versions of Command.
// This is a manually updated array (for performance sake)
// which stores all valid Command versions for testing.
func (drushVersionList *DrushVersionList) ListContents() []string {
	return drushVersionList.list
}

// ListLocal returns a list of all local versions of Command.
// This is a manually updated array (for performance sake)
// which stores all valid Command versions for testing.
func (drushVersionList *DrushVersionList) ListLegacy() {
	drushVersionList.list = supportedLegacyVersions
	sort.Strings(drushVersionList.list)
}

// PrintLocal prints a list of all local versions, see ListLocal().
func (drushVersionList *DrushVersionList) PrintLocal() {
	drushVersionList.ListInstalled()
	for _, value := range drushVersionList.list {
		fmt.Println(value)
	}
}

// ListRemote will fetch a list of all available versions from composer.
// Versions must start with integers 6,7,8 or 9 to be returned.
func (drushVersionList *DrushVersionList) ListAll() {
	drushVersionsCommand, _ := composer.Show("drush/drush -a")
	time.Sleep(time.Second * 5)
	
	for _, v := range strings.Split(string(drushVersionsCommand), "\n") {
		if strings.HasPrefix(v, "versions") {

			var acceptableVersions = make([]string, 0)
			composerVersions := strings.Split(string(drushVersionsCommand), ", ")

			for x, id := range supportedLegacyVersions {
				num := strings.Split(id, ".")[0]
				if i, e := strconv.ParseInt(num, 10, 10); e == nil {
					if i <= 5 {
						acceptableVersions = append(acceptableVersions, supportedLegacyVersions[x])
					}
				}
			}

			for x, id := range composerVersions {
				num := strings.Split(id, ".")[0]
				if i, e := strconv.ParseInt(num, 10, 10); e == nil {
					if i >= 6 {
						acceptableVersions = append(acceptableVersions, composerVersions[x])
					}
				}
			}
			drushVersionList.list = acceptableVersions
		}
	}
}

// PrintRemote will print all available remote versions via composer.
// See ListRemote() for more information.
func (drushVersionList *DrushVersionList) PrintRemote() {
	drushVersionList.ListAll()
	for _, value := range drushVersionList.list {
		fmt.Println(value)
	}
}

// ListInstalled returns a list of all available installed versions and includes
// an identifier for the currently used version.
func (drushVersionList *DrushVersionList) ListInstalled() DrushVersionList {
	usr, _ := user.Current()
	workingDir := usr.HomeDir + "/.dvm/versions"
	thisDrush := GetActiveVersion()
	//thisDrush := "7.2.0"
	files, _ := ioutil.ReadDir(workingDir)
	installedVersions := NewDrushVersionList()
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "drush-") {
			thisVersion := strings.Replace(file.Name(), "drush-", "", -1)
			if thisDrush == thisVersion {
				fmt.Sprintf("%v*\n", thisVersion)
				installedVersions.list = append(installedVersions.list, fmt.Sprintf("%v (in use)", thisVersion))
			} else {
				fmt.Sprintln(thisVersion)
				installedVersions.list = append(installedVersions.list, thisVersion)
			}
		}
	}
	return installedVersions
}

// PrintInstalled prints a list of all installed drush versions.
// See ListInstalled() for more information.
func (drushVersionList *DrushVersionList) PrintInstalled() {
	InstalledVersions := drushVersionList.ListInstalled()
	for _, value := range InstalledVersions.list {
		fmt.Println(value)
	}
}

// GetInstalled returns a list of all installed drush versions.
func (drushVersionList *DrushVersionList) GetInstalled() []string {
	InstalledVersions := drushVersionList.ListInstalled()
	var versions []string
	for _, value := range InstalledVersions.list {
		versions = append(versions, value)
	}
	return versions
}

// IsInstalled returns a boolean of the status of an input version.
func (drushVersionList *DrushVersionList) IsInstalled(version string) bool {
	usr, _ := user.Current()
	workingDir := usr.HomeDir + sep + ".dvm" + sep + "versions" + sep + "drush-" + version
	if _, err := os.Stat(workingDir); err == nil {
		return true
	}
	return false
}

// GetActiveVersion returns the currently active Command version
func GetActiveVersion() string {
	drushOutputVersion, drushOutputError := drush.Run("version --format=string")
	if drushOutputError != nil {
		fmt.Println(drushOutputError)
		os.Exit(1)
	}
	return string(strings.Replace(string(drushOutputVersion), "\n", "", -1))
}
