package main

import (
	"github.com/akobir-au/stacks_integration_api/config"
	"github.com/akobir-au/stacks_integration_api/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
