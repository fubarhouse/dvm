// plugin contains operations tools for managing drush modules
// available through the `drush dl` command with a Go API.
package plugin

import (
	"io/ioutil"
	log "github.com/Sirupsen/logrus"
	"os"
	"os/exec"
	"os/user"
)

// drushPackage is a struct to store information on a drush package.
// This is used by associated methods to manage individual packages.
type drushPackage struct {
	name   string
	status bool
}

// NewDrushPackage will return a new drush package.
func NewDrushPackage(name string) drushPackage {
	drushPackage := new(drushPackage)
	drushPackage.name = name
	drushPackage.status = drushPackage.Status()
	return *drushPackage
}

// Status will return the status of a given drush Package
// Status is determined by the availability of the local
// file system of a drush module.
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

// List will list a set of installed packages available in the local file system.
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
		log.Println(Package)
	}
	return installedPackages
}

// Install will install a drush module to the local file system in ~/.drush/
func (drushPackage *drushPackage) Install() {
	// Installs a specified Command package from the current versions core version.
	usr, _ := user.Current()
	workingDir := usr.HomeDir + "/.drush"
	_, err := os.Stat(workingDir + "/" + drushPackage.name + "/")
	if err != nil {
		// err
		drushPackageOut, drushPackageError := exec.Command("drush", "dl", drushPackage.name).Output()
		if drushPackageError == nil {
			drushPackage.status = drushPackage.Status()
			log.Printf("Successfully installed Command package %v\n", drushPackage.name)
		} else {
			log.Printf("Could not install %v\n%v", drushPackageError, string(drushPackageOut))
		}
	} else {
		log.Printf("Unsuccessfully installed %v: already installed\n", drushPackage.name)
	}
}

// Uninstall will remove a drush module from ~/.drush/
func (drushPackage *drushPackage) Uninstall() {
	// Uninstall any drush package based on string input via drushPackage
	usr, _ := user.Current()
	workingDir := usr.HomeDir + "/.drush"
	_, err := os.Stat(usr.HomeDir + "/" + drushPackage.name + "/")
	if err != nil {
		_, drushPackageError := exec.Command("sh", "-c", "rm -rf "+workingDir+"/"+drushPackage.name).Output()
		if drushPackageError == nil {
			drushPackage.status = drushPackage.Status()
			log.Printf("Successfully uninstalled %v\n", drushPackage.name)
		} else {
			log.Printf("Could not uninstall %v\n", drushPackageError)
		}
	} else {
		log.Printf("Unsuccessfully uninstalled %v: already uninstalled\n", drushPackage.name)
	}
}

// Reinstall will trigger the removal and re-installation of a drush module.
// This is useful when changing between major versions of Drush for compatibility.
func (drushPackage *drushPackage) Reinstall() {
	// Reinstalls a Command package.
	// Installations are grabbed from the current versions major version.
	drushPackage.Uninstall()
	drushPackage.Install()
}
