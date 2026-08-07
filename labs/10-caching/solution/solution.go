package solution

import (
	"context"
	"sync"
)

type entry struct {
	ready chan struct{}
	value string
	err   error
}
type Cache struct {
	mu      sync.Mutex
	values  map[string]string
	pending map[string]*entry
}

func New() *Cache { return &Cache{values: map[string]string{}, pending: map[string]*entry{}} }
func (c *Cache) Get(ctx context.Context, key string, load func(context.Context) (string, error)) (string, error) {
	c.mu.Lock()
	if value, ok := c.values[key]; ok {
		c.mu.Unlock()
		return value, nil
	}
	if current, ok := c.pending[key]; ok {
		c.mu.Unlock()
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-current.ready:
			return current.value, current.err
		}
	}
	current := &entry{ready: make(chan struct{})}
	c.pending[key] = current
	c.mu.Unlock()
	current.value, current.err = load(ctx)
	c.mu.Lock()
	if current.err == nil {
		c.values[key] = current.value
	}
	delete(c.pending, key)
	close(current.ready)
	c.mu.Unlock()
	return current.value, current.err
}
