package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("* Name Url Author")

	for _, f := range feeds {
		fmt.Println("* " + f.FName + " " + f.FUrl + " " + f.UName.String)
	}

	return nil
}
