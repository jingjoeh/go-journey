package starter_test

import (
	target "bootcamp/13-performance/starter"
	"slices"
	"testing"
)

func TestDeduplicate(t *testing.T) {
	in := []string{"b", "a", "b", "c", "a"}
	if got := target.Deduplicate(in); !slices.Equal(got, []string{"b", "a", "c"}) {
		t.Fatalf("got %v", got)
	}
	if !slices.Equal(in, []string{"b", "a", "b", "c", "a"}) {
		t.Fatal("input mutated")
	}
}
func BenchmarkDeduplicate(b *testing.B) {
	in := []string{"a", "b", "a", "c", "d", "b"}
	for b.Loop() {
		target.Deduplicate(in)
	}
}
