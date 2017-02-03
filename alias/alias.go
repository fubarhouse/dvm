package alias

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Alias struct {
	name string
	path string
	uri  string
}

func NewAlias(name, path, alias string) *Alias {
	return &Alias{name, path, alias}
}

func (Alias *Alias) SetName(value string) {
	Alias.name = value
}

func (Alias *Alias) GetName() string {
	return Alias.name
}

func (Alias *Alias) SetUri(value string) {
	Alias.uri = value
}

func (Alias *Alias) GetUri() string {
	return Alias.uri
}

func (Alias *Alias) SetPath(value string) {
	Alias.path = value
}

func (Alias *Alias) GetPath() string {
	return Alias.path
}

func (Alias *Alias) Install() {
	log.Println("Adding alias", Alias.uri)
	tpl, err := template.ParseGlob("templates/*")
	file, err := os.Create("templates/" + Alias.GetUri() + ".alias.drushrc.php")
	tpl.Execute(file, struct {
		Name string
		Path string
		Uri  string
	}{Alias.GetName(), Alias.GetPath(), Alias.GetUri()})

	if err != nil {
		log.Println("Error reading files:", err)
	}

	defer file.Close()
}

func (Alias *Alias) Uninstall() {
	log.Println("Removing alias", Alias.uri)
	os.Remove("templates/" + Alias.GetUri() + ".alias.drushrc.php")

}

func (Alias *Alias) Reinstall() {
	Alias.Uninstall()
	Alias.Install()

}

func (Alias *Alias) GetStatus() bool {
	_, err := os.Stat("templates/" + Alias.GetUri() + ".alias.drushrc.php")
	if err != nil {
		return false
	} else {
		return true
	}
}

func (Alias *Alias) PrintStatus() {
	_, err := os.Stat("templates/" + Alias.GetUri() + ".alias.drushrc.php")
	if err != nil {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}
}
