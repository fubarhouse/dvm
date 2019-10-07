package config

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
		viper.SetDefault("config.path", binPath)
	}

	return viper.GetString("config.path")
}

