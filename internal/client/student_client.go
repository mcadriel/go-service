package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mcadriel/go-service/internal/auth"
	"github.com/mcadriel/go-service/internal/config"
	"github.com/mcadriel/go-service/internal/model"
)

type Client interface {
	FetchStudentByID(string) (*model.Student, error)
}

type client struct {
	cfg  *config.Config
	csrf *auth.CSRFTokenGenerator
	jwt  *auth.JWTTokenGenerator
}

func NewClient(cfg *config.Config) Client {
	return &client{
		cfg:  cfg,
		csrf: auth.NewCSRFTokenGenerator(cfg),
		jwt:  auth.NewJWTTokenGenerator(cfg),
	}
}

func (c *client) FetchStudentByID(id string) (*model.Student, error) {
	url := fmt.Sprintf("%s/%s", c.cfg.StudentServiceURL, id)

	// Generate CSRF token and HMAC
	csrfToken, csrfHmac, _, err := c.csrf.GenerateCSRFTokenAndHMAC()
	if err != nil {
		return nil, fmt.Errorf("error generating csrf: %w", err)
	}

	// Generate Access Token with csrf_hmac
	accessToken, err := c.jwt.GenerateAccessTokenWithCSRF(csrfHmac)
	if err != nil {
		return nil, fmt.Errorf("error generating access token: %w", err)
	}

	// Generate Refresh Token
	refreshToken, err := c.jwt.GenerateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("error generating refresh token: %w", err)
	}

	// Build Cookie header
	cookieHeader := fmt.Sprintf("accessToken=%s; refreshToken=%s; csrfToken=%s", accessToken, refreshToken, csrfToken)

	// Build HTTP request with CSRF header
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", cookieHeader)
	req.Header.Set("x-csrf-token", csrfToken)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch student data: %s", resp.Status)
	}

	// Parse response
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var student model.Student
	if err := json.Unmarshal(bodyBytes, &student); err != nil {
		return nil, err
	}

	return &student, nil
}
