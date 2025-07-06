package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mcadriel/go-service/internal/auth"
	"github.com/mcadriel/go-service/internal/config"
	"github.com/mcadriel/go-service/internal/model"
)

type Client interface {
	FetchStudentByID(context.Context, string) (*model.Student, error)
}

type client struct {
	cfg    *config.Config
	csrf   *auth.CSRFTokenGenerator
	jwt    *auth.JWTTokenGenerator
	client *http.Client
}

func NewClient(cfg *config.Config, httpClient *http.Client) Client {
	return &client{
		cfg:    cfg,
		csrf:   auth.NewCSRFTokenGenerator(cfg),
		jwt:    auth.NewJWTTokenGenerator(cfg),
		client: httpClient,
	}
}

func (c *client) FetchStudentByID(ctx context.Context, id string) (*model.Student, error) {
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

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set Cookies
	req.AddCookie(&http.Cookie{Name: "accessToken", Value: accessToken})
	req.AddCookie(&http.Cookie{Name: "refreshToken", Value: refreshToken})
	req.AddCookie(&http.Cookie{Name: "csrfToken", Value: csrfToken})

	// Set CSRF Header
	req.Header.Set("x-csrf-token", csrfToken)

	// Send request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Handle non-200 status
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to fetch student data: %s - %s", resp.Status, string(body))
	}

	// Parse response
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var student model.Student
	if err := json.Unmarshal(bodyBytes, &student); err != nil {
		return nil, fmt.Errorf("error decoding student json: %w", err)
	}

	return &student, nil
}
