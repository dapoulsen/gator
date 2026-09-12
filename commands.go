package main

import "errors"

type commands struct {
	handledCommands map[string]func(*state, command) error
}

type command struct {
	Name string
	Args []string
}

func (c *commands) run(s *state, cmd command) error {
	command, exists := c.handledCommands[cmd.Name]
	if !exists {
		return errors.New("The given command doesn't exist")
	}

	return command(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handledCommands[name] = f
}
