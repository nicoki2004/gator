package main

import (
	"fmt"

	"github.com/nicoki2004/gator/internal/state"
)

func handlerLogin(s *state.State, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("the login handler expects a single argument: the username")
	}
	userName := cmd.Args[0]

	err := s.Cfg.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Printf("User: %s - has been set", userName)

	return nil
}
