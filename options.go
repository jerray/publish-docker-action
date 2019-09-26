package main

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type GitHub struct {
	Workflow   string `env:"GITHUB_WORKFLOW"`
	Action     string `env:"GITHUB_ACTION"`
	Actor      string `env:"GITHUB_ACTOR"`
	Repository string `env:"GITHUB_REPOSITORY"`
	Commit     string `env:"GITHUB_SHA"`
	EventName  string `env:"GITHUB_EVENT_NAME"`
	EventPath  string `env:"GITHUB_EVENT_PATH"`
	Ref        string `env:"GITHUB_REF"`
}

type Inputs struct {
	Username         string   `env:"INPUT_USERNAME"`
	Password         string   `env:"INPUT_PASSWORD"`
	Registry         string   `env:"INPUT_REGISTRY"`
	Repository       string   `env:"INPUT_REPOSITORY"`
	Cache            string   `env:"INPUT_CACHE"`
	Dockerfile       string   `env:"INPUT_FILE" envDefault:"Dockerfile"`
	Path             string   `env:"INPUT_PATH" envDefault:"."`
	Tags             []string `env:"INPUT_TAGS" envDefault:"latest" envSeparator:","`
	BuildArgs        []string `env:"INPUT_BUILD_ARGS" envSeparator:","`
	AllowPullRequest bool     `env:"INPUT_ALLOW_PULL_REQUEST"`
	AutoTag          bool     `env:"INPUT_AUTO_TAG"`
}

type Options struct {
	GitHub GitHub
	Inputs Inputs
}

func LoadOptions() (*Options, error) {
	github := GitHub{}
	if err := env.Parse(&github); err != nil {
		return nil, fmt.Errorf("failed to parse github envrionments: %s", err)
	}

	inputs := Inputs{}
	if err := env.Parse(&inputs); err != nil {
		return nil, fmt.Errorf("failed to parse inputs: %s", err)
	}

	if inputs.Repository == "" {
		inputs.Repository = github.Repository
	}

	return &Options{
		GitHub: github,
		Inputs: inputs,
	}, nil
}
