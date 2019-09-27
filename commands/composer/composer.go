package composer

import (
	"os/exec"
	"strings"
)

func run(subcommand string, args []string) ([]byte, error) {

	path, _ := exec.LookPath("composer")
	cmdArgs := []string{
		path,
		subcommand,
	}
	for _, arg := range args {
		cmdArgs = append(cmdArgs, arg)
	}
	Command := exec.Cmd{
		Path: path,
		Args: cmdArgs,
	}
	return Command.Output()
}

func Install(input string) ([]byte, error) {
	args := strings.Split(input, " ")
	return run("install", args)
}

func Require(input string) ([]byte, error) {
	args := strings.Split(input, " ")
	return run("require", args)
}

func Show(input string) ([]byte, error) {
	args := strings.Split(input, " ")
	return run("show", args)
}