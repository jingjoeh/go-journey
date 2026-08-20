package starter

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrInvalidPort = errors.New("invalid port")

func ParsePort(raw string) (int, error) {

	port, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("cannot parse %q to port %w", raw, ErrInvalidPort)
	}

	if port < 1 || port > 65535 {
		return 0, fmt.Errorf("%d : %w", port, ErrInvalidPort)
	}

	return port, nil
}
