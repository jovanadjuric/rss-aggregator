package main

import (
	"database/sql"
	"fmt"
	"os"

	config "github.com/jovanadjuric/rss-aggregator/internal/config"
	"github.com/jovanadjuric/rss-aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	s := initState()
	cmds := registerCommands()
	args := os.Args

	if len(args) < 2 {
		fmt.Println("too few arguments")
		os.Exit(1)
	}

	err := cmds.run(s, command{
		name: args[1],
		args: args[2:],
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initState() *state {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading config file", err)
	}

	db, err := sql.Open("postgres", *cfg.Db_Url)
	if err != nil {
		fmt.Println("error connecting to database", err)
	}

	dbQueries := database.New(db)

	s := state{cfg: &cfg, db: dbQueries}

	return &s
}

func registerCommands() *commands {
	cmds := &commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddfeed)
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", handlerFollow)
	cmds.register("following", handlerFollowing)

	return cmds
}
