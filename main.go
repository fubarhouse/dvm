// @TODO: Setup any applicable interfaces and remove pointers to make?
// @TODO: Implement flags properly.
// @TODO: Use git tags to discover content dynamically?

package main

import (
	//flag "github.com/ogier/pflag"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"io/ioutil"
	"strings"
)

const PATH_COMPOSER = "/usr/local/bin/composer"
const PATH_DRUSH =  "/usr/local/bin/drush"
const PATH_UNZIP =  "/usr/bin/unzip"
const PATH_WGET =  "/usr/local/bin/wget"

type drushPackage struct {
	// A struct to store information on a drush package.
	// This is used by associated methods to manage individual packages.
	name string
	status bool
}

func newDrushPackage(name string) drushPackage {
	// An API to create/store a Drush package object.
	drushPackage := new(drushPackage)
	drushPackage.name = name
	drushPackage.status = drushPackage.Status()
	return *drushPackage
}

func (drushPackage *drushPackage) Status() bool {
	usr, _ := user.Current()
	workingDir := usr.HomeDir + "/.drush"
	files, _ := ioutil.ReadDir(workingDir)
	installedPackages := []string{}
	for _, file := range files {
		if file.IsDir() == true {
			activeItemDirectory := file.Name()
			_, err := os.Stat(workingDir + "/" + activeItemDirectory + "/")
			if err == nil {
				installedPackages = append(installedPackages, file.Name())
			}
		}
	}
	for _, Package := range installedPackages {
		if drushPackage.name == Package {
			return true
		}
	}
	return false
}

func (drushPackage *drushPackage) List() []string {
	usr, _ := user.Current()
	workingDir := usr.HomeDir + "/.drush"
	files, _ := ioutil.ReadDir(workingDir)
	installedPackages := []string{}
	for _, file := range files {
		if file.IsDir() == true {
			activeItemDirectory := file.Name()
			if activeItemDirectory != "cache" {
				installedPackages = append(installedPackages, file.Name())
			}
		}
	}
	for _, Package := range installedPackages {
		fmt.Println(Package)
	}
	return installedPackages
}

func (drushPackage *drushPackage) Install() {
	// Installs a specified Drush package from the current versions core version.
	usr, _ := user.Current()
	workingDir := usr.HomeDir + "/.drush"
	_, err := os.Stat(workingDir + "/" + drushPackage.name + "/")
	if err != nil {
		// err
		_, drushPackageError := exec.Command(PATH_DRUSH, "dl", drushPackage.name).Output()
		if drushPackageError == nil {
			drushPackage.status = drushPackage.Status()
			fmt.Printf("Successfully installed Drush package %v\n", drushPackage.name)
		} else {
			fmt.Printf("Could not install Drush package %v\n", drushPackageError)
		}
	} else {
		fmt.Printf("Unsuccessfully installed Drush package %v: already installed\n", drushPackage.name)
	}
}

func (drushPackage *drushPackage) Uninstall() {
	// Uninstall any drush package based on string input via drushPackage
	usr, _ := user.Current()
	workingDir := usr.HomeDir + "/.drush"
	_, err := os.Stat(usr.HomeDir + "/" + drushPackage.name + "/")
	if err != nil {
		_, drushPackageError := exec.Command("sh", "-c", "rm -rf " + workingDir + "/" + drushPackage.name).Output()
		if drushPackageError == nil {
			drushPackage.status = drushPackage.Status()
			fmt.Printf("Successfully uninstalled Drush package %v\n", drushPackage.name)
		} else {
			fmt.Printf("Could not uninstall Drush package %v\n", drushPackageError)
		}
	} else {
		fmt.Printf("Unsuccessfully unistalled Drush package %v: already uninstalled\n", drushPackage.name)
	}
}

func (drushPackage *drushPackage) Reinstall() {
	// Reinstalls a Drush package.
	// Installations are grabbed from the current versions major version.
	drushPackage.Uninstall()
	drushPackage.Install()
}

