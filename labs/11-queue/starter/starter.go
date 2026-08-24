package starter

import (
	"context"
	"errors"
)

var ErrInvalidAttempts = errors.New("max attempts must be positive")

func Process(ctx context.Context, maxAttempts int, handle func(context.Context) error) error {
	if maxAttempts < 1 {
		return ErrInvalidAttempts
	}
	var atErr error
	for range maxAttempts {
		if err := ctx.Err(); err != nil {
			return err
		}
		err := handle(ctx)
		if err == nil {
			return nil
		} else {
			atErr = err
		}
	}

	return atErr
}
