package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/nicoki2004/gator/internal/database"
	"github.com/nicoki2004/gator/internal/rss"
	"github.com/nicoki2004/gator/internal/state"
)

const (
	rssUrl = "https://www.wagslane.dev/index.xml"
)

func handlerAgg(s *state.State, cmd command) error {
	// feed, err := rss.FetchFeed(context.Background(), rssUrl)
	// if err != nil {
	// 	return fmt.Errorf("Error requesting feed: %v", err)
	// }
	// rss.PrintFeed(feed)

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <1m, 1s, 5m>", cmd.Name)
	}

	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Invalid Duration: %w", err)
	}

	// Crear el Ticker
	ticker := time.NewTicker(time_between_reqs)

	fmt.Printf("Collecting feeds every %s...\n", time_between_reqs)

	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("Error crapping feeds: %w", err)
		}
	}
}

func scrapeFeeds(s *state.State) error {
	nextfeed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to fetch next unprocessed feed from database: %w", err)
	}

	_, err = s.Db.MarkFeedFetched(context.Background(), nextfeed.ID)
	if err != nil {
		return fmt.Errorf("Error marking fetched feed: %w", err)
	}

	rssFeeds, err := rss.FetchFeed(context.Background(), nextfeed.Url)
	if err != nil {
		return fmt.Errorf("Error fetching feeds: %w", err)
	}

	// rss.PrintFeed(feed)
	err = savePost(s, rssFeeds, nextfeed.ID)
	if err != nil {
		return fmt.Errorf("Error saving posts: %w", err)
	}
	return nil
}

// Save al posts feeds to a DB
func savePost(s *state.State, feeds *rss.RSSFeed, feed_id uuid.UUID) error {
	for _, post := range feeds.Channel.Item {
		title := sql.NullString{
			String: post.Title,
			Valid:  post.Title != "", // Es válido solo si no está vacío
		}

		description := sql.NullString{
			String: post.Description,
			Valid:  post.Description != "",
		}

		pubDate, err := time.Parse(time.RFC1123Z, post.PubDate)
		if err != nil {
			pubDate = time.Now()
		}
		param := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			PublishedAt: pubDate,
			Title:       title,
			Url:         post.Link,
			Description: description,
			FeedID:      feed_id,
		}

		_, err = s.Db.CreatePost(context.Background(), param)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value") {
				continue
			}
			return fmt.Errorf("Error saving in Db: %w", err)
		}

	}

	return nil
}
