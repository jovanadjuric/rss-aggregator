package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jovanadjuric/rss-aggregator/internal/database"
)

func handlerAddfeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return errors.New("addfeed handler expects two arguments, name and url")
	}

	if s.cfg.Current_User_Name == nil {
		return errors.New("no user is currently logged in")
	}

	currentUser, err := s.db.GetUser(context.Background(), *s.cfg.Current_User_Name)
	if err != nil {
		return err
	}

	uid := uuid.New()
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{ID: uid, CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.args[0], Url: cmd.args[1], UserID: currentUser.ID})
	if err != nil {
		return err
	}

	uid = uuid.New()
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uid, CreatedAt: time.Now(), UpdatedAt: time.Now(), FeedID: feed.ID, UserID: currentUser.ID})
	if err != nil {
		return err
	}

	fmt.Println(feed.ID)

	fmt.Println(feed.Name)
	fmt.Println(feed.Url)

	return nil
}
