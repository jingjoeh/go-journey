package starter

import (
	"context"
	"sync"
)

type Cache struct {
	mu     sync.Mutex
	values map[string]string
}

func New() *Cache { return &Cache{values: make(map[string]string)} }

func (c *Cache) Get(ctx context.Context, key string, load func(context.Context) (string, error)) (string, error) {
	panic("TODO: implement Cache.Get")
}
