package cache

import (
	"context"

	lru "github.com/hashicorp/golang-lru"
	"github.com/nicksrandall/dataloader"
)

type Cache struct {
	arcCache *lru.ARCCache
}

func New(size int) (*Cache, error) {
	arcCache, err := lru.NewARC(size)
	if err != nil {
		return nil, err
	}

	cache := Cache{
		arcCache: arcCache,
	}

	return &cache, nil
}

func (c *Cache) Get(_ context.Context, key dataloader.Key) (dataloader.Thunk, bool) {
	v, ok := c.arcCache.Get(key)
	if ok {
		return v.(dataloader.Thunk), ok
	}
	return nil, ok
}

func (c *Cache) Set(_ context.Context, key dataloader.Key, value dataloader.Thunk) {
	c.arcCache.Add(key, value)
}

func (c *Cache) Delete(_ context.Context, key dataloader.Key) bool {
	if c.arcCache.Contains(key) {
		c.arcCache.Remove(key)
		return true
	}
	return false
}

func (c *Cache) Clear() {
	c.arcCache.Purge()
}
