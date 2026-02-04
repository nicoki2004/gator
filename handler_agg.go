package main

import (
	"context"
	"fmt"

	"github.com/nicoki2004/gator/internal/rss"
	"github.com/nicoki2004/gator/internal/state"
)

const (
	rssUrl = "https://www.wagslane.dev/index.xml"
)

func handlerAgg(s *state.State, cmd command) error {
	feed, err := rss.FetchFeed(context.Background(), rssUrl)
	if err != nil {
		return fmt.Errorf("Error requesting feed: %v", err)
	}

	rss.PrintFeed(feed)
	return nil
}
