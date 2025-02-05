package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jovanadjuric/rss-aggregator/internal/database"
)

func handlerAddfeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return errors.New("addfeed handler expects two arguments, name and url")
	}

	uid := uuid.New()
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{ID: uid, CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.args[0], Url: cmd.args[1], UserID: user.ID})
	if err != nil {
		return err
	}

	uid = uuid.New()
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uid, CreatedAt: time.Now(), UpdatedAt: time.Now(), FeedID: feed.ID, UserID: user.ID})
	if err != nil {
		return err
	}

	fmt.Println(feed.ID)

	fmt.Println(feed.Name)
	fmt.Println(feed.Url)

	return nil
}
