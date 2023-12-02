package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithoutProp := os.Args[1:]
	envDevPath := argsWithoutProp[0]
	commandWithArgs := argsWithoutProp[1:]

	envs, err := ReadDir(envDevPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("envs from folder: ", envs)
	// TODO: pass stdin and stdout, stderr to command
	returnCode := RunCmd(commandWithArgs, envs)
	// TODO: pass exit code from command outside
	fmt.Println("Command finished with signal: ", returnCode)

}
