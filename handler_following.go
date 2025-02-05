package main

import (
	"context"
	"fmt"

	"github.com/jovanadjuric/rss-aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, f := range feeds {
		fmt.Println("* " + f.FeedName)
	}

	return nil
}
