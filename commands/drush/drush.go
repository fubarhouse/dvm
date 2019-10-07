package drush

import (
	"github.com/fubarhouse/dvm/config"
	"os/exec"
	"strings"
)

func run(args []string) ([]byte, error) {

	path, _ := exec.LookPath("drush")
	cmdArgs := []string{
		path,
	}

	for _, arg := range args {
		cmdArgs = append(cmdArgs, arg)
	}
	Command := exec.Cmd{
		Path: config.Path(),
		Args: cmdArgs,
	}
	return Command.Output()
}

func Run(input string) ([]byte, error) {
	args := strings.Split(input, " ")
	return run(args)
}
