package main

import (
	"context"
	"fmt"

	"github.com/nicoki2004/gator/internal/database"
	"github.com/nicoki2004/gator/internal/state"
)

func handlerUsers(s *state.State, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %v", cmd.Name)
	}

	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get users: %w", err)
	}

	printUsers(users, s.Cfg.CurrentUserName)
	return nil
}

func printUsers(users []database.User, current string) {
	for _, user := range users {
		cr := ""
		if user.Name == current {
			cr = " (current)"
		}
		fmt.Printf("* %s%s\n", user.Name, cr)
	}
}
