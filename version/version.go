// version is a package which manages a particular version of Drush
package version

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/fubarhouse/dvm/conf"
	"github.com/fubarhouse/dvm/versionlist"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

// DrushVersion is a struct containing information on a given version of Drush.
type DrushVersion struct {
	// A struct to store a single version and to identify validity via OOP.
	// This is used by many methods to process input data.
	version      string
	validVersion bool
}

// NewDrushVersion will return a new DrushVersion.
func NewDrushVersion(version string) DrushVersion {
	// An API to create/store a Drush version object.
	retVal := DrushVersion{version, false}
	retVal.validVersion = retVal.Exists()
	if retVal.validVersion == false {
		log.Fatalf("Input drush v%v was not found in Git tag history or composer.", retVal.version)
	}
	return retVal
}

// assertFileSystem will ensure the filesystem at ~/.dvm/versions is created for use.
func assertFileSystem() {
	usr, _ := user.Current()
	Directory := usr.HomeDir + "/.dvm/versions/"
	_, StatErr := os.Stat(Directory)
	if StatErr != nil {
		MkdirErr := os.MkdirAll(Directory, 0775)
		if MkdirErr != nil {
			log.Fatalf("Unsuccessfully attempted to create the directory %v with mode 0775.", Directory)
		} else {
			log.Infof("Successfully create the directory %v with mode 0775.", Directory)
		}
	}
}

