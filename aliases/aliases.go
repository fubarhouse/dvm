package aliases

import (
	"github.com/fubarhouse/golang-drush/command"
	"strings"
)

type AliasList struct {
	// A simple Alias List for attaching methods.
	value []string
}

func NewAliasList() *AliasList {
	// Create a new but empty Alias List
	return &AliasList{}
}

func (list *AliasList) Add(item string) {
	// Add an alias to the alias list.
	list.value = append(list.value, item)
}

func (list *AliasList) Generate(key string) {
	// Add a set of aliases to an Alias List based of a string value.
	sites := command.NewDrushCommand()
	sites.Set("", "sa", false)
	values, _ := sites.Output()
	values = strings.Split(values[0], "\n")
	for _, alias := range values {
		if strings.Contains(alias, key) == true {
			list.Add(alias)
		}
	}
}

func (list *AliasList) Values() []string {
	// Return values from the Alias List object
	return list.value
}