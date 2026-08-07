package solution_test

import (
	target "bootcamp/10-caching/solution"
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestCacheLoadsOnce(t *testing.T) {
	c := target.New()
	var calls atomic.Int32
	load := func(context.Context) (string, error) { calls.Add(1); time.Sleep(time.Millisecond); return "value", nil }
	var wg sync.WaitGroup
	for range 8 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			got, err := c.Get(context.Background(), "k", load)
			if err != nil || got != "value" {
				t.Errorf("got (%q,%v)", got, err)
			}
		}()
	}
	wg.Wait()
	if calls.Load() != 1 {
		t.Fatalf("loads %d", calls.Load())
	}
}
