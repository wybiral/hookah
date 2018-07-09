package protocols

import (
	"os/exec"

	"github.com/google/shlex"
	"github.com/wybiral/hookah/pkg/node"
)

// Exec creates an exec Node
func Exec(cmd string) (*node.Node, error) {
	parts, err := shlex.Split(cmd)
	if err != nil {
		return nil, err
	}
	c := exec.Command(parts[0], parts[1:]...)
	w, err := c.StdinPipe()
	if err != nil {
		return nil, err
	}
	r, err := c.StdoutPipe()
	if err != nil {
		return nil, err
	}
	err = c.Start()
	if err != nil {
		return nil, err
	}
	closer := &cmdcloser{c: c}
	return &node.Node{R: r, W: w, C: closer}, nil
}

type cmdcloser struct {
	c *exec.Cmd
}

func (c *cmdcloser) Close() error {
	return c.c.Process.Kill()
}
