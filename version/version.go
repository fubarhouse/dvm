package version

import (
	"fmt"
	"github.com/fubarhouse/dvm/versionlist"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

const PATH_DRUSH = "/usr/local/bin/drush"
const PATH_UNZIP = "/usr/bin/unzip"
const PATH_WGET = "/usr/local/bin/wget"
const PATH_COMPOSER = "/usr/local/bin/composer"

type DrushVersion struct {
	// A struct to store a single version and to identify validity via OOP.
	// This is used by many methods to process input data.
	version      string
	validVersion bool
}

func NewDrushVersion(version string) DrushVersion {
	// An API to create/store a Command version object.
	retVal := DrushVersion{version, false}
	retVal.validVersion = retVal.Exists()
	if retVal.validVersion == false {
		log.Fatalf("Input drush v%v was not found in Git tag history or composer.", retVal.version)
	}
	return retVal
}

func (drushVersion *DrushVersion) Exists() bool {
	// Takes in a Command version object and tests if it exists
	// in any available Command version list object.
	drushVersions := versionlist.NewDrushVersionList()
	drushVersions.ListLocal()
	for _, versionItem := range drushVersions.ListContents() {
		if drushVersion.version == versionItem {
			return true
		}
	}
	drushVersions.ListRemote()
	for _, versionItem := range drushVersions.ListContents() {
		if drushVersion.version == versionItem {
			return true
		}
	}
	return false
}

func (drushVersion *DrushVersion) Status() bool {
	// Check the installation state of any individual Command version object.
	usr, _ := user.Current()
	_, err := os.Stat(usr.HomeDir + "/.dvm/versions/drush-" + drushVersion.version)
	if err == nil {
		return true
	}
	return false
}

func (drushVersion *DrushVersion) LegacyInstall() {
	// Basically the main() func for Legacy versions which encapsulates
	// the code/dependencies for installing legacy Command versions.
	drushVersion.LegacyInstallVersion()
	drushVersion.LegacyInstallTable()
}

func (drushVersion *DrushVersion) LegacyInstallTable() {
	// ConsoleTable is essentially always missing from older Command versions.
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
func (drushVersion *DrushVersion) LegacyInstallVersion() {
	// Installs from a zip file which was located via git tags (manual input see ListLocal()).
	// @TODO: Rewrite in the Golang way.
	usr, _ := user.Current()
	fmt.Println("Downloading and extracting legacy Command version ", drushVersion.version)
	zipFileName := drushVersion.version + ".zip"
	remotePath := "https://github.com/drush-ops/drush/archive/" + zipFileName
	zipPath := usr.HomeDir + "/.dvm/versions/"
	zipFile := zipPath + zipFileName
	exec.Command("sh", "-c", "mkdir -p "+zipPath).Run()
	_, wgetErr := exec.Command("sh", "-c", PATH_WGET+" "+remotePath).Output()
	if wgetErr != nil {
		fmt.Println("wget returned error:", wgetErr)
	}
	exec.Command("sh", "-c", "mv "+zipFileName+" "+zipPath).Run()
	exec.Command("sh", "-c", "cd "+zipPath+" && "+PATH_UNZIP+" "+zipFile).Run()
	exec.Command("sh", "-c", "rm -f "+zipFile).Run()
	drushVersion.Status()
}

func (drushVersion *DrushVersion) Install() {
	// Installs a version of Command supported by composer.
	usr, _ := user.Current()
	_, err := os.Stat(usr.HomeDir + "/.dvm/versions/drush-" + drushVersion.version)
	if err != nil {
		majorVersion := fmt.Sprintf("%c", drushVersion.version[0])
		workingDir := usr.HomeDir + "/.dvm/versions"
		fmt.Printf("Attempting to install Command v%v\n", drushVersion.version)

		if majorVersion == "6" || majorVersion == "7" || majorVersion == "8" || majorVersion == "9" {
			_, installError := exec.Command("sh", "-c", "cd "+workingDir+"/ && mkdir -p ./drush-"+drushVersion.version+" && cd ./drush-"+drushVersion.version+" && "+PATH_COMPOSER+" require \"drush/drush:"+drushVersion.version+"\"").Output()
			if installError != nil {
				fmt.Printf("Could not install Command %v, cleaning installation...\n", drushVersion.version)
				fmt.Println(installError)
				exec.Command("sh", "-c", "rm -rf "+workingDir+"/drush-"+drushVersion.version).Output()
			}
		} else {
			drushVersion.LegacyInstall()
		}
	} else {
		fmt.Printf("Command v%v is already installed.\n", drushVersion.version)
	}
}

func (drushVersion *DrushVersion) Uninstall() {
	// Uninstalls a drush version which was installed using DVM convention.
	usr, _ := user.Current()
	_, err := os.Stat(usr.HomeDir + "/.dvm/versions/drush-" + drushVersion.version)
	if err == nil {
		workingDir := usr.HomeDir + "/.dvm/versions"
		fmt.Printf("Removing installation of Command v%v\n", drushVersion.version)
		_, rmErr := exec.Command("sh", "-c", "rm -rf "+workingDir+"/drush-"+drushVersion.version).Output()
		if rmErr != nil {
			fmt.Println(rmErr)
		}
	} else {
		fmt.Printf("Command v%v is not installed.\n", drushVersion.version)
	}
}

func (drushVersion *DrushVersion) Reinstall() {
	// Uninstall and Install an input Command version.
	drushVersion.Uninstall()
	drushVersion.Install()
}

func (drushVersion *DrushVersion) SetDefault() {
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
			fmt.Println("Could not remove "+PATH_DRUSH+": ", rmErr)
		} else {
			fmt.Println("Symlink successfully removed.")
		}
		// Add symlink
		_, rmErr = exec.Command("sh", "-c", "ln -sF "+symlinkDest+" "+symlinkSource).Output()
		if rmErr != nil {
			fmt.Println("Could not sym "+PATH_DRUSH+": ", rmErr)
		} else {
			fmt.Println("Symlink successfully created.")
		}
		// Verify version
		currVer, rmErr := exec.Command("sh", "-c", PATH_DRUSH+" --version").Output()
		if rmErr != nil {
			fmt.Println("Command returned error: ", rmErr)
			os.Exit(1)
		} else {
			if string(currVer) == drushVersion.version {
				fmt.Printf("Command is now set to v%v", drushVersion.version)
			}
		}
	} else {
		log.Fatal("Command version entered is not valid.")
	}
}

func GetActiveVersion() string {
	// Returns the currently active Command version
	drushOutputVersion, drushOutputError := exec.Command(PATH_DRUSH, "version", "--format=string").Output()
	if drushOutputError != nil {
		fmt.Println(drushOutputError)
		os.Exit(1)
	}
	return string(strings.Replace(string(drushOutputVersion), "\n", "", -1))
}
