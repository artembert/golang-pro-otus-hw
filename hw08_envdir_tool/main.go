package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	envDevPath := args[0]
	commandWithArgs := args[1:]

	envs, err := ReadDir(envDevPath)
	if err != nil {
		panic(err)
	}

	returnCode := RunCmd(commandWithArgs, envs)
	os.Exit(returnCode)
}
