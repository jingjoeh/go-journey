package starter_test

import (
	target "bootcamp/08-auth/starter"
	"errors"
	"testing"
)

func TestTokenRoundTrip(t *testing.T) {
	token, err := target.Sign("user-42", []byte("test-secret"))
	if err != nil {
		t.Fatal(err)
	}
	got, err := target.Verify(token, []byte("test-secret"))
	if err != nil || got != "user-42" {
		t.Fatalf("got (%q, %v)", got, err)
	}
	if _, err := target.Verify(token+"x", []byte("test-secret")); !errors.Is(err, target.ErrInvalidToken) {
		t.Fatalf("got %v", err)
	}
}
