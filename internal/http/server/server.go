package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mcadriel/go-service/internal/config"
	"github.com/mcadriel/go-service/internal/handler"
)

type HTTPServer struct {
	r *gin.Engine
}

func NewHTTPServer() *HTTPServer {
	s := &HTTPServer{
		r: gin.Default(),
	}
	s.routes()
	return s
}

func (s *HTTPServer) routes() {
	api := s.r.Group("/api/v1")
	handler.RegisterRoutes(api)
}

func (s *HTTPServer) Start() {
	s.r.Run(":" + config.ListeningPort)
}
