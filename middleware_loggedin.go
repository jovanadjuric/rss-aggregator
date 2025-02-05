package main

import (
	"context"
	"errors"

	"github.com/jovanadjuric/rss-aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		if *s.cfg.Current_User_Name == "" {
			return errors.New("not logged in")
		}

		user, err := s.db.GetUser(context.Background(), *s.cfg.Current_User_Name)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
