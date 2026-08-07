package starter

import "errors"

var ErrInvalidToken = errors.New("invalid token")

func Sign(subject string, secret []byte) (string, error) {
	panic("TODO: implement Sign")
}

func Verify(token string, secret []byte) (string, error) {
	panic("TODO: implement Verify")
}
