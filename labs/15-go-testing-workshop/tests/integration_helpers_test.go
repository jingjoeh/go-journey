//go:build integration

package tests_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const defaultTestDatabaseURL = "postgres://postgres:postgres@localhost:55432/go_testing_workshop?sslmode=disable"

func openTestDatabase(t *testing.T) *sql.DB {
	t.Helper()

	databaseURL := os.Getenv("TEST_DATABASE_URL")
	if databaseURL == "" {
		databaseURL = defaultTestDatabaseURL
	}

	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		t.Fatalf("open PostgreSQL test database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		t.Skipf("PostgreSQL is not available at TEST_DATABASE_URL: %v", err)
	}

	applyMigrations(t, db)
	cleanDatabase(t, db)
	t.Cleanup(func() {
		cleanDatabase(t, db)
		_ = db.Close()
	})

	return db
}

func applyMigrations(t *testing.T, db *sql.DB) {
	t.Helper()

	migrationPath := filepath.Join("..", "migrations", "001_create_orders.sql")
	migration, err := os.ReadFile(migrationPath)
	if err != nil {
		t.Fatalf("read migration %s: %v", migrationPath, err)
	}
	if _, err := db.Exec(string(migration)); err != nil {
		t.Fatalf("apply test migration: %v", err)
	}
}

func cleanDatabase(t *testing.T, db *sql.DB) {
	t.Helper()

	if _, err := db.Exec("TRUNCATE TABLE orders, products RESTART IDENTITY CASCADE"); err != nil {
		t.Fatalf("clean test database: %v", err)
	}
}

func insertProduct(t *testing.T, db *sql.DB, name string, stock int) int {
	t.Helper()

	var productID int
	if err := db.QueryRow(
		"INSERT INTO products (name, stock) VALUES ($1, $2) RETURNING id",
		name,
		stock,
	).Scan(&productID); err != nil {
		t.Fatalf("insert product fixture: %v", err)
	}
	return productID
}

func productStock(t *testing.T, db *sql.DB, productID int) int {
	t.Helper()

	var stock int
	if err := db.QueryRow(
		"SELECT stock FROM products WHERE id = $1",
		productID,
	).Scan(&stock); err != nil {
		t.Fatalf("read product stock: %v", err)
	}
	return stock
}

func orderCount(t *testing.T, db *sql.DB, productID int) int {
	t.Helper()

	var count int
	if err := db.QueryRow(
		"SELECT COUNT(*) FROM orders WHERE product_id = $1",
		productID,
	).Scan(&count); err != nil {
		t.Fatalf("count product orders: %v", err)
	}
	return count
}

func databaseState(t *testing.T, db *sql.DB, productID int) string {
	t.Helper()
	return fmt.Sprintf("stock=%d orders=%d", productStock(t, db, productID), orderCount(t, db, productID))
}
