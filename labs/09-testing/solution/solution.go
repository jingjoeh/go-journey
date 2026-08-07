package solution

import (
	"errors"
	"strings"
)

var ErrInvalidEmail = errors.New("invalid email")

func NormalizeEmail(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	if strings.Count(raw, "@") != 1 {
		return "", ErrInvalidEmail
	}
	parts := strings.SplitN(raw, "@", 2)
	if parts[0] == "" || parts[1] == "" {
		return "", ErrInvalidEmail
	}
	return parts[0] + "@" + strings.ToLower(parts[1]), nil
}
