package solution_test

import (
	target "bootcamp/02-error-handling/solution"
	"errors"
	"testing"
)

func TestParsePort(t *testing.T) {
	got, err := target.ParsePort("8080")
	if err != nil || got != 8080 {
		t.Fatalf("got (%d, %v)", got, err)
	}
	for _, raw := range []string{"0", "65536", "abc"} {
		_, err := target.ParsePort(raw)
		if !errors.Is(err, target.ErrInvalidPort) {
			t.Errorf("%q: got %v", raw, err)
		}
	}
}
