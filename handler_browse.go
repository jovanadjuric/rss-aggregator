package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jovanadjuric/rss-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2

	if len(cmd.args) > 0 {
		l, err := strconv.ParseInt(cmd.args[0], 10, 32)
		if err != nil {
			return err
		}
		limit = int32(l)
	}

	posts, err := s.db.GetPostsByUserId(context.Background(), database.GetPostsByUserIdParams{UserID: user.ID, Limit: limit})
	if err != nil {
		return err
	}

	for _, p := range posts {
		fmt.Println()
		fmt.Println("* " + p.PTitle + " << " + p.FName)
		fmt.Println("* " + p.PDescription)
	}

	return nil
}
