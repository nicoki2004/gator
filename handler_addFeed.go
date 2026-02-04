package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/nicoki2004/gator/internal/database"
	"github.com/nicoki2004/gator/internal/state"
)

func handlerAddFeed(s *state.State, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       url,
		Name:      name,
		UserID:    user.ID,
	}

	feed, err := s.Db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Couldn't create feed: %w", err)
	}

	err = feedFollow(s, feed, user)
	if err != nil {
		return fmt.Errorf("couldn't create the feed follow: %w", err)
	}

	printFeed(feed, s.Cfg.CurrentUserName)
	return nil
}

func printFeed(feed database.Feed, current string) {
	fmt.Printf("*------------------------------------------\n")
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:       	 %s\n", current)
}
