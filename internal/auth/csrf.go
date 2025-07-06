package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Returns: csrf_token, csrf_hmac_hash, expiry, error
func GenerateCSRFTokenAndHMAC() (string, string, int64, error) {
	secret := os.Getenv("CSRF_TOKEN_SECRET")
	if secret == "" {
		return "", "", 0, fmt.Errorf("CSRF_TOKEN_SECRET is not set")
	}

	expMs := os.Getenv("CSRF_TOKEN_TIME_IN_MS")
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
