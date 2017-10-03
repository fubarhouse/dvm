// dvm is a drush version management binary for unix systems.
package main

import (
	"fmt"
	"github.com/fubarhouse/dvm/plugin"
	"github.com/fubarhouse/dvm/version"
	"github.com/fubarhouse/dvm/versionlist"
	"os"

	"github.com/fubarhouse/dvm/conf"
)

// @TODO: Use git tags to discover content dynamically?

// print_usage provides the user with examples of application usage.
func print_usage() {
	fmt.Println("Example usages:")
	fmt.Println("-")
	fmt.Println("dvm install 7.0.0\t\t\tInstall a specified version of Drush")
	fmt.Println("dvm uninstall 7.0.0\t\t\tUninstall a specified version of Drush")
	fmt.Println("dvm reinstall 7.0.0\t\t\tReinstall a specified version of Drush")
	fmt.Println("dvm use 7.0.0\t\t\t\tSpecify the version of Drush to set as in use")
	fmt.Println("-")
	fmt.Println("dvm package install registry_rebuild\tInstall a Drush module")
	fmt.Println("dvm package uninstall registry_rebuild\tUnistall a Drush module")
	fmt.Println("dvm package reinstall registry_rebuild\tReistall a Drush module")
	fmt.Println("-")
	fmt.Println("dvm list installed\t\t\tList installed Drush versions")
	fmt.Println("dvm list available\t\t\tList available Drush versions")
}

func main() {

	_, StatErr := os.Stat("/usr/local/bin/drush")
	if StatErr != nil {
		if len(os.Args) < 2 {
			if os.Args[1] != "install" && os.Args[1] != "reinstall" && os.Args[1] != "use" {
				print_usage()
				fmt.Println("-")
				fmt.Println("No active version in use, please install and activate a drush version.")
				os.Exit(0)
			}
		}
	}
	if os.Args != nil && len(os.Args) >= 2 {

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
		if os.Args[1] == "config" {
			if os.Args[2] == "set" {
				conf.Set("", "")
			} else if os.Args[2] == "get" {

			}
		}
	} else {
		print_usage()
		os.Exit(0)
	}
	os.Exit(0)
}
