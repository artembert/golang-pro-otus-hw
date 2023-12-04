package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	SetupEnv(env)
	command := exec.Command(cmd[0], cmd[1:]...) // #nosec G204

	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr

	err := command.Start()
	if err != nil {
		fmt.Println("error while starting command: ", err)
		return -1
	}
	err = command.Wait()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return exitError.ExitCode()
		}
		fmt.Println("error while starting command: ", err)
		return -1
	}
	return
}
