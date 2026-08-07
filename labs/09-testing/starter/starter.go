package starter

import "errors"

var ErrInvalidEmail = errors.New("invalid email")

func NormalizeEmail(raw string) (string, error) {
	panic("TODO: implement NormalizeEmail")
}
