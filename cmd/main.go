package main

import (
	"fmt"
	"github.com/GalahadKingsman/messenger_users/internal/app"
	"github.com/GalahadKingsman/messenger_users/internal/config"
	"github.com/caarlos0/env/v11"
	"log"
)

func main() {

	cfg := &config.Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", cfg)

	app.Run(cfg)
}
