package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/jovanadjuric/rss-aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("unfollow handler expects a single argument, the url")
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.FID})
	if err != nil {
		return err
	}

	fmt.Println("Feed follow deleted")

	return nil
}
