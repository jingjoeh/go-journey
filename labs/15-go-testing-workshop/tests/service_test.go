package tests_test

import "testing"

func TestPlaceOrderInvalidQuantity(t *testing.T) {
	t.Skip("TODO: learner exercise")

	// TODO: Arrange an OrderService with MockOrderRepository.
	// TODO: Call PlaceOrder with a non-positive quantity.
	// TODO: Verify ErrInvalidQuantity and that the repository was not called.
}

func TestPlaceOrderSuccess(t *testing.T) {
	t.Skip("TODO: learner exercise")

	// TODO: Arrange the mock result and expected repository arguments.
	// TODO: Call PlaceOrder with a valid product ID and quantity.
	// TODO: Verify the returned Order and the mock expectations.
}

func TestPlaceOrderRepositoryError(t *testing.T) {
	t.Skip("TODO: learner exercise")

	// TODO: Arrange a sentinel repository error.
	// TODO: Call PlaceOrder with a valid request.
	// TODO: Use ErrorIs to verify that wrapping preserved the original cause.
}
