package solution_test

import (
	target "bootcamp/04-concurrency/solution"
	"context"
	"errors"
	"slices"
	"testing"
)

func TestSquareAll(t *testing.T) {
	got, err := target.SquareAll(context.Background(), []int{3, -2, 5}, 2)
	if err != nil || !slices.Equal(got, []int{9, 4, 25}) {
		t.Fatalf("got (%v, %v)", got, err)
	}
	if _, err := target.SquareAll(context.Background(), nil, 0); !errors.Is(err, target.ErrInvalidWorkers) {
		t.Fatalf("got %v", err)
	}
}
