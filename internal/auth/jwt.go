package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken() (string, error) {
	secret := os.Getenv("JWT_ACCESS_TOKEN_SECRET")
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

func GenerateRefreshToken() (string, error) {
	secret := os.Getenv("JWT_REFRESH_TOKEN_SECRET")
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

func GenerateAccessTokenWithCSRF(csrfHmac string) (string, error) {
	secret := os.Getenv("JWT_ACCESS_TOKEN_SECRET")
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
