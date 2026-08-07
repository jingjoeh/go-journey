package solution

import (
	"context"
	"errors"
	"fmt"
)

var ErrInvalidAmount = errors.New("amount must be positive")

type AccountStore interface {
	AddBalance(context.Context, string, int64) error
}
type Transactor interface {
	WithinTx(context.Context, func(AccountStore) error) error
}

func Transfer(ctx context.Context, tx Transactor, from, to string, amount int64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	return tx.WithinTx(ctx, func(store AccountStore) error {
		if err := store.AddBalance(ctx, from, -amount); err != nil {
			return fmt.Errorf("debit %s: %w", from, err)
		}
		if err := store.AddBalance(ctx, to, amount); err != nil {
			return fmt.Errorf("credit %s: %w", to, err)
		}
		return nil
	})
}
