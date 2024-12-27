package main

import (
	"fmt"
	"os"

	"github.com/jovanadjuric/rss-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading config file", err)
	}

	username := os.Getenv("USER")
	if username == "" {
		username = os.Getenv("USERNAME")
	}

	cfg.SetUser(username)

	updatedCfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading config file", err)
	}

	fmt.Println(*updatedCfg.Current_User_Name)
	fmt.Println(*updatedCfg.Db_Url)
}
