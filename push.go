package main

import (
	"os"
	"os/exec"
)

func push(inputs Inputs) error {
	for _, tag := range inputs.Tags {
		cmd := exec.Command("docker", "push", tag)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
