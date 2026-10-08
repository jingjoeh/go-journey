package starter

import (
	"context"
	"errors"
)

var (
	ErrInvalidQuantity = errors.New("quantity must be positive")
	ErrOutOfStock      = errors.New("product is out of stock")
)

type Order struct {
	ID        int
	ProductID int
	Quantity  int
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, productID, quantity int) (Order, error)
}
