package solution

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"
)

var ErrInvalidToken = errors.New("invalid token")

func Sign(subject string, secret []byte) (string, error) {
	if subject == "" || len(secret) == 0 {
		return "", ErrInvalidToken
	}
	payload := base64.RawURLEncoding.EncodeToString([]byte(subject))
	mac := hmac.New(sha256.New, secret)
	_, _ = mac.Write([]byte(payload))
	return payload + "." + base64.RawURLEncoding.EncodeToString(mac.Sum(nil)), nil
}
func Verify(token string, secret []byte) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 2 || len(secret) == 0 {
		return "", ErrInvalidToken
	}
	got, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", ErrInvalidToken
	}
	mac := hmac.New(sha256.New, secret)
	_, _ = mac.Write([]byte(parts[0]))
	if !hmac.Equal(got, mac.Sum(nil)) {
		return "", ErrInvalidToken
	}
	subject, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil || len(subject) == 0 {
		return "", ErrInvalidToken
	}
	return string(subject), nil
}
