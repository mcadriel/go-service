package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/mcadriel/go-service/internal/config"
)

type CSRFTokenGenerator struct {
	cfg *config.Config
}

func NewCSRFTokenGenerator(cfg *config.Config) *CSRFTokenGenerator {
	return &CSRFTokenGenerator{
		cfg: cfg,
	}
}

// Returns: csrf_token, csrf_hmac_hash, expiry, error
func (c *CSRFTokenGenerator) GenerateCSRFTokenAndHMAC() (string, string, int64, error) {
	secret := c.cfg.CsrfTokenSecret
	if secret == "" {
		return "", "", 0, fmt.Errorf("CSRF_TOKEN_SECRET is not set")
	}

	expMs := c.cfg.CsrfTokenTimeInMs
	expDur, err := strconv.Atoi(expMs)
	if err != nil {
		expDur = 950000
	}

	expiresAt := time.Now().Add(time.Duration(expDur) * time.Millisecond).Unix()
	csrfToken := fmt.Sprintf("%d", expiresAt)

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(csrfToken))
	csrfHmac := hex.EncodeToString(h.Sum(nil))

	return csrfToken, csrfHmac, expiresAt, nil
}