type drushVersion struct {
	// A struct to store a single version and to identify validity via OOP.
	// This is used by many methods to process input data.
	version string
	validVersion bool
}

func newDrushVersion(version string) drushVersion {
	// An API to create/store a Drush version object.
	retVal := drushVersion{version, false}
	retVal.validVersion = retVal.Exists()
	if retVal.validVersion == false {
		log.Fatal("Input drush version was not found in Git tag history or composer.")
	}
	return retVal
}

func (drushVersion *drushVersion) Exists() bool {
	// Takes in a Drush version object and tests if it exists
	// in any available Drush version list object.
	drushVersions := newDrushVersionList()
	drushVersions.ListLocal()
	for _, versionItem := range drushVersions.list {
		if drushVersion.version == versionItem {
			return true
		}
	}
	drushVersions.ListRemote()
	for _, versionItem := range drushVersions.list {
		if drushVersion.version == versionItem {
			return true
		}
	}
	return false
}

func (drushVersion *drushVersion) Status() bool {
	// Check the installation state of any individual Drush version object.
	usr, _ := user.Current()
	_, err := os.Stat(usr.HomeDir + "/.dvm/versions/drush-" + drushVersion.version)
	if err == nil {
		return true
	}
	return false
}

func (drushVersion *drushVersion) LegacyInstall() {
	// Basically the main() func for Legacy versions which encapsulates
	// the code/dependencies for installing legacy Drush versions.
	drushVersion.LegacyInstallVersion()
	drushVersion.LegacyInstallTable()
}

func (drushVersion *drushVersion) LegacyInstallTable() {
	// ConsoleTable is essentially always missing from older Drush versions.
	// This ensures the script is available to the legacy version.
	// @TODO: Restore functionality in the Golang way...
	//usr, _ := user.Current()
	//fmt.Println("Fixing dependency issue with module Console_Table")
	//ctFileName := "Table.inc"
	//ctRemotePath := "https://raw.githubusercontent.com/pear/Console_Table/master/Table.php"
	//ctPath := usr.HomeDir + "/.dvm/versions/drush-" + drushVersion.version + "/includes/"
	//ctFile := ctPath + ctFileName
	//_, wgetErr := exec.Command("sh", "-c", PATH_WGET + " " + ctRemotePath).Output()
	//if wgetErr != nil {
	//	fmt.Println("wget returned error:", wgetErr)
	//}
	//exec.Command("mv ./" + ctFileName + " " + ctFile).Run()
}
func (drushVersion *drushVersion) LegacyInstallVersion() {
	// Installs from a zip file which was located via git tags (manual input see ListLocal()).
	// @TODO: Rewrite in the Golang way.
	usr, _ := user.Current()
	fmt.Println("Downloading and extracting legacy Drush version ", drushVersion.version)
	zipFileName := drushVersion.version + ".zip"
	remotePath := "https://github.com/drush-ops/drush/archive/" + zipFileName
	zipPath := usr.HomeDir + "/.dvm/versions/"
	zipFile := zipPath + zipFileName
	exec.Command("sh", "-c", "mkdir -p " + zipPath).Run()
	_, wgetErr := exec.Command("sh", "-c", PATH_WGET + " " + remotePath).Output()
	if wgetErr != nil {
		fmt.Println("wget returned error:", wgetErr)
	}
	exec.Command("sh", "-c", "mv " + zipFileName + " " + zipPath).Run()
	exec.Command("sh", "-c", "cd " + zipPath + " && " + PATH_UNZIP + " " + zipFile).Run()
	exec.Command("sh", "-c", "rm -f " + zipFile).Run()
	drushVersion.Status()
}

