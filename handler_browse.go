package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/nicoki2004/gator/internal/database"
	"github.com/nicoki2004/gator/internal/state"
)

func handlerBrowse(s *state.State, cmd command, user database.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s <posts limit>", cmd.Name)
	}

	limit := int32(2)

	// Validate if args[0] is presset an a valid number
	if len(cmd.Args) > 0 {
		if l, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = int32(l)
		} else {
			// Opcional: avisar si el argumento no era un número válido
			fmt.Printf("Invalid limit '%s', using default: 2\n", cmd.Args[0])
		}
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return fmt.Errorf("Error Getting the posts: %w", err)
	}

	for _, item := range posts {
		fmt.Printf("Title:       %s\n", item.Title)
		fmt.Printf("Description: %s\n", item.Description)
		fmt.Printf("Link:        %s\n", item.Url)
		fmt.Printf("Published:   %s\n", item.PublishedAt)
		fmt.Println("-----------------------------------------------")
	}

	return nil
}
