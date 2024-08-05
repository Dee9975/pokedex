package cache

import (
	"time"
)

type Cache struct {
	Caches map[string]*Entry
}

type Entry struct {
	Key       string
	CreatedAt time.Time
	Value     []byte
}

func NewCache() *Cache {
	return &Cache{
		Caches: make(map[string]*Entry),
	}
}

func (c *Cache) Get(key string) *Entry {
	return c.Caches[key]
}

func (c *Cache) Set(key string, value []byte) {
	c.Caches[key] = &Entry{
		Key:       key,
		CreatedAt: time.Now(),
		Value:     value,
	}
}

func (c *Cache) ClearOld() {
	for _, entry := range c.Caches {
		if entry.CreatedAt.Add(30*time.Second).Compare(time.Now()) < 0 {
			delete(c.Caches, entry.Key)
		}
	}
}
