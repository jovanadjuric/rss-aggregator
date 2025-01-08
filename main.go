package main

import (
	"fmt"
	"os"

	config "github.com/jovanadjuric/rss-aggregator/internal/config"
	_ "github.com/lib/pq"
)

type state struct {
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

	updatedCfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading config file", err)
	}

	fmt.Println(*updatedCfg.Current_User_Name)
	fmt.Println(*updatedCfg.Db_Url)
}

func initState() *state {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading config file", err)
	}
	s := state{cfg: &cfg}

	return &s
}

func registerCommands() *commands {
	cmds := &commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	return cmds
}
