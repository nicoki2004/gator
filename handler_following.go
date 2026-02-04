package main

import (
	"context"
	"fmt"

	"github.com/nicoki2004/gator/internal/database"
	"github.com/nicoki2004/gator/internal/state"
)

func handlerFollowing(s *state.State, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
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
