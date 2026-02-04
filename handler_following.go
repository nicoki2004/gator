package main

import (
	"context"
	"fmt"

	"github.com/nicoki2004/gator/internal/state"
)

func handlerFollowing(s *state.State, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	user, err := s.Db.GetUserByName(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error getting current user: %w", err)
	}

	feedFollows, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error getting feeds: %w", err)
	}

	for _, feedF := range feedFollows {
		fmt.Printf("-----------------------------\n")
		fmt.Printf("* %s\n", feedF.FeedName)
	}

	return nil
}
