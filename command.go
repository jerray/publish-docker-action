package main

import (
	"io"
	"os/exec"
)

type Commander interface {
	Run(name string, args ...string) error
}

type Command struct {
	out io.Writer
	err io.Writer
}

func NewCommand(w io.Writer, e io.Writer) *Command {
	if e == nil {
		e = w
	}
	return &Command{w, e}
}

func (c *Command) Run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = c.out
	cmd.Stderr = c.err
	return cmd.Run()
}
