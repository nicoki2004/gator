package main

import (
	"context"
	"fmt"

	"github.com/nicoki2004/gator/internal/state"
)

func handlerFeeds(s *state.State, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {

		user, err := s.Db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}

		printFeed(feed, user.Name)
	}
	return nil
}
