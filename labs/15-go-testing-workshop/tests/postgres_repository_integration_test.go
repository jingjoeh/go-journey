//go:build integration

package tests_test

import (
	"bootcamp/15-go-testing-workshop/starter"
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostgreSQLCreateOrderSuccess(t *testing.T) {

	db := openTestDatabase(t)

	productID := insertProduct(t, db, "MacBook Pro", 10)

	t.Logf("Created product ID: %d", productID)

	repo := starter.NewPostgreSQLOrderRepository(db)

	order, err := repo.CreateOrder(context.Background(), productID, 3)
	require.NoError(t, err)

	assert.NotEmpty(t, order.ID)
	assert.Equal(t, productID, order.ProductID)
	assert.Equal(t, 3, order.Quantity)

	count := orderCount(t, db, productID)
	assert.Equal(t, 1, count)

	stock := productStock(t, db, productID)
	assert.Equal(t, 7, stock)
}

func TestPostgreSQLCreateOrderOutOfStock(t *testing.T) {

	db := openTestDatabase(t)

	productID := insertProduct(t, db, "Macbook Pro", 2)

	repo := starter.NewPostgreSQLOrderRepository(db)

	order, err := repo.CreateOrder(context.Background(), productID, 5)

	require.ErrorIs(t, err, starter.ErrOutOfStock)

	assert.Empty(t, order.ID)

	count := orderCount(t, db, productID)
	assert.Equal(t, 0, count)

	stock := productStock(t, db, productID)
	assert.Equal(t, 2, stock)
}

func TestPostgreSQLCreateOrderRollsBackAfterInsertFailure(t *testing.T) {

	db := openTestDatabase(t)
	productID := insertProduct(t, db, "Macbook Pro", 3)

	repo := starter.NewPostgreSQLOrderRepository(db)
	order, err := repo.CreateOrder(context.Background(), productID, -1)

	require.Error(t, err)
	assert.Empty(t, order.ID)

	count := orderCount(t, db, productID)
	assert.Equal(t, 0, count)

	stock := productStock(t, db, productID)
	assert.Equal(t, 3, stock)
}

func TestPostgreSQLCreateOrderConcurrentLastItem(t *testing.T) {

	db := openTestDatabase(t)
	productID := insertProduct(t, db, "Macbook Pro", 1)
	repo := starter.NewPostgreSQLOrderRepository(db)
	service := starter.NewOrderService(repo)

	var err1 error
	var err2 error

	var order1 starter.Order
	var order2 starter.Order

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		order1, err1 = service.PlaceOrder(context.Background(), productID, 1)
	}()

	go func() {
		defer wg.Done()
		order2, err2 = service.PlaceOrder(context.Background(), productID, 1)
	}()
	wg.Wait()

	orders := []starter.Order{order1, order2}
	errs := []error{err1, err2}

	successCount := 0
	outOfStockCount := 0

	for i, err := range errs {
		if err == nil {
			successCount++
			assert.Positive(t, orders[i].ID)
		} else if errors.Is(err, starter.ErrOutOfStock) {
			outOfStockCount++
			assert.Empty(t, orders[i].ID)
		} else {
			t.Errorf("unexpected error: %v", err)
		}
	}

	assert.Equal(t, 1, successCount)
	assert.Equal(t, 1, outOfStockCount)

	count := orderCount(t, db, productID)
	stock := productStock(t, db, productID)

	assert.Equal(t, 1, count)
	assert.Equal(t, 0, stock)

}
