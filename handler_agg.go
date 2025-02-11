package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
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
		layout := "Mon, 02 Jan 2006 15:04:05 -0700"

		pubDate, err := time.Parse(layout, item.PubDate)
		var publishedAt sql.NullTime

		if err != nil || pubDate.IsZero() {
			fmt.Println("time parsing failed", err)
			publishedAt = sql.NullTime{Valid: false}
		} else {
			publishedAt = sql.NullTime{Time: pubDate, Valid: true}
		}

		uid := uuid.New()
		post, err := s.db.CreatePost(context.Background(), database.CreatePostParams{ID: uid, CreatedAt: time.Now(), UpdatedAt: time.Now(), PublishedAt: publishedAt, Title: item.Title, Description: item.Description, Url: item.Link, FeedID: nextFeed.ID})
		if err != nil {
			if err.Error() == `pq: duplicate key value violates unique constraint "posts_url_key"` {
				continue
			}

			fmt.Println(err)
		}

		fmt.Println("Successfully saved post " + post.ID.String())

	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{ID: nextFeed.ID, UpdatedAt: time.Now()})
	if err != nil {
		return err
	}

	return nil
}
