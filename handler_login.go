package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nicoki2004/gator/internal/state"
)

func handlerLogin(s *state.State, cmd command) error {
	if err := requireArgs(cmd, 1); err != nil {
		return err
	}

	// if len(cmd.Args) != 1 {
	// return fmt.Errorf("usage: %s <name>", cmd.Name)
	// }
	name := cmd.Args[0]

	_, err := s.Db.GetUserByName(context.Background(), name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("user %q does not exist", name)
		}
		return fmt.Errorf("database error while looking up user: %w", err)
	}

	err = s.Cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}
