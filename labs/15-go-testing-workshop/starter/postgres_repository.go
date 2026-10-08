package starter

import (
	"context"
	"database/sql"
	"fmt"
)

type PostgreSQLOrderRepository struct {
	db *sql.DB
}

func NewPostgreSQLOrderRepository(db *sql.DB) *PostgreSQLOrderRepository {
	return &PostgreSQLOrderRepository{db: db}
}

func (r *PostgreSQLOrderRepository) CreateOrder(
	ctx context.Context,
	productID int,
	quantity int,
) (Order, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return Order{}, fmt.Errorf("begin create order transaction: %w", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	result, err := tx.ExecContext(ctx, `
		UPDATE products
		SET stock = stock - $1
		WHERE id = $2 AND stock >= $1
	`, quantity, productID)
	if err != nil {
		return Order{}, fmt.Errorf("decrease product stock: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Order{}, fmt.Errorf("read updated product count: %w", err)
	}
	if rowsAffected == 0 {
		return Order{}, ErrOutOfStock
	}

	order := Order{
		ProductID: productID,
		Quantity:  quantity,
	}
	if err := tx.QueryRowContext(ctx, `
		INSERT INTO orders (product_id, quantity)
		VALUES ($1, $2)
		RETURNING id
	`, productID, quantity).Scan(&order.ID); err != nil {
		return Order{}, fmt.Errorf("insert order: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return Order{}, fmt.Errorf("commit create order transaction: %w", err)
	}

	return order, nil
}
