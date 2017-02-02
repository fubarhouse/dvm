// @TODO: Setup any applicable interfaces and remove pointers to make?
// @TODO: Implement flags properly.
// @TODO: Use git tags to discover content dynamically?

package main

import (
	"flag"
	"fmt"
	"github.com/fubarhouse/dvm/plugin"
	"github.com/fubarhouse/dvm/version"
	"github.com/fubarhouse/dvm/versionlist"
	"os"
)

func main() {

	flagPackage := flag.String("package", "", "Use package flag to tell DVM to target package.")
	flagVersion := flag.String("version", "", "Version to perform action on.")
	flagList := flag.String("list", "", "List to print (installed|available)")
	flagInstall := flag.Bool("install", false, "Version of Command to install")
	flagUninstall := flag.Bool("uninstall", false, "Version of Command to uninstall")
	flagReinstall := flag.Bool("reinstall", false, "Version of Command to reinstall")
	flagSetDefault := flag.Bool("default", false, "Version of Command set as system default")
	flag.Parse()
	if string(*flagVersion) != "" {
		this := version.NewDrushVersion(string(*flagVersion))
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
		this := plugin.NewDrushPackage(*flagPackage)
		if bool(*flagInstall) == true {
			this.Install()
		} else if bool(*flagReinstall) == true {
			this.Reinstall()
		} else if bool(*flagUninstall) == true {
			this.Uninstall()
		} else if string(*flagList) == "available" {
			fmt.Printf("Invalid use --package=%v --list=available. \"available\" is not acceptable by --list flag when targetting Packages.\n", string(*flagPackage))
		} else if string(*flagList) == "installed" {
			this.List()
		} else {
			flag.Usage()
		}
	} else if string(*flagList) != "" {
		Drushes := versionlist.NewDrushVersionList()
		switch string(*flagList) {
		case "available":
			Drushes.PrintRemote()

		case "installed":
			Drushes.PrintInstalled()
		}
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
