package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/nicoki2004/gator/internal/database"
	"github.com/nicoki2004/gator/internal/state"
)

func handlerFollow(s *state.State, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.Db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Couldn't find the feed: %w", err)
	}

	err = feedFollow(s, feed, user)
	if err != nil {
		return fmt.Errorf("Error creating feed follow: \n", err)
	}

	return nil
}

func feedFollow(s *state.State, feed database.Feed, user database.User) error {
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err := s.Db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Error creating feed follow: %w\n", err)
	}

	fmt.Printf("Feed Follow Created for: \n")
	fmt.Printf("User: %s\n", user.Name)
	fmt.Printf("Feed: %s\n", feed.Name)

	return nil
}
