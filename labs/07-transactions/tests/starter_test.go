package starter_test

import (
	target "bootcamp/07-transactions/starter"
	"context"
	"errors"
	"testing"
)

type fakeTx struct {
	deltas map[string]int64
	calls  int
}

func (f *fakeTx) WithinTx(ctx context.Context, fn func(target.AccountStore) error) error {
	f.calls++
	return fn(f)
}
func (f *fakeTx) AddBalance(_ context.Context, id string, delta int64) error {
	f.deltas[id] += delta
	return nil
}
func TestTransfer(t *testing.T) {
	tx := &fakeTx{deltas: map[string]int64{}}
	if err := target.Transfer(context.Background(), tx, "a", "b", 25); err != nil {
		t.Fatal(err)
	}
	if tx.calls != 1 || tx.deltas["a"] != -25 || tx.deltas["b"] != 25 {
		t.Fatalf("%+v", tx)
	}
	if err := target.Transfer(context.Background(), tx, "a", "b", 0); !errors.Is(err, target.ErrInvalidAmount) {
		t.Fatalf("got %v", err)
	}
}