func (drushVersion *drushVersion) Install() {
	// Installs a version of Drush supported by composer.
	usr, _ := user.Current()
	_, err := os.Stat(usr.HomeDir + "/.dvm/versions/drush-" + drushVersion.version)
	if err != nil {
		majorVersion := fmt.Sprintf("%c",drushVersion.version[0])
		workingDir := usr.HomeDir + "/.dvm/versions"
		fmt.Printf("Attempting to install Drush v%v\n", drushVersion.version)

		if majorVersion == "6" || majorVersion == "7" || majorVersion == "8" || majorVersion == "9" {
			_, installError := exec.Command("sh", "-c", "cd " + workingDir + "/ && mkdir -p ./drush-" + drushVersion.version + " && cd ./drush-" + drushVersion.version + " && " + PATH_COMPOSER + " require \"drush/drush:" + drushVersion.version + "\"").Output()
			if installError != nil {
				fmt.Printf("Could not install Drush %v, cleaning installation...\n", drushVersion.version)
				fmt.Println(installError)
				exec.Command("sh", "-c", "rm -rf " + workingDir + "/drush-" + drushVersion.version).Output()
			}
		} else {
			drushVersion.LegacyInstall()
		}
	} else {
		fmt.Printf("Drush v%v is already installed.\n", drushVersion.version)
	}
}

func (drushVersion *drushVersion) Uninstall() {
	// Uninstalls a drush version which was installed using DVM convention.
	usr, _ := user.Current()
	_, err := os.Stat(usr.HomeDir + "/.dvm/versions/drush-" + drushVersion.version)
	if err == nil {
		workingDir := usr.HomeDir + "/.dvm/versions"
		fmt.Printf("Removing installation of Drush v%v\n", drushVersion.version)
		_, rmErr := exec.Command("sh", "-c", "rm -rf " + workingDir + "/drush-" + drushVersion.version).Output()
		if rmErr != nil {
			fmt.Println(rmErr)
		}
	} else {
		fmt.Printf("Drush v%v is not installed.\n", drushVersion.version)
	}
}

func (drushVersion *drushVersion) Reinstall() {
	// Uninstall and Install an input Drush version.
	drushVersion.Uninstall()
	drushVersion.Install()
}

func (drushVersion *drushVersion) SetDefault() {
	// Removes whatever is located at PATH_DRUSH
	// Adds a symlink to an installed version.
	usr, _ := user.Current()
	workingDir := usr.HomeDir + "/.dvm/versions"
	majorVersion := fmt.Sprintf("%c", drushVersion.version[0])
	symlinkSource := ""
	symlinkDest := ""
	if majorVersion == "6" || majorVersion == "7" || majorVersion == "8" || majorVersion == "9" {
		// If the version is supported by composer:
		symlinkSource = PATH_DRUSH
		symlinkDest = workingDir + "/drush-" + drushVersion.version + "/vendor/bin/drush"
	} else {
		// If it isn't supported by Composer...
		symlinkSource = PATH_DRUSH
		symlinkDest = workingDir + "/drush-" + drushVersion.version + "/drush"
	}

	if drushVersion.validVersion == true {
		// Remove symlink
		_, rmErr := exec.Command("sh", "-c", "rm -f "+symlinkSource).Output()
		if rmErr != nil {
			fmt.Println("Could not remove " + PATH_DRUSH + ": ", rmErr)
		} else {
			fmt.Println("Symlink successfully removed.")
		}
		// Add symlink
		_, rmErr = exec.Command("sh", "-c", "ln -sF "+symlinkDest+" "+symlinkSource).Output()
		if rmErr != nil {
			fmt.Println("Could not sym " + PATH_DRUSH + ": ", rmErr)
		} else {
			fmt.Println("Symlink successfully created.")
		}
		// Verify version
		currVer, rmErr := exec.Command("sh", "-c", PATH_DRUSH + " --version").Output()
		if rmErr != nil {
			fmt.Println("Drush returned error: ", rmErr)
			os.Exit(1)
		} else {
			if string(currVer) == drushVersion.version {
				fmt.Printf("Drush is now set to v%v", drushVersion.version)
			}
		}
	} else {
		log.Fatal("Drush version entered is not valid.")
	}
}

