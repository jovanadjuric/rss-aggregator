package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jovanadjuric/rss-aggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("follow handler expects a single argument, the url")
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	uuid := uuid.New()
	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid, CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: feed.FID})
	if err != nil {
		return err
	}

	fmt.Println("Feed follow created")
	fmt.Println(follow.FeedName)
	fmt.Println(follow.UserName)

	return nil
}
