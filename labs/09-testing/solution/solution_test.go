package solution_test

import (
	target "bootcamp/09-testing/solution"
	"errors"
	"testing"
)

func TestNormalizeEmail(t *testing.T) {
	tests := []struct {
		in, want string
		invalid  bool
	}{
		{" Alice@EXAMPLE.COM ", "Alice@example.com", false},
		{"missing", "", true}, {"a@@b", "", true}, {"@host", "", true},
	}
	for _, tc := range tests {
		got, err := target.NormalizeEmail(tc.in)
		if tc.invalid != errors.Is(err, target.ErrInvalidEmail) || got != tc.want {
			t.Errorf("%q: got (%q, %v)", tc.in, got, err)
		}
	}
}
