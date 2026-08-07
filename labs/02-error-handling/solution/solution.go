package solution

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrInvalidPort = errors.New("invalid port")

func ParsePort(raw string) (int, error) {
	port, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("%w: %q is not a number", ErrInvalidPort, raw)
	}
	if port < 1 || port > 65535 {
		return 0, fmt.Errorf("%w: %d is outside 1..65535", ErrInvalidPort, port)
	}
	return port, nil
}
