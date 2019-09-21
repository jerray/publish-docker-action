package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	opts, err := LoadOptions()
	if err != nil {
		fmt.Printf("failed to load options: %s", err)
		os.Exit(1)
	}
	github, inputs := opts.GitHub, opts.Inputs

	typ, name := resolveRef(github)
	if typ == RefTypePull && !inputs.AllowPullRequest {
		fmt.Printf("if you want to build a pull request, please set `with.allow_pull_request` to `true`")
		os.Exit(1)
	}

	resolveAutoTag(typ, name, &inputs)

	for i, t := range inputs.Tags {
		inputs.Tags[i] = strings.Join([]string{inputs.Repository, t}, ":")
	}

	err = build(inputs)
	if err != nil {
		fmt.Printf("failed to build image: %s", err)
		os.Exit(1)
	}

	err = push(inputs)
	if err != nil {
		fmt.Printf("failed to push image: %s", err)
		os.Exit(1)
	}
}
