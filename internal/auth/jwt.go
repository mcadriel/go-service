package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mcadriel/go-service/internal/config"
)

type JWTTokenGenerator struct {
	cfg *config.Config
}

func NewJWTTokenGenerator(cfg *config.Config) *JWTTokenGenerator {
	return &JWTTokenGenerator{
		cfg: cfg,
	}
}

func (j *JWTTokenGenerator) GenerateAccessToken() (string, error) {
	secret := j.cfg.JwtAccessTokenSecret
	if secret == "" {
		return "", jwt.ErrTokenMalformed
	}

	claims := jwt.MapClaims{
		"sub": "pdf-service",
		"exp": time.Now().Add(15 * time.Minute).Unix(),
		"iss": "go-service",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (j *JWTTokenGenerator) GenerateRefreshToken() (string, error) {
	secret := j.cfg.JwtRefreshTokenSecret
	if secret == "" {
		return "", jwt.ErrTokenMalformed
	}

	claims := jwt.MapClaims{
		"sub":  "pdf-service",
		"role": "internal-service",
		"exp":  time.Now().Add(8 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (j *JWTTokenGenerator) GenerateAccessTokenWithCSRF(csrfHmac string) (string, error) {
	secret := j.cfg.JwtAccessTokenSecret
	if secret == "" {
		return "", jwt.ErrTokenMalformed
	}

	claims := jwt.MapClaims{
		"sub":       "pdf-service",
		"role":      "internal-service",
		"iss":       "go-service",
		"csrf_hmac": csrfHmac,
		"exp":       time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
