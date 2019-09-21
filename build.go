package main

import (
	"os"
	"os/exec"
)

func build(inputs Inputs) error {
	args := []string{
		"build",
		"--file", inputs.Dockerfile,
	}

	if inputs.Cache != "" {
		args = append(args, "--from-cache", inputs.Cache)
	}

	for _, tag := range inputs.Tags {
		args = append(args, "--tag", tag)
	}

	args = append(args, inputs.Path)

	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
