package main

import (
	"github.com/mcadriel/go-service/internal/config"
	"github.com/mcadriel/go-service/internal/http/server"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	s := server.NewHTTPServer()
	s.Start()
}