// Exists will return a bool based on the availability status of a drush version.
func (drushVersion *DrushVersion) Exists() bool {
	// Takes in a Drush version object and tests if it exists
	// in any available Drush version list object.
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

// Status will check the installation state of any individual Drush version object.
func (drushVersion *DrushVersion) Status() bool {
	usr, _ := user.Current()
	_, err := os.Stat(usr.HomeDir + "/.dvm/versions/drush-" + drushVersion.version)
	if err == nil {
		return true
	}
	return false
}

// LegacyInstall is basically the main() func for Legacy versions which encapsulates
// the code/dependencies for installing legacy Drush versions.
func (drushVersion *DrushVersion) LegacyInstall() {
	drushVersion.LegacyInstallVersion()
	drushVersion.LegacyInstallTable()
}

// LegacyInstallTable is essentially always missing from older Drush versions.
// This ensures the script is available to the legacy version.
func (drushVersion *DrushVersion) LegacyInstallTable() {
	// @TODO: Restore functionality in the Golang way...
	//usr, _ := user.Current()
	//log.Infoln("Fixing dependency issue with module Console_Table")
	//ctFileName := "Table.inc"
	//ctRemotePath := "https://raw.githubusercontent.com/pear/Console_Table/master/Table.php"
	//ctPath := usr.HomeDir + "/.dvm/versions/drush-" + drushVersion.version + "/includes/"
	//ctFile := ctPath + ctFileName
	//_, wgetErr := exec.Command("sh", "-c", "wget", ctRemotePath).Output()
	//if wgetErr != nil {
	//	log.Infoln("wget returned error:", wgetErr)
	//}
	//exec.Command("mv ./" + ctFileName + " " + ctFile).Run()
}

// LegacyInstallVersion will install from a zip file which was located via git tags (manual input see ListLocal()).
func (drushVersion *DrushVersion) LegacyInstallVersion() {
	// @TODO: Rewrite in the Golang way.
	usr, _ := user.Current()
	log.Infoln("Downloading and extracting legacy Drush version ", drushVersion.version)
	zipFileName := drushVersion.version + ".zip"
	remotePath := "https://github.com/drush-ops/drush/archive/" + zipFileName
	zipPath := usr.HomeDir + "/.dvm/versions/"
	zipFile := zipPath + zipFileName
	exec.Command("sh", "-c", "mkdir -p "+zipPath).Run()
	_, wgetErr := exec.Command("sh", "-c", "wget", remotePath).Output()
	if wgetErr != nil {
		log.Warnln("wget returned error:", wgetErr)
	}
	exec.Command("sh", "-c", "mv "+zipFileName+" "+zipPath).Run()
	exec.Command("sh", "-c", "cd "+zipPath+" && unzip "+zipFile).Run()
	exec.Command("sh", "-c", "rm -f "+zipFile).Run()
	drushVersion.Status()
}

// Install will install a version of drush version with composer in a common location.
func (drushVersion *DrushVersion) Install() {
	assertFileSystem()
	// Installs a version of Drush supported by composer.
	usr, _ := user.Current()
	_, err := os.Stat(usr.HomeDir + "/.dvm/versions/drush-" + drushVersion.version)
	if err != nil {
		majorVersion := fmt.Sprintf("%c", drushVersion.version[0])
		workingDir := usr.HomeDir + "/.dvm/versions"
		log.Infof("Attempting to install Drush v%v", drushVersion.version)

		if majorVersion == "6" || majorVersion == "7" || majorVersion == "8" || majorVersion == "9" {
			_, installError := exec.Command("sh", "-c", "cd "+workingDir+"/ && mkdir -p ./drush-"+drushVersion.version+" && cd ./drush-"+drushVersion.version+" && composer require \"drush/drush:"+drushVersion.version+"\"").Output()
			if installError != nil {
				log.Errorf("Could not install Drush %v, cleaning installation...", drushVersion.version)
				log.Errorln(installError)
				exec.Command("sh", "-c", "rm -rf "+workingDir+"/drush-"+drushVersion.version).Output()
			} else {
				log.Infof("Successfully installed Drush v%v", drushVersion.version)
			}
		} else {
			drushVersion.LegacyInstall()
		}
	} else {
		log.Infof("Drush v%v is already installed.", drushVersion.version)
	}
}

// Uninstall will remove the file system associated to a given drush version.
func (drushVersion *DrushVersion) Uninstall() {
	// Uninstalls a drush version which was installed using DVM convention.
	usr, _ := user.Current()
	_, err := os.Stat(usr.HomeDir + "/.dvm/versions/drush-" + drushVersion.version)
	if err == nil {
		workingDir := usr.HomeDir + "/.dvm/versions"
		log.Infof("Removing installation of Drush v%v", drushVersion.version)
		_, rmErr := exec.Command("sh", "-c", "rm -rf "+workingDir+"/drush-"+drushVersion.version).Output()
		if rmErr != nil {
			log.Errorln(rmErr)
		} else {
			log.Infof("Successfully uninstalled Drush v%v", drushVersion.version)
		}
	} else {
		log.Errorf("Drush v%v is not installed.", drushVersion.version)
	}
}

// Reinstall will remove and reinstall a drush version.
func (drushVersion *DrushVersion) Reinstall() {
	// Uninstall and Install an input Drush version.
	drushVersion.Uninstall()
	drushVersion.Install()
}

// SetDefault will remove and add a symlink to an specified installation of drush.
func (drushVersion *DrushVersion) SetDefault() {
	Drushes := versionlist.NewDrushVersionList()
	if Drushes.IsInstalled(drushVersion.version) {
		usr, _ := user.Current()
		workingDir := usr.HomeDir + "/.dvm/versions"
		majorVersion := fmt.Sprintf("%c", drushVersion.version[0])
		symlinkSource := ""
		symlinkDest := ""
		if majorVersion == "6" || majorVersion == "7" || majorVersion == "8" || majorVersion == "9" {
			// If the version is supported by composer:
			symlinkSource = conf.Path()
			if _, err := os.Stat(workingDir + "/drush-" + drushVersion.version + "/vendor/bin/drush"); err == nil {
				symlinkDest = workingDir + "/drush-" + drushVersion.version + "/vendor/bin/drush"
			}
		} else {
			// If it isn't supported by Composer...
			symlinkSource = conf.Path()
			symlinkDest = workingDir + "/drush-" + drushVersion.version + "/drush"
		}

		if drushVersion.validVersion == true {
			// Remove symlink
			_, rmErr := exec.Command("sh", "-c", "rm -f "+symlinkSource).Output()
			if rmErr != nil {
				log.Println("Could not remove "+conf.Path()+": ", rmErr)
			} else {
				log.Println("Symlink successfully removed.")
			}
			// Add symlink
			_, rmErr = exec.Command("sh", "-c", "ln -sF "+symlinkDest+" "+symlinkSource).Output()
			if rmErr != nil {
				log.Println("Could not sym "+conf.Path()+": ", rmErr)
			} else {
				log.Println("Symlink successfully created.")
			}
			// Verify version
			currVer, rmErr := exec.Command("sh", "-c", conf.Path()+" --version").Output()
			if rmErr != nil {
				log.Println("Drush returned error: ", rmErr)
				os.Exit(1)
			} else {
				if string(currVer) == drushVersion.version {
					log.Printf("Drush is now set to v%v", drushVersion.version)
				}
			}
		} else {
			log.Fatal("Drush version entered is not valid.")
		}
	} else {
		log.Fatalf("Drush version %v is not installed.", drushVersion.version)
	}
}

// GetActiveVersion will return the currently active drush version.
func GetActiveVersion() string {
	drushOutputVersion, drushOutputError := exec.Command("drush", "version", "--format=string").Output()
	if drushOutputError != nil {
		log.Println(drushOutputError)
		os.Exit(1)
	}
	return string(strings.Replace(string(drushOutputVersion), "\n", "", -1))
}
