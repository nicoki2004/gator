package main

import (
	"context"
	"fmt"

	"github.com/nicoki2004/gator/internal/state"
)

func handlerReset(s *state.State, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	err := s.Db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete all users: %w", err)
	}

	fmt.Println("Users deleted successfully!")
	return nil
}
