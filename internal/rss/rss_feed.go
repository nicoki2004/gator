package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Return failed with status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	rssFeed := &RSSFeed{}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return nil, err
	}

	return rssFeed, nil
}

func PrintFeed(feed *RSSFeed) {
	// Imprimir información general del canal
	fmt.Printf("--- Feed: %s ---\n", html.UnescapeString(feed.Channel.Title))
	fmt.Printf("Link:        %s\n", feed.Channel.Link)
	fmt.Printf("Description: %s\n", html.UnescapeString(feed.Channel.Description))
	fmt.Println("-----------------------------------------------")

	// Iterar sobre los artículos
	for _, item := range feed.Channel.Item {
		fmt.Printf("Title:       %s\n", html.UnescapeString(item.Title))
		fmt.Printf("Description: %s\n", html.UnescapeString(item.Description))
		fmt.Printf("Link:        %s\n", item.Link)
		// Opcional: imprimir fecha si existe
		if item.PubDate != "" {
			fmt.Printf("Published:   %s\n", item.PubDate)
		}
		fmt.Println("-----------------------------------------------")
	}
}
