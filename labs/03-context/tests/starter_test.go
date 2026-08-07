package starter_test

import (
	target "bootcamp/03-context/starter"
	"context"
	"errors"
	"testing"
	"time"
)

func TestWait(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := target.Wait(ctx, time.Hour); !errors.Is(err, context.Canceled) {
		t.Fatalf("got %v", err)
	}
	if err := target.Wait(context.Background(), 0); err != nil {
		t.Fatal(err)
	}
}
