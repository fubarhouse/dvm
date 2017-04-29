package plugin

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
)

const PATH_DRUSH = "drush"

type drushPackage struct {
	// A struct to store information on a drush package.
	// This is used by associated methods to manage individual packages.
	name   string
	status bool
}

func NewDrushPackage(name string) drushPackage {
	// An API to create/store a Command package object.
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
	// Installs a specified Command package from the current versions core version.
	usr, _ := user.Current()
	workingDir := usr.HomeDir + "/.drush"
	_, err := os.Stat(workingDir + "/" + drushPackage.name + "/")
	if err != nil {
		// err
		_, drushPackageError := exec.Command(PATH_DRUSH, "dl", drushPackage.name).Output()
		if drushPackageError == nil {
			drushPackage.status = drushPackage.Status()
			fmt.Printf("Successfully installed Command package %v\n", drushPackage.name)
		} else {
			fmt.Printf("Could not install Command package %v\n", drushPackageError)
		}
	} else {
		fmt.Printf("Unsuccessfully installed Command package %v: already installed\n", drushPackage.name)
	}
}

func (drushPackage *drushPackage) Uninstall() {
	// Uninstall any drush package based on string input via drushPackage
	usr, _ := user.Current()
	workingDir := usr.HomeDir + "/.drush"
	_, err := os.Stat(usr.HomeDir + "/" + drushPackage.name + "/")
	if err != nil {
		_, drushPackageError := exec.Command("sh", "-c", "rm -rf "+workingDir+"/"+drushPackage.name).Output()
		if drushPackageError == nil {
			drushPackage.status = drushPackage.Status()
			fmt.Printf("Successfully uninstalled Command package %v\n", drushPackage.name)
		} else {
			fmt.Printf("Could not uninstall Command package %v\n", drushPackageError)
		}
	} else {
		fmt.Printf("Unsuccessfully unistalled Command package %v: already uninstalled\n", drushPackage.name)
	}
}

func (drushPackage *drushPackage) Reinstall() {
	// Reinstalls a Command package.
	// Installations are grabbed from the current versions major version.
	drushPackage.Uninstall()
	drushPackage.Install()
}
