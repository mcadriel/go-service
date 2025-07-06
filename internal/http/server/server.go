package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mcadriel/go-service/internal/config"
	"github.com/mcadriel/go-service/internal/handler"
	"github.com/mcadriel/go-service/internal/service"
)

type HTTPServer struct {
	r   *gin.Engine
	cfg *config.Config
}

func NewHTTPServer(cfg *config.Config) *HTTPServer {
	s := &HTTPServer{
		r:   gin.Default(),
		cfg: cfg,
	}
	s.routes()
	return s
}

func (s *HTTPServer) routes() {
	api := s.r.Group("/api/v1")

	service := service.NewService(s.cfg)
	handler.RegisterRoutes(api, service)
}

func (s *HTTPServer) Start() {
	s.r.Run(":" + s.cfg.ListeningPort)
}
