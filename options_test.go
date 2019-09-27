package main

import (
	"os"
	"reflect"
	"testing"
)

func TestLoadOptions(t *testing.T) {
	type innerCase struct {
		name   string
		env    map[string]string
		expect *Options
	}
	cases := []innerCase{
		{
			name: "default values",
			env: map[string]string{
				"GITHUB_WORKFLOW":          "main",
				"GITHUB_ACTION":            "action name",
				"GITHUB_ACTOR":             "actor",
				"GITHUB_REPOSITORY":        "jerray/publish-docker-action",
				"GITHUB_SHA":               "3b5ec30992fc35b8dd59716f093bf3166a92925b",
				"GITHUB_EVENT_NAME":        "push",
				"GITHUB_EVENT_PATH":        "/",
				"GITHUB_REF":               "refs/heads/master",
				"INPUT_USERNAME":           "",
				"INPUT_PASSWORD":           "",
				"INPUT_REGISTRY":           "",
				"INPUT_REPOSITORY":         "",
				"INPUT_CACHE":              "",
				"INPUT_ALLOW_PULL_REQUEST": "false",
				"INPUT_AUTO_TAG":           "false",
			},
			expect: &Options{
				GitHub: GitHub{
					Workflow:   "main",
					Action:     "action name",
					Actor:      "actor",
					Repository: "jerray/publish-docker-action",
					Commit:     "3b5ec30992fc35b8dd59716f093bf3166a92925b",
					EventName:  "push",
					EventPath:  "/",
					Ref:        "refs/heads/master",
				},
				Inputs: Inputs{
					Username:         "",
					Password:         "",
					Registry:         "",
					Repository:       "",
					Cache:            "",
					Dockerfile:       "Dockerfile",
					Path:             ".",
					Tags:             []string{"latest"},
					BuildArgs:        nil,
					AllowPullRequest: false,
					AutoTag:          false,
				},
			},
		},
		{
			name: "specific values",
			env: map[string]string{
				"GITHUB_WORKFLOW":          "my workflow",
				"GITHUB_ACTION":            "action name",
				"GITHUB_ACTOR":             "actor",
				"GITHUB_REPOSITORY":        "jerray/publish-docker-action",
				"GITHUB_SHA":               "3b5ec30992fc35b8dd59716f093bf3166a92925b",
				"GITHUB_EVENT_NAME":        "push",
				"GITHUB_EVENT_PATH":        "/",
				"GITHUB_REF":               "refs/heads/master",
				"INPUT_USERNAME":           "jerray",
				"INPUT_PASSWORD":           "password",
				"INPUT_REGISTRY":           "docker.pkg.github.com",
				"INPUT_REPOSITORY":         "",
				"INPUT_CACHE":              "",
				"INPUT_FILE":               "./builds/Dockerfile",
				"INPUT_PATH":               "./builds",
				"INPUT_TAGS":               "tag1,tag2",
				"INPUT_BUILD_ARGS":         "X=1,Y=2,Z=3",
				"INPUT_ALLOW_PULL_REQUEST": "true",
				"INPUT_AUTO_TAG":           "true",
			},
			expect: &Options{
				GitHub: GitHub{
					Workflow:   "my workflow",
					Action:     "action name",
					Actor:      "actor",
					Repository: "jerray/publish-docker-action",
					Commit:     "3b5ec30992fc35b8dd59716f093bf3166a92925b",
					EventName:  "push",
					EventPath:  "/",
					Ref:        "refs/heads/master",
				},
				Inputs: Inputs{
					Username:         "jerray",
					Password:         "password",
					Registry:         "docker.pkg.github.com",
					Repository:       "",
					Cache:            "",
					Dockerfile:       "./builds/Dockerfile",
					Path:             "./builds",
					Tags:             []string{"tag1", "tag2"},
					BuildArgs:        []string{"X=1", "Y=2", "Z=3"},
					AllowPullRequest: true,
					AutoTag:          true,
				},
			},
		},
	}

	for _, c := range cases {
		func(c innerCase) {
			t.Run(c.name, func(t *testing.T) {
				for k, v := range c.env {
					_ = os.Setenv(k, v)
				}
				opts, _ := LoadOptions()
				for k := range c.env {
					_ = os.Unsetenv(k)
				}
				if !reflect.DeepEqual(c.expect, opts) {
					t.Errorf("expect options is %#v, actual is %#v", c.expect, opts)
				}
			})
		}(c)
	}
}
