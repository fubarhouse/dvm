package conf

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"os/user"
)

const sep = string(os.PathSeparator)

// Path will retrieve config.path from the config file.
func Path() string {
	x, _ := user.Current()
	y := x.HomeDir
	cp := y + sep + ".dvm"

	binDir := cp + sep + "bin"
	binPath := binDir + sep + "drush"
	if _, fileErr := os.Stat(binDir); fileErr != nil {
		e := os.MkdirAll(binDir, 0755)
		if e != nil {
			log.Printf("Error, could not create path %v\n", binDir)
		}
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(cp)
	err := viper.ReadInConfig()

	if err != nil {
		log.Println("No configuration file loaded - using defaults")
		viper.SetDefault("config.path", binPath)
	}

	return viper.GetString("config.path")
}

// Set will set a value in the config to the input.
// Viper does not support writing to config files yet,
// so for now this is largely useless...
func Set(name, value string) error {
	x, _ := user.Current()
	y := x.HomeDir
	cp := y + sep + ".dvm"

	viper.SetConfigName("config")
	viper.AddConfigPath(cp)
	err := viper.ReadInConfig()

	if err != nil {
		log.Println("No configuration file found - cannot set configuration.")
		return err
	} else {
		viper.Set(name, value)
		log.Printf("Configuration %v was set to %v\n", name, viper.Get(name))
		return nil
	}
}
