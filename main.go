package main

import (
	"fmt"
	"os"
)

func main() {
	opts, err := LoadOptions()
	if err != nil {
		fmt.Printf("failed to load options: %s", err)
		os.Exit(1)
	}
	github, inputs := opts.GitHub, opts.Inputs

	if err := resolveInputs(github, &inputs); err != nil {
		fmt.Printf("failed to resolve inputs: %s", err)
		os.Exit(1)
	}

	cmd := NewCommand(os.Stdout, os.Stderr)

	if err := build(cmd, inputs); err != nil {
		fmt.Printf("failed to build image: %s", err)
		os.Exit(1)
	}

	if err := push(cmd, inputs); err != nil {
		fmt.Printf("failed to push image: %s", err)
		os.Exit(1)
	}
}
