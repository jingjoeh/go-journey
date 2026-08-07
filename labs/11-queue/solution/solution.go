package solution

import (
	"context"
	"errors"
)

var ErrInvalidAttempts = errors.New("max attempts must be positive")

func Process(ctx context.Context, maxAttempts int, handle func(context.Context) error) error {
	if maxAttempts < 1 {
		return ErrInvalidAttempts
	}
	var err error
	for range maxAttempts {
		if ctxErr := ctx.Err(); ctxErr != nil {
			return ctxErr
		}
		if err = handle(ctx); err == nil {
			return nil
		}
	}
	return err
}
