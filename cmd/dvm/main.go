// dvm is a drush version management binary for unix systems.
package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/fubarhouse/dvm/conf"
	"github.com/fubarhouse/dvm/plugin"
	"github.com/fubarhouse/dvm/version"
	"github.com/fubarhouse/dvm/versionlist"
	"github.com/spf13/viper"
)

// print_usage provides the user with examples of application usage.
func print_usage() {
	fmt.Println("Example usages:")
	fmt.Println("-")
	fmt.Println("dvm install 7.0.0\t\t\t\tInstall a specified version of Drush")
	fmt.Println("dvm uninstall 7.0.0\t\t\t\tUninstall a specified version of Drush")
	fmt.Println("dvm reinstall 7.0.0\t\t\t\tReinstall a specified version of Drush")
	fmt.Println("dvm use 7.0.0\t\t\t\t\tSpecify the version of Drush to set as in use")
	fmt.Println("-")
	fmt.Println("dvm package install registry_rebuild\t\tInstall a Drush module")
	fmt.Println("dvm package uninstall registry_rebuild\t\tUnistall a Drush module")
	fmt.Println("dvm package reinstall registry_rebuild\t\tReistall a Drush module")
	fmt.Println("-")
	fmt.Println("dvm list installed\t\t\t\tList installed Drush versions")
	fmt.Println("dvm list available\t\t\t\tList available Drush versions")
	fmt.Println("-")
	fmt.Println("dvm config get <config_name>\t\t\tList installed Drush versions")
	fmt.Println("dvm config set <config_name> <config_value>\tList available Drush versions")
}

func main() {

	x, _ := user.Current()
	y := x.HomeDir
	cp := y + "/.dvm"

	viper.SetConfigName("config")
	viper.AddConfigPath(cp)
	err := viper.ReadInConfig()

	if len(os.Args) == 1 {
		print_usage()
		os.Exit(0)
	}

	if err != nil {
		log.Println("No configuration file loaded - using defaults")
	}

	_, StatErr := os.Stat(viper.GetString("config.path"))
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
			if len(os.Args) > 2 {
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
			} else {
				print_usage()
				fmt.Println("-")
				fmt.Println("No version argument specified.")
				os.Exit(0)
			}
		}

		if os.Args[1] == "package" {
			if len(os.Args) > 2 {
				if os.Args[2] == "install" || os.Args[2] == "uninstall" || os.Args[2] == "reinstall" {
					if len(os.Args) > 3 {
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
				} else {
					print_usage()
					fmt.Println("-")
					fmt.Println("No valid package action specified.")
					os.Exit(0)
				}
			} else {
				print_usage()
				fmt.Println("-")
				fmt.Println("No package action specified.")
				os.Exit(0)
			}
		}

		if len(os.Args) > 2 {
			if os.Args[1] == "list" {
				Drushes := versionlist.NewDrushVersionList()
				if os.Args[2] == "available" {
					Drushes.PrintRemote()
				} else if os.Args[2] == "installed" {
					Drushes.PrintInstalled()
				} else {
					print_usage()
					fmt.Println("-")
					fmt.Println("No valid action specified.")
					os.Exit(0)
				}
			}
		}

		if os.Args[1] == "config" {
			if os.Args[2] == "set" {
				conf.Set(os.Args[3], os.Args[4])
			} else if os.Args[2] == "get" {

			}
		}
	} else {
		print_usage()
		os.Exit(0)
	}
	os.Exit(0)
}