func getActiveVersion() string {
	// Returns the currently active Drush version
	drushOutputVersion, drushOutputError := exec.Command(PATH_DRUSH, "version", "--format=string").Output()
	if drushOutputError != nil {
		fmt.Println(drushOutputError)
		return ""
	}
	return string(strings.Replace(string(drushOutputVersion), "\n", "", -1))
}

type drushVersionList struct {
	// A struct to store associated versions in a simple []string.
	// This is used by methods to store and use multiple version data.
	list []string
}

func newDrushVersionList() drushVersionList {
	// An API to create/store a Drush version list object.
	retVal := drushVersionList{}
	return retVal
}

func (drushVersionList *drushVersionList) ListLocal() {
	// Return a list of all local versions of Drush.
	// This is a manually updated array (for performance sake)
	// which stores all valid Drush versions for testing.
	drushVersionList.list = []string{"1.0.0+drupal5", "1.0.0+drupal6", "1.0.0-beta1+drupal5", "1.0.0-beta2+drupal5", "1.0.0-beta3+drupal5", "1.0.0-beta4+drupal5", "1.0.0-rc1+drupal5", "1.0.0-rc1+drupal6", "1.0.0-rc11+drupal6", "1.0.0-rc2+drupal5", "1.0.0-rc2+drupal6", "1.0.0-rc2+drupal7", "1.0.0-rc3+drupal5", "1.1.0+drupal5", "1.1.0+drupal6", "1.2.0+drupal5", "1.2.0+drupal6", "1.3.0+drupal5", "1.4.0+drupal5", "2.0.0", "2.0.0-alpha1+drupal5", "2.0.0-alpha1+drupal6", "2.0.0-alpha1+drupal7", "2.0.0-alpha2+drupal5", "2.0.0-alpha2+drupal6", "2.0.0-alpha2+drupal7", "2.0.0-rc1", "2.1.0", "3.0.0", "3.0.0-alpha1", "3.0.0-beta1", "3.0.0-rc1", "3.0.0-rc2", "3.0.0-rc3", "3.0.0-rc4", "3.1.0", "3.2.0", "3.3.0", "4.0.0", "4.0.0-rc1", "4.0.0-rc10", "4.0.0-rc3", "4.0.0-rc4", "4.0.0-rc5", "4.0.0-rc6", "4.0.0-rc7", "4.0.0-rc8", "4.0.0-rc9", "4.1.0", "4.2.0", "4.3.0", "4.4.0", "4.5.0", "4.5.0-rc1", "4.6.0", "5.0.0", "5.0.0-rc1", "5.0.0-rc2", "5.0.0-rc3", "5.0.0-rc4", "5.0.0-rc5", "5.1.0", "5.2.0", "5.3.0", "5.4.0", "5.5.0", "5.6.0", "5.7.0", "5.8.0", "5.9.0", "6.0.0-rc1", "6.0.0-rc2", "6.0.0-rc3", "6.0.0-rc4", "6.1.0-rc1", "6.0.0", "6.1.0", "6.2.0", "6.3.0", "6.4.0", "6.5.0", "6.6.0", "7.0.0-alpha1", "7.0.0-alpha2", "7.0.0-alpha3", "7.0.0-alpha4", "7.0.0-alpha5", "7.0.0-alpha6", "7.0.0-alpha7", "7.0.0-alpha8", "7.0.0-alpha9", "7.0.0-rc1", "7.0.0-rc2", "7.0.0", "7.1.0", "7.2.0", "8.0.0-beta11", "8.0.0-beta12", "8.0.0-beta14", "8.0.0-rc1", "8.0.0-rc2", "8.0.0-rc3", "8.0.0-rc4", "8.0.0", "8.0.1", "8.0.2", "8.0.3", "8.0.5"}
}

