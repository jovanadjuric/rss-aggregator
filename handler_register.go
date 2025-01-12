package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jovanadjuric/rss-aggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("register handler expects a single argument, the username")
	}

	uuid := uuid.New()
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid, CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.args[0]})
	if err != nil {
		return err
	}

	s.cfg.SetUser(user.Name)

	fmt.Println("user has been registered")
	fmt.Println(user.ID)

	return nil
}
