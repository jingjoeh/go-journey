package starter

import "errors"

var ErrInvalidPort = errors.New("invalid port")

func ParsePort(raw string) (int, error) {
	panic("TODO: implement ParsePort")
}
