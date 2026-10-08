package starter

import (
	"context"
	"fmt"
)

type OrderService struct {
	repository OrderRepository
}

func NewOrderService(repository OrderRepository) *OrderService {
	return &OrderService{repository: repository}
}

func (s *OrderService) PlaceOrder(
	ctx context.Context,
	productID int,
	quantity int,
) (Order, error) {
	if quantity <= 0 {
		return Order{}, ErrInvalidQuantity
	}

	order, err := s.repository.CreateOrder(ctx, productID, quantity)
	if err != nil {
		return Order{}, fmt.Errorf("create order: %w", err)
	}

	return order, nil
}
