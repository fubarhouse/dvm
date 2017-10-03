package conf

import (
	"github.com/spf13/viper"
	"log"
	"os/user"
)

// Path will retrieve config.path from the config file.
func Path() string {
	x, _ := user.Current()
	y := x.HomeDir
	cp := y + "/.dvm"

	viper.SetConfigName("config")
	viper.AddConfigPath(cp)
	err := viper.ReadInConfig()

	if err != nil {
		log.Println("No configuration file loaded - using defaults")
		viper.SetDefault("config.path", "/usr/local/bin/drush")
	}

	return viper.GetString("config.path")
}

// Set will set a value in the config to the input.
// Viper does not support writing to config files yet,
// so for now this is largely useless...
func Set(name, value string) error {
	x, _ := user.Current()
	y := x.HomeDir
	cp := y + "/.dvm"

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
