package main

import (
	"errors"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0])
	command.Env = append(os.Environ(), ToPairsSlice(env)...)

	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr

	err := command.Start()
	if err != nil {
		panic(err)
	}
	err = command.Wait()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return err.(*exec.ExitError).ExitCode()
		}
		panic(err)
	}
	return
}
