package solution_test

import (
	target "bootcamp/01-go-basics/solution"
	"slices"
	"testing"
)

func TestSumPositive(t *testing.T) {
	input := []int{-2, 0, 3, 4}
	if got := target.SumPositive(input); got != 7 {
		t.Fatalf("got %d, want 7", got)
	}
	if !slices.Equal(input, []int{-2, 0, 3, 4}) {
		t.Fatal("input was mutated")
	}
	if got := target.SumPositive(nil); got != 0 {
		t.Fatalf("nil: got %d", got)
	}
}
