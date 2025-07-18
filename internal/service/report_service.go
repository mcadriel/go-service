package service

import (
	"context"
	"net/http"
	"time"

	"github.com/mcadriel/go-service/internal/client"
	"github.com/mcadriel/go-service/internal/config"
	"github.com/mcadriel/go-service/internal/pdf"
)

type Service interface {
	GenerateStudentReport(string) ([]byte, error)
}

type service struct {
	cfg    *config.Config
	client client.Client
}

func NewService(cfg *config.Config) Service {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	c := client.NewClient(cfg, httpClient)
	return &service{
		cfg:    cfg,
		client: c,
	}
}

func (s *service) GenerateStudentReport(studentID string) ([]byte, error) {
	student, err := s.client.FetchStudentByID(context.Background(), studentID)
	if err != nil {
		return nil, err
	}

	return pdf.GenerateStudentReport(student)
}
