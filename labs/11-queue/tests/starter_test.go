package starter_test

import (
	target "bootcamp/11-queue/starter"
	"context"
	"errors"
	"testing"
)

func TestProcess(t *testing.T) {
	calls := 0
	err := target.Process(context.Background(), 3, func(context.Context) error {
		calls++
		if calls < 3 {
			return errors.New("temporary")
		}
		return nil
	})
	if err != nil || calls != 3 {
		t.Fatalf("got (%d,%v)", calls, err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := target.Process(ctx, 3, func(context.Context) error { t.Fatal("called"); return nil }); !errors.Is(err, context.Canceled) {
		t.Fatalf("got %v", err)
	}
}
