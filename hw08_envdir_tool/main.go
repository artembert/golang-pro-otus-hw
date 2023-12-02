package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	envDevPath := args[0]
	commandWithArgs := args[1:]

	envs, err := ReadDir(envDevPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("envs from folder: ", envs)
	// TODO: pass stdin and stdout, stderr to command
	returnCode := RunCmd(commandWithArgs, envs)
	fmt.Printf("Command \"%s\" finished with signal: %d \n", strings.Join(commandWithArgs, " "), returnCode)
	os.Exit(returnCode)
}
