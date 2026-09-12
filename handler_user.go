package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 || len(cmd.Args) > 1 {
		return errors.New("A username is needed")
	}

	err := s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("The user has been set with username: ", cmd.Args[0])
	return nil
}
