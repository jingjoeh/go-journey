package orderprocessor_test

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	orderprocessor "sr-go-bootcamp/exercises/concurrency/orderprocessor"
)

const testTimeout = 2 * time.Second

func TestProcessOrdersProcessesAllOrders(t *testing.T) {
	orders := makeOrders(40)
	processed := make([]atomic.Int32, len(orders))

	err := runWithTimeout(t, func(ctx context.Context) error {
		return orderprocessor.ProcessOrders(ctx, orders, 4, func(_ context.Context, order orderprocessor.Order) error {
			if order.ID < 0 || order.ID >= len(processed) {
				return fmt.Errorf("unexpected order ID %d", order.ID)
			}
			processed[order.ID].Add(1)
			return nil
		})
	})
	if err != nil {
		t.Fatalf("ProcessOrders() error = %v", err)
	}

	for id := range processed {
		if got := processed[id].Load(); got != 1 {
			t.Errorf("order %d processed %d times; want exactly once", id, got)
		}
	}
}

func TestProcessOrdersBoundsConcurrency(t *testing.T) {
	const workerCount = 4

	orders := makeOrders(24)
	release := make(chan struct{})
	reachedLimit := make(chan struct{})
	var reachedOnce sync.Once
	var active atomic.Int32
	var maximum atomic.Int32

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	result := make(chan error, 1)
	go func() {
		result <- orderprocessor.ProcessOrders(ctx, orders, workerCount, func(ctx context.Context, _ orderprocessor.Order) error {
			current := active.Add(1)
			defer active.Add(-1)
			updateMaximum(&maximum, current)
			if current == workerCount {
				reachedOnce.Do(func() { close(reachedLimit) })
			}

			select {
			case <-release:
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		})
	}()

	select {
	case <-reachedLimit:
		close(release)
	case err := <-result:
		t.Fatalf("ProcessOrders() returned before reaching %d concurrent processors: %v", workerCount, err)
	case <-ctx.Done():
		t.Fatal("ProcessOrders() did not reach the configured concurrency before timeout")
	}

	if err := waitForResult(t, ctx, result); err != nil {
		t.Fatalf("ProcessOrders() error = %v", err)
	}
	if got := maximum.Load(); got > workerCount {
		t.Fatalf("maximum concurrency = %d; want at most %d", got, workerCount)
	}
}

func TestProcessOrdersWithOneWorker(t *testing.T) {
	orders := makeOrders(12)
	var active atomic.Int32
	var maximum atomic.Int32
	var processed atomic.Int32

	err := runWithTimeout(t, func(ctx context.Context) error {
		return orderprocessor.ProcessOrders(ctx, orders, 1, func(_ context.Context, _ orderprocessor.Order) error {
			current := active.Add(1)
			defer active.Add(-1)
			updateMaximum(&maximum, current)
			processed.Add(1)
			return nil
		})
	})
	if err != nil {
		t.Fatalf("ProcessOrders() error = %v", err)
	}
	if got := processed.Load(); got != int32(len(orders)) {
		t.Errorf("processed = %d; want %d", got, len(orders))
	}
	if got := maximum.Load(); got != 1 {
		t.Errorf("maximum concurrency = %d; want 1", got)
	}
}

func TestProcessOrdersReturnsProcessorErrorAndCancelsWork(t *testing.T) {
	const workerCount = 4

	orders := makeOrders(100)
	wantErr := errors.New("processor failed")
	workersStarted := make(chan struct{})
	var workersStartedOnce sync.Once
	var started atomic.Int32
	var canceled atomic.Int32

	err := runWithTimeout(t, func(ctx context.Context) error {
		return orderprocessor.ProcessOrders(ctx, orders, workerCount, func(ctx context.Context, order orderprocessor.Order) error {
			if started.Add(1) == workerCount {
				workersStartedOnce.Do(func() { close(workersStarted) })
			}

			if order.ID == 0 {
				select {
				case <-workersStarted:
					return wantErr
				case <-ctx.Done():
					return ctx.Err()
				}
			}

			select {
			case <-ctx.Done():
				canceled.Add(1)
				return ctx.Err()
			case <-workersStarted:
				select {
				case <-ctx.Done():
					canceled.Add(1)
					return ctx.Err()
				}
			}
		})
	})

	if !errors.Is(err, wantErr) {
		t.Fatalf("ProcessOrders() error = %v; want %v", err, wantErr)
	}
	if got := started.Load(); got >= int32(len(orders)) {
		t.Errorf("started %d of %d orders after an early error; want remaining production stopped", got, len(orders))
	}
	if canceled.Load() == 0 {
		t.Error("no in-flight processor observed cancellation")
	}
}

func TestProcessOrdersStopsOnParentCancellation(t *testing.T) {
	const workerCount = 3

	parent, cancelParent := context.WithCancel(context.Background())
	defer cancelParent()

	workersStarted := make(chan struct{})
	var workersStartedOnce sync.Once
	var started atomic.Int32
	var exited atomic.Int32
	result := make(chan error, 1)

	go func() {
		result <- orderprocessor.ProcessOrders(parent, makeOrders(50), workerCount, func(ctx context.Context, _ orderprocessor.Order) error {
			if started.Add(1) == workerCount {
				workersStartedOnce.Do(func() { close(workersStarted) })
			}
			<-ctx.Done()
			exited.Add(1)
			return ctx.Err()
		})
	}()

	select {
	case <-workersStarted:
		cancelParent()
	case err := <-result:
		t.Fatalf("ProcessOrders() returned before parent cancellation: %v", err)
	case <-time.After(testTimeout):
		cancelParent()
		t.Fatal("workers did not start before timeout")
	}

	ctx, cancelWait := context.WithTimeout(context.Background(), testTimeout)
	defer cancelWait()
	if err := waitForResult(t, ctx, result); !errors.Is(err, context.Canceled) {
		t.Fatalf("ProcessOrders() error = %v; want context.Canceled", err)
	}
	if got := exited.Load(); got < workerCount {
		t.Errorf("processors exited after cancellation = %d; want at least %d", got, workerCount)
	}
}

func TestProcessOrdersWithZeroOrders(t *testing.T) {
	var called atomic.Bool

	err := runWithTimeout(t, func(ctx context.Context) error {
		return orderprocessor.ProcessOrders(ctx, nil, 3, func(context.Context, orderprocessor.Order) error {
			called.Store(true)
			return nil
		})
	})
	if err != nil {
		t.Fatalf("ProcessOrders() error = %v", err)
	}
	if called.Load() {
		t.Fatal("processor was called with zero orders")
	}
}

func TestProcessOrdersRejectsInvalidWorkerCount(t *testing.T) {
	for _, workerCount := range []int{0, -1} {
		t.Run(fmt.Sprintf("workerCount=%d", workerCount), func(t *testing.T) {
			var called atomic.Bool
			err := orderprocessor.ProcessOrders(
				context.Background(),
				[]orderprocessor.Order{{ID: 1}},
				workerCount,
				func(context.Context, orderprocessor.Order) error {
					called.Store(true)
					return nil
				},
			)
			if err == nil {
				t.Fatal("ProcessOrders() error = nil; want validation error")
			}
			if called.Load() {
				t.Fatal("processor was called for an invalid worker count")
			}
		})
	}
}

func makeOrders(count int) []orderprocessor.Order {
	orders := make([]orderprocessor.Order, count)
	for id := range orders {
		orders[id] = orderprocessor.Order{ID: id}
	}
	return orders
}

func runWithTimeout(t *testing.T, run func(context.Context) error) error {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	result := make(chan error, 1)
	go func() {
		result <- run(ctx)
	}()
	return waitForResult(t, ctx, result)
}

func waitForResult(t *testing.T, ctx context.Context, result <-chan error) error {
	t.Helper()

	select {
	case err := <-result:
		return err
	case <-ctx.Done():
		t.Fatal("ProcessOrders() did not return before timeout; possible goroutine leak or deadlock")
		return ctx.Err()
	}
}

func updateMaximum(maximum *atomic.Int32, candidate int32) {
	for {
		current := maximum.Load()
		if candidate <= current || maximum.CompareAndSwap(current, candidate) {
			return
		}
	}
}
