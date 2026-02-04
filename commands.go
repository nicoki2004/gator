package main

import (
	"fmt"

	"github.com/nicoki2004/gator/internal/state"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state.State, command) error
}

func (c *commands) run(s *state.State, cmd command) error {
	handler, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("command not found: %q", cmd.Name)
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state.State, command) error) {
	c.registeredCommands[name] = f
}
