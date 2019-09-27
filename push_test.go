package main

import (
	"reflect"
	"testing"
)

func Test_push(t *testing.T) {
	cases := []struct {
		tags    []string
		expects []commandArguments
	}{
		{
			tags: []string{"master", "latest"},
			expects: []commandArguments{
				{"docker", []string{"push", "master"}},
				{"docker", []string{"push", "latest"}},
			},
		},
		{
			tags: []string{"v1.0.0", "v1.0", "v1"},
			expects: []commandArguments{
				{"docker", []string{"push", "v1.0.0"}},
				{"docker", []string{"push", "v1.0"}},
				{"docker", []string{"push", "v1"}},
			},
		},
		{
			tags: []string{"image:2", "image:2.0", "image:2.0.1"},
			expects: []commandArguments{
				{"docker", []string{"push", "image:2"}},
				{"docker", []string{"push", "image:2.0"}},
				{"docker", []string{"push", "image:2.0.1"}},
			},
		},
	}
	for _, c := range cases {
		cmd := &mockCommander{}
		inputs := Inputs{Tags: c.tags}
		_ = push(cmd, inputs)
		for i, expect := range c.expects {
			if !reflect.DeepEqual(expect, cmd.commands[i]) {
				t.Errorf("expect command %v, actual is %v", expect, cmd.commands[i])
			}
		}
	}
}
