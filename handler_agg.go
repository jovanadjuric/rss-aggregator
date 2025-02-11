package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jovanadjuric/rss-aggregator/internal/database"
	"github.com/jovanadjuric/rss-aggregator/internal/rss_feed"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("agg handler expects a single argument, the duration")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("Collecting feeds every " + cmd.args[0])

	ticker := time.NewTicker(timeBetweenReqs)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("\nFetching feed " + nextFeed.Name)

	feed, err := rss_feed.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	for _, item := range feed.Channel.Items {
		fmt.Println("* " + item.Title)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{ID: nextFeed.ID, UpdatedAt: time.Now()})
	if err != nil {
		return err
	}

	return nil
}
