package main

import (
	"github.com/fubarhouse/dvm/plugin"
	"github.com/fubarhouse/dvm/version"
	"github.com/fubarhouse/dvm/versionlist"
	"os"
)

// @TODO: Use git tags to discover content dynamically?

func main() {

	if os.Args[1] == "install" || os.Args[1] == "uninstall" || os.Args[1] == "reinstall" || os.Args[1] == "use" {
		Action := os.Args[1]
		Version := os.Args[2]
		this := version.NewDrushVersion(Version)
		if Action == "install" {
			this.Install()
		} else if Action == "reinstall" {
			this.Reinstall()
		} else if Action == "uninstall" {
			this.Uninstall()
		} else if Action == "use" {
			this.SetDefault()
		}
	}
	if os.Args[1] == "package" {
		if os.Args[2] == "install" || os.Args[2] == "uninstall" || os.Args[2] == "reinstall" {
			Action := os.Args[2]
			PackageName := os.Args[3]
			this := plugin.NewDrushPackage(PackageName)
			if Action == "install" {
				this.Install()
			} else if Action == "reinstall" {
				this.Reinstall()
			} else if Action == "uninstall" {
				this.Uninstall()
			}
		}
	}
	if os.Args[1] == "list" {
		Drushes := versionlist.NewDrushVersionList()
		if os.Args[2] == "available" {
			Drushes.PrintRemote()
		} else if os.Args[2] == "installed" {
			Drushes.PrintInstalled()
		}
	}
}
