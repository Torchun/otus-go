package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	var exitError *exec.ExitError

	// lookup over taken envs & unset all found
	for varName, params := range env {
		_, ok := os.LookupEnv(varName)
		if ok {
			// var exists in os env
			err := os.Unsetenv(varName)
			if err != nil {
				fmt.Println(fmt.Errorf("env unset error: %w", err))
				return exitError.ExitCode()
			}
		}

		// already unset, skip to next without setting new value
		if params.NeedRemove {
			continue
		}

		// need to be set with new value
		err := os.Setenv(varName, params.Value)
		if err != nil {
			fmt.Println(fmt.Errorf("env set error: %w", err))
			return exitError.ExitCode()
		}
	}

	command := exec.Command(cmd[0], cmd[1:]...) //nolint: gosec
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Env = os.Environ()

	if err := command.Run(); err != nil {
		fmt.Println(fmt.Errorf("command run error: %w", err))
		return exitError.ExitCode()
	}

	return 0
}
