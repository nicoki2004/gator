package rss

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchFeed(t *testing.T) {
	// 1. Crear un servidor de prueba con un XML de ejemplo
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		xml := `
		<rss version="2.0">
			<channel>
				<title>Test Feed</title>
				<link>https://test.com</link>
				<description>A test feed</description>
				<item>
					<title>Test Item</title>
					<link>https://test.com/item</link>
					<description>Item description</description>
				</item>
			</channel>
		</rss>`
		w.Write([]byte(xml))
	}))
	defer ts.Close()

	feed, err := FetchFeed(context.Background(), ts.URL)
	if err != nil {
		t.Fatalf("FetchFeed failed: %v", err)
	}

	if feed.Channel.Title != "Test Feed" {
		t.Errorf("Expected title 'Test Feed', got '%s'", feed.Channel.Title)
	}

	if len(feed.Channel.Item) != 1 {
		t.Errorf("Expected 1 item, got %d", len(feed.Channel.Item))
	}
}
