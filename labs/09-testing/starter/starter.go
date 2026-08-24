package starter

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidEmail = errors.New("invalid email")

func NormalizeEmail(raw string) (string, error) {

	raw = strings.TrimSpace(raw)

	if raw == "" {
		return "", ErrInvalidEmail
	}

	parts := strings.Split(raw, "@")
	if len(parts) != 2 {
		return "", ErrInvalidEmail
	}
	if parts[0] == "" || parts[1] == "" {
		return "", ErrInvalidEmail
	}

	return fmt.Sprintf("%s@%s", parts[0], strings.ToLower(parts[1])), nil
}
