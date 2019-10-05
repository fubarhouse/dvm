// version is a package which manages a particular version of Drush
package version

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/fubarhouse/dvm/commands/composer"
	"github.com/fubarhouse/dvm/commands/drush"
	"github.com/fubarhouse/dvm/conf"
	"github.com/fubarhouse/dvm/versionlist"
)

var (
	// usr is the current user.
	usr, _ = user.Current()
	// dvmDirectory is the absolute path to the dvm directory.
	dvmDirectory = usr.HomeDir + sep + ".dvm" + sep
)

// DrushVersion is a struct containing information on a given version of Drush.
type DrushVersion struct {
	// A struct to store a single version and to identify validity via OOP.
	// This is used by many methods to process input data.
	fullVersion  string
	majorVersion int64
	validVersion bool
}

// move will move a file
// temporary proxy to consolidate functionality.
func move(oldPath string, newPath string) error {
	return os.Rename(oldPath, newPath)
}

// remove will remove a path.
// temporary proxy to consolidate functionality.
func remove(path string) error {
	return os.RemoveAll(path)
}

// mkdir will make a directory.
// temporary proxy to consolidate functionality.
func mkdir(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// NewDrushVersion will return a new DrushVersion.
func NewDrushVersion(version string) DrushVersion {
	// An API to create/store a Drush version object.
	retVal := DrushVersion{
		fullVersion: version,
		validVersion: false,
	}
	retVal.SetVersionIdentifier(version)
	retVal.validVersion = retVal.Exists()
	if retVal.validVersion == false {
		// @TODO: Messaging weird when called from non-install command (use)
		log.Fatalf("Input drush v%v was not found in Git tag history or composer.", retVal.fullVersion)
	}
	return retVal
}

// assertFileSystem will ensure the filesystem at ~/.dvm/versions is created for use.
func assertFileSystem() {
	Directory := dvmDirectory + "versions" + sep
	_, StatErr := os.Stat(Directory)
	if StatErr != nil {
		MkdirErr := mkdir(Directory, 0775)
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
	drushVersions.ListAll()
	for _, versionItem := range drushVersions.ListContents() {
		if drushVersion.fullVersion == versionItem {
			return true
		}
	}
	return false
}

// Status will check the installation state of any individual Drush version object.
func (drushVersion *DrushVersion) Status() bool {
	_, err := os.Stat(dvmDirectory + "versions" + sep + "drush-" + drushVersion.fullVersion)
	if err == nil {
		return true
	}
	return false
}

// Install will install a version of drush version with composer in a common location.
func (drushVersion *DrushVersion) Install() {
	assertFileSystem()
	// Installs a version of Drush supported by composer.
	_, err := os.Stat(dvmDirectory + "versions" + sep + "drush-" + drushVersion.fullVersion)
	if err != nil {

		log.Infof("Attempting to install Drush v%v", drushVersion.fullVersion)
		if drushVersion.majorVersion >= 6 {
			directory := fmt.Sprintf("%vversions%vdrush-%v", dvmDirectory, sep, drushVersion.fullVersion)
			mkdir(directory, 0755)
			fmt.Println(directory)
			_, installError := composer.Require(fmt.Sprintf("drush/drush:%v --working-dir=%v", drushVersion.fullVersion, directory))
			if installError != nil {
				log.Errorf("Could not install Drush %v, cleaning installation...", drushVersion.fullVersion)
				log.Errorln(installError)
				remove(fmt.Sprintf("versions%v/drush-%v", dvmDirectory, sep, drushVersion.fullVersion))
			} else {
				log.Infof("Successfully installed Drush v%v", drushVersion.fullVersion)
			}
		} else {
			drushVersion.LegacyInstall()
		}
	} else {
		log.Infof("Drush v%v is already installed.", drushVersion.fullVersion)
	}
}

// Uninstall will remove the file system associated to a given drush version.
func (drushVersion *DrushVersion) Uninstall() {
	// Uninstalls a drush version which was installed using DVM convention.
	_, err := os.Stat(fmt.Sprintf("%vversions%vdrush-%v", dvmDirectory, sep, drushVersion.fullVersion))
	if err == nil {
		log.Infof("Removing installation of Drush v%v", drushVersion.fullVersion)
		rmErr := remove(fmt.Sprintf("%vversions%v/drush-%v", dvmDirectory, sep, drushVersion.fullVersion))
		if rmErr != nil {
			log.Errorln(rmErr)
		} else {
			log.Infof("Successfully uninstalled Drush v%v", drushVersion.fullVersion)
		}
	} else {
		log.Errorf("Drush v%v is not installed.", drushVersion.fullVersion)
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
	if Drushes.IsInstalled(drushVersion.fullVersion) {
		workingDir := dvmDirectory + "versions"
		symlinkSource := ""
		symlinkDest := ""
		if drushVersion.majorVersion > 6 {
			// If the version is supported by composer:
			symlinkSource = conf.Path()
			//if _, err := os.Stat(workingDir + sep + "drush-" + drushVersion.fullVersion + sep + "vendor" + sep + "bin" + sep + "drush"); err == nil {
				symlinkDest = workingDir + sep + "drush-" + drushVersion.fullVersion + sep + "vendor" + sep + "bin" + sep + "drush"
			//}
		} else {
			// If it isn't supported by Composer...
			symlinkSource = conf.Path()
			symlinkDest = workingDir + sep + "drush-" + drushVersion.fullVersion + sep + "drush"
		}

		if drushVersion.validVersion == true {
			// Remove symlink
			rmErr := remove(symlinkSource)
			if rmErr != nil {
				log.Println("Could not remove "+conf.Path()+": ", rmErr)
			} else {
				log.Println("Symlink successfully removed.")
			}
			// Add symlink
			rmErr = os.Symlink(symlinkDest, symlinkSource)
			if rmErr != nil {
				log.Println("Could not sym "+conf.Path()+": ", rmErr)
				log.Println(symlinkDest, "|||", symlinkSource)
			} else {
				log.Println("Symlink successfully created.")
				log.Printf("To use it, run %v or make it available to $PATH", conf.Path())
			}
			// Verify version
			currVer, rmErr := exec.Command(conf.Path(), "--version").Output()
			if rmErr != nil {
				log.Println("Drush returned error: ", rmErr)
				os.Exit(1)
			} else {
				if string(currVer) == drushVersion.fullVersion {
					log.Printf("Drush is now set to v%v", drushVersion.fullVersion)
				}
			}
		} else {
			log.Fatal("Drush version entered is not valid.")
		}
	} else {
		log.Fatalf("Drush version %v is not installed.", drushVersion.fullVersion)
	}
}


// SetVersionIdentifier will parse the input fullVersion to identify the major version.
func (drushVersion *DrushVersion) SetVersionIdentifier(input string) {
	// Assume semantic versioning conventions.
	versionParts := strings.Split(input, ".")
	versionInt, _ := strconv.ParseInt(versionParts[0], 10, 10)
	drushVersion.majorVersion = versionInt
}

// GetActiveVersion will return the currently active drush version.
func GetActiveVersion() string {
	drushOutputVersion, drushOutputError := drush.Run("version --format=string")
	if drushOutputError != nil {
		log.Println(drushOutputError)
		os.Exit(1)
	}
	return string(strings.Replace(string(drushOutputVersion), "\n", "", -1))
}
