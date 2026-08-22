package starter

import (
	"context"
	"errors"
	"sync"
)

var ErrInvalidWorkers = errors.New("workers must be positive")

type SquareJob struct {
	Index int
	Input int
}

func SquareAll(ctx context.Context, values []int, workers int) ([]int, error) {

	if workers < 1 {
		return nil, ErrInvalidWorkers
	}

	results := make([]int, len(values))
	jobChan := make(chan SquareJob)
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-jobChan:
					if !ok {
						return
					}
					results[job.Index] = square(job.Input)
				}
			}
		}()
	}

produceLoop:
	for i, v := range values {
		select {
		case jobChan <- SquareJob{
			Index: i,
			Input: v,
		}:
		case <-ctx.Done():
			break produceLoop
		}
	}
	close(jobChan)
	wg.Wait()
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	return results, nil
}

func square(value int) int {
	return value * value
}
