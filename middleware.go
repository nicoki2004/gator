package main

import (
	"context"

	"github.com/nicoki2004/gator/internal/database"
	"github.com/nicoki2004/gator/internal/state"
)

func middlewareLoggedIn(handler func(s *state.State, cmd command, user database.User) error) func(*state.State, command) error {
	return func(s *state.State, cmd command) error {
		userName := s.Cfg.CurrentUserName

		user, err := s.Db.GetUserByName(context.Background(), userName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
