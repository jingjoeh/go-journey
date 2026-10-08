package mocks

import (
	"context"

	order "bootcamp/15-go-testing-workshop/starter"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) CreateOrder(
	ctx context.Context,
	productID int,
	quantity int,
) (order.Order, error) {
	arguments := m.Called(ctx, productID, quantity)

	createdOrder, _ := arguments.Get(0).(order.Order)
	return createdOrder, arguments.Error(1)
}
