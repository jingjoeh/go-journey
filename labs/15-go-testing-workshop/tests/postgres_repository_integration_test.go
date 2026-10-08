//go:build integration

package tests_test

import "testing"

func TestPostgreSQLCreateOrderSuccess(t *testing.T) {
	t.Skip("TODO: learner exercise")

	// TODO: Open an isolated test database and insert a product fixture.
	// TODO: Create an order through PostgreSQLOrderRepository.
	// TODO: Verify the returned order, persisted row, and remaining stock.
}

func TestPostgreSQLCreateOrderOutOfStock(t *testing.T) {
	t.Skip("TODO: learner exercise")

	// TODO: Insert a product whose stock is lower than the requested quantity.
	// TODO: Verify ErrOutOfStock.
	// TODO: Verify both stock and order count remain unchanged.
}

func TestPostgreSQLCreateOrderRollsBackAfterInsertFailure(t *testing.T) {
	t.Skip("TODO: learner exercise")

	// TODO: Force the INSERT to fail after the stock UPDATE.
	// Hint: the orders table rejects non-positive quantities.
	// TODO: Verify the error and prove that neither partial state change remains.
}

func TestPostgreSQLCreateOrderConcurrentLastItem(t *testing.T) {
	t.Skip("TODO: learner exercise")

	// TODO: Insert one product with stock=1.
	// TODO: Start two CreateOrder calls concurrently for that same product.
	// TODO: Verify exactly one success, one ErrOutOfStock, stock=0, and one order row.
}
