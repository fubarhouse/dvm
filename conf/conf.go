package conf

import (
	"os"
	"os/user"
	"github.com/naoina/toml"
)

type config struct {
	Dvm struct {
		Name string
		Path string
	}
}

func getSettings() *config {
	x, _ := user.Current()
	y := x.HomeDir
	cp := y + "/.dvm/config.toml"

	println(cp)
	f, err := os.Open(cp)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var c config
	if err := toml.NewDecoder(f).Decode(&c); err != nil {
		panic(err)
	}
	return &c;
}

func Path() string {
	c := getSettings()
	return c.Dvm.Path
}

func Set(name, value string) {
	getSettings()
}