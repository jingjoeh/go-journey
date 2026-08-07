package starter

import (
	"context"
	"errors"
)

var ErrInvalidWorkers = errors.New("workers must be positive")

func SquareAll(ctx context.Context, values []int, workers int) ([]int, error) {
	panic("TODO: implement SquareAll")
}
