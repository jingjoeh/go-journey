package starter

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"
)

var ErrInvalidToken = errors.New("invalid token")

// Sign creates a token in the form base64url(subject).base64url(signature).
// The signature authenticates the original subject bytes with HMAC-SHA256.
func Sign(subject string, secret []byte) (string, error) {
	if subject == "" {
		return "", ErrInvalidToken
	}

	// HMAC combines the subject and secret without exposing the secret in the token.
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(subject))
	signature := mac.Sum(nil)

	// RawURLEncoding is URL-safe and omits padding characters.
	return strings.Join([]string{
		base64.RawURLEncoding.EncodeToString([]byte(subject)),
		base64.RawURLEncoding.EncodeToString(signature),
	}, "."), nil
}

// Verify validates a signed token and returns its subject.
// Every malformed or unauthenticated token is reported as ErrInvalidToken.
func Verify(token string, secret []byte) (string, error) {
	var tokens = strings.Split(token, ".")
	if len(tokens) != 2 {
		return "", ErrInvalidToken
	}

	subject, err := base64.RawURLEncoding.DecodeString(tokens[0])
	if err != nil {
		return "", ErrInvalidToken
	}
	inputToken, err := base64.RawURLEncoding.DecodeString(tokens[1])
	if err != nil {
		return "", ErrInvalidToken
	}

	// Recreate the signature from the decoded subject using the same algorithm as Sign.
	mac := hmac.New(sha256.New, secret)
	mac.Write(subject)
	acceptSignature := mac.Sum(nil)

	// hmac.Equal compares signatures without leaking timing information.
	if hmac.Equal(acceptSignature, inputToken) {
		return string(subject), nil
	}

	return "", ErrInvalidToken
}
