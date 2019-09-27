package main

import (
	"reflect"
	"testing"
)

func Test_build(t *testing.T) {
	cases := []struct {
		inputs Inputs
		expect commandArguments
	}{
		// default
		{
			inputs: Inputs{
				Dockerfile: "Dockerfile",
				Path:       ".",
			},
			expect: commandArguments{"docker", []string{"build", "--file", "Dockerfile", "."}},
		},

		// change docker file
		{
			inputs: Inputs{
				Dockerfile: "./builds/Dockerfile",
				Path:       ".",
			},
			expect: commandArguments{"docker", []string{"build", "--file", "./builds/Dockerfile", "."}},
		},

		// change build context
		{
			inputs: Inputs{
				Dockerfile: "Dockerfile",
				Path:       "./builds",
			},
			expect: commandArguments{"docker", []string{"build", "--file", "Dockerfile", "./builds"}},
		},

		// with cache
		{
			inputs: Inputs{
				Dockerfile: "Dockerfile",
				Path:       ".",
				Cache:      "cached_image:latest",
			},
			expect: commandArguments{"docker", []string{"build", "--file", "Dockerfile",
				"--from-cache", "cached_image:latest",
				"."}},
		},

		// with build args
		{
			inputs: Inputs{
				Dockerfile: "Dockerfile",
				Path:       ".",
				BuildArgs:  []string{"USER=root", "GROUP=super"},
			},
			expect: commandArguments{"docker", []string{"build", "--file", "Dockerfile",
				"--build-arg", "USER=root",
				"--build-arg", "GROUP=super",
				".",
			}},
		},

		// with tags
		{
			inputs: Inputs{
				Dockerfile: "Dockerfile",
				Path:       ".",
				Tags:       []string{"image1:latest", "image2:latest"},
			},
			expect: commandArguments{"docker", []string{"build", "--file", "Dockerfile",
				"--tag", "image1:latest",
				"--tag", "image2:latest",
				".",
			}},
		},
	}
	for _, c := range cases {
		cmd := &mockCommander{}
		_ = build(cmd, c.inputs)
		if !reflect.DeepEqual(c.expect, cmd.commands[0]) {
			t.Errorf("expect command %v, actual is %v", c.expect, cmd.commands[0])
		}
	}
}
