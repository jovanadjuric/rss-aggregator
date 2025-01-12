package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, u := range users {
		name := u.Name

		if s.cfg.Current_User_Name != nil && *s.cfg.Current_User_Name == name {
			name += " (current)"
		}

		fmt.Println("* " + name)
	}

	return nil
}
