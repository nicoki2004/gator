package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nicoki2004/gator/internal/database"
	"github.com/nicoki2004/gator/internal/state"
)

func middlewareLoggedIn(handler func(s *state.State, cmd command, user database.User) error) func(*state.State, command) error {
	return func(s *state.State, cmd command) error {
		if s.Cfg.CurrentUserName == "" {
			return errors.New("you must be logged in first: use 'login <username>'")
		}

		user, err := s.Db.GetUserByName(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("logged-in user %q no longer exists", s.Cfg.CurrentUserName)
			}
			return fmt.Errorf("database error: %w", err)
		}

		return handler(s, cmd, user)
	}
}
