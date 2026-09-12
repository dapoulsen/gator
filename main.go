package main

import (
	"fmt"
	"os"

	"github.com/dapoulsen/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		os.Exit(0)
	}
	cfg.SetUser("daniel")

	cfg, err = config.Read()
	if err != nil {
		os.Exit(0)
	}
	fmt.Printf("Struct db url: %s\nConfig username: %s\n", cfg.DBURL, cfg.CurrentUserName)
}
