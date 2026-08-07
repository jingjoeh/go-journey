package starter

import (
	"context"
	"errors"
)

var ErrInvalidAttempts = errors.New("max attempts must be positive")

func Process(ctx context.Context, maxAttempts int, handle func(context.Context) error) error {
	panic("TODO: implement Process")
}
