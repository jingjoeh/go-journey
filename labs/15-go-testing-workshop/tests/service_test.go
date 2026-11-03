package tests_test

import (
	"bootcamp/15-go-testing-workshop/mocks"
	"bootcamp/15-go-testing-workshop/starter"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPlaceOrderInvalidQuantity(t *testing.T) {
	mockRepo := new(mocks.MockOrderRepository)
	service := starter.NewOrderService(mockRepo)

	_, err := service.PlaceOrder(t.Context(), 1, -1)
	assert.ErrorIs(t, err, starter.ErrInvalidQuantity)
	mockRepo.AssertNumberOfCalls(t, "CreateOrder", 0)
}

func TestPlaceOrderSuccess(t *testing.T) {

	mockRepo := new(mocks.MockOrderRepository)
	service := starter.NewOrderService(mockRepo)

	mockRepo.On("CreateOrder", mock.Anything, 1, 1).Return(starter.Order{
		ID:        1,
		ProductID: 1,
		Quantity:  1,
	}, nil).Once()

	order, err := service.PlaceOrder(t.Context(), 1, 1)

	require.NoError(t, err)
	assert.Equal(t, 1, order.ProductID)
	assert.Equal(t, 1, order.Quantity)
	mockRepo.AssertNumberOfCalls(t, "CreateOrder", 1)
}

func TestPlaceOrderRepositoryError(t *testing.T) {

	repoErr := errors.New("database unavailable")

	mockRepo := new(mocks.MockOrderRepository)
	service := starter.NewOrderService(mockRepo)

	mockRepo.On("CreateOrder", mock.Anything, mock.Anything, mock.Anything).Return(starter.Order{}, repoErr).Once()

	_, err := service.PlaceOrder(t.Context(), 1, 1)

	require.Error(t, err)
	assert.ErrorIs(t, err, repoErr)
	mockRepo.AssertNumberOfCalls(t, "CreateOrder", 1)

}
