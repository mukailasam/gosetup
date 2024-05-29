package main

import (
	"os"
	"os/exec"
)

func runner(arg ...string) error {
	cmd := exec.Command("go", arg...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return errNoConnection
	}

	return nil
}
