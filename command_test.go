package main

import (
	"bytes"
	"testing"
)

type commandArguments struct {
	name string
	args []string
}

type mockCommander struct {
	commands []commandArguments
}

func newMockCommander() *mockCommander {
	return &mockCommander{make([]commandArguments, 0)}
}

func (c *mockCommander) Run(name string, args ...string) error {
	c.commands = append(c.commands, commandArguments{name, args})
	return nil
}

func TestNewCommand(t *testing.T) {
	w := new(bytes.Buffer)
	c := NewCommand(w, nil)
	if c.out != w {
		t.Errorf("unexpected command stdout")
	}
	if c.err != w {
		t.Errorf("unexpected command stderr")
	}
}

func TestCommand_Run(t *testing.T) {
	cases := []struct {
		inputs []string
		expect string
	}{
		{[]string{"hello, world!"}, "hello, world!\n"},
		{[]string{"publish", "docker"}, "publish docker\n"},
		{[]string{"there", "are", "four", "arguments"}, "there are four arguments\n"},
	}
	for _, c := range cases {
		w := new(bytes.Buffer)
		cmd := NewCommand(w, w)
		_ = cmd.Run("echo", c.inputs...)
		if actual := w.String(); c.expect != actual {
			t.Errorf("expect command echo value is `%s`, actual is `%s`", c.expect, actual)
		}
	}
}
