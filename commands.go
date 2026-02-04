package main

import (
	"fmt"

	"github.com/nicoki2004/gator/internal/state"
)

type command struct {
	Name string
	Args []string
}

type commandDefinition struct {
	name        string
	description string
	usage       string
	minArgs     int
	handler     func(*state.State, command) error
}

type commands struct {
	registeredCommands map[string]commandDefinition
}

func (c *commands) run(s *state.State, cmd command) error {
	def, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}

	// Validación automática de argumentos
	if len(cmd.Args) < def.minArgs {
		return fmt.Errorf("usage: %s\nerror: %s requires at least %d arguments",
			def.usage, def.name, def.minArgs)
	}

	return def.handler(s, cmd)
	// handler, ok := c.registeredCommands[cmd.Name]
	// if !ok {
	// 	return fmt.Errorf("command not found: %q", cmd.Name)
	// }
	// return handler(s, cmd)
}

func (c *commands) register(def commandDefinition) {
	c.registeredCommands[def.name] = def
}

// func (c *commands) register(name string, f func(*state.State, command) error) {
// 	c.registeredCommands[name] = f
// }
