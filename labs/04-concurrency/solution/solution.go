package solution

import (
	"context"
	"errors"
	"sync"
)

var ErrInvalidWorkers = errors.New("workers must be positive")

func SquareAll(ctx context.Context, values []int, workers int) ([]int, error) {
	if workers < 1 {
		return nil, ErrInvalidWorkers
	}
	out := make([]int, len(values))
	jobs := make(chan int)
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				out[i] = values[i] * values[i]
			}
		}()
	}
	for i := range values {
		select {
		case jobs <- i:
		case <-ctx.Done():
			close(jobs)
			wg.Wait()
			return nil, ctx.Err()
		}
	}
	close(jobs)
	wg.Wait()
	return out, nil
}
