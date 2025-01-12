package main

import (
	"context"
	"fmt"

	"github.com/jovanadjuric/rss-aggregator/internal/rss_feed"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := rss_feed.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
