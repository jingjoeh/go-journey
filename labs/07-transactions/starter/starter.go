package starter

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

	if amount < 1 {
		return ErrInvalidAmount
	}

	return tx.WithinTx(ctx, func(as AccountStore) error {
		if err := as.AddBalance(ctx, from, -amount); err != nil {
			return fmt.Errorf("%w : failed to transfer %d from %s", err, amount, from)
		}

		if err := as.AddBalance(ctx, to, amount); err != nil {
			return fmt.Errorf("%w : failed to transfer %d to %s", err, amount, to)
		}

		return nil
	})

}
