package main

import (
	"log"

	"github.com/morfin60/parallel-handlers/internal/config"
	"github.com/morfin60/parallel-handlers/internal/http/server"
	"github.com/morfin60/parallel-handlers/internal/multiplexer"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err.Error())
	}

	server := server.New(cfg)

	multiplexer.RegisterHandlers(server.Handler(), cfg)

	err = server.Start()
	if err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