func (drushVersionList *drushVersionList) PrintLocal() {
	// Print a list of all local versions, see ListLocal().
	drushVersionList.ListLocal()
	for _, value := range drushVersionList.list {
		fmt.Println(value)
	}
}

func (drushVersionList *drushVersionList) ListRemote() {
	// Fetches a list of all available versions from composer.
	// Versions must start with integers 6,7,8 or 9 to be returned.
	drushVersionsObj := newDrushVersionList()
	drushVersionsCommand, _ := exec.Command("sh", "-c", PATH_COMPOSER + " show drush/drush -a | grep versions | sort | uniq").Output()
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

func (drushVersionList *drushVersionList) PrintRemote() {
	// Print all available remote versions on composer.
	// See ListRemote() for more information.
	drushVersionList.ListRemote()
	for _, value := range drushVersionList.list {
		fmt.Println(value)
	}
}

func (drushVersionList *drushVersionList) ListInstalled() drushVersionList {
	// Returns a list of all available installed versions and includes
	// an identifier for the currently used version.
	usr, _ := user.Current()
	workingDir := usr.HomeDir + "/.dvm/versions"
	thisDrush := getActiveVersion()
	files, _ := ioutil.ReadDir(workingDir)
	installedVersions := newDrushVersionList()
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "drush-") {
			thisVersion := strings.Replace(file.Name(), "drush-", "", -1)
			if thisDrush == thisVersion {
				fmt.Sprintf("%v*\n", thisVersion)
				installedVersions.list = append(installedVersions.list, thisVersion + "*")
			} else {
				fmt.Sprintln(thisVersion)
				installedVersions.list = append(installedVersions.list, thisVersion)
			}
		}
	}
	return installedVersions
}

func (drushVersionList *drushVersionList) PrintInstalled() {
	// Prints a list of all installed Drush versions.
	// See ListInstalled() for more information.
	drushVersionList.ListInstalled()
	for _, value := range drushVersionList.list {
		fmt.Println(value)
	}
}

func main() {

	flagPackage := flag.String("package", "", "Use package flag to tell DVM to target package.")
	flagVersion := flag.String("version", "", "Version to perform action on.")
	flagList := flag.String("list", "", "List to print (installed|available)")
	flagInstall := flag.Bool("install", false, "Version of Drush to install")
	flagUninstall := flag.Bool("uninstall", false, "Version of Drush to uninstall")
	flagReinstall := flag.Bool("reinstall", false, "Version of Drush to reinstall")
	flagSetDefault := flag.Bool("default", false, "Version of Drush set as system default")

	flag.Parse()

	if string(*flagVersion) != "" {
		this := newDrushVersion(string(*flagVersion))
		if bool(*flagInstall) == true {
			this.Install()
		} else if bool(*flagReinstall) == true {
			this.Reinstall()
		} else if bool(*flagUninstall) == true {
			this.Uninstall()
		} else if bool(*flagSetDefault) == true {
			this.SetDefault()
		} else {
			flag.Usage()
		}
	} else if string(*flagPackage) != "" {
		this := newDrushPackage(*flagPackage)
		if bool(*flagInstall) == true {
			this.Install()
		} else if bool(*flagReinstall) == true {
			this.Reinstall()
		} else if bool(*flagUninstall) == true {
			this.Uninstall()
		} else if string(*flagList) == "available" {
			fmt.Printf("Invalid use --package=%v --list=available. \"available\" is not acceptable by --list flag when targetting Packages.\n", string(*flagPackage))
		} else if string(*flagList) != "" {
			this.List()
		} else {
			flag.Usage()
		}
	} else if string(*flagList) != "" {
		switch string(*flagList) {
		case "available":
			Drushes := newDrushVersionList()
			Drushes.PrintRemote()

		case "installed":
			Drushes := newDrushVersionList()
			Drushes.PrintInstalled()
		}
	} else {
		flag.Usage()
	}

}
