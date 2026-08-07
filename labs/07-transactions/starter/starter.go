package starter

import (
	"context"
	"errors"
)

var ErrInvalidAmount = errors.New("amount must be positive")

type AccountStore interface {
	AddBalance(context.Context, string, int64) error
}

type Transactor interface {
	WithinTx(context.Context, func(AccountStore) error) error
}

func Transfer(ctx context.Context, tx Transactor, from, to string, amount int64) error {
	panic("TODO: implement Transfer")
}
