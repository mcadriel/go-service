package main

import (
	"github.com/mcadriel/go-service/internal/config"
	"github.com/mcadriel/go-service/internal/http/server"
)

func main() {
	// Load environment variables
	cfg := config.LoadConfig()
	s := server.NewHTTPServer(cfg)
	s.Start()
}
