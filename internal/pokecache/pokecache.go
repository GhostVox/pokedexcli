package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheMap map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
	stop     chan struct{}
}

func NewCache(interval time.Duration) *Cache {
	var newcache = Cache{
		cacheMap: make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
		stop:     make(chan struct{}),
		interval: interval,
	}
	go newcache.reapLoop()
	return &newcache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.cacheMap[key]; ok {
		return
	}
	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.cacheMap[key]; !ok {
		return nil, false
	}
	return c.cacheMap[key].val, true
}
func (c *Cache) Stop() {
	close(c.stop)
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		select {

		case <-ticker.C:
			c.mu.Lock()
			for key, element := range c.cacheMap {
				if t := time.Since(element.createdAt); t > c.interval {
					delete(c.cacheMap, key)
				}

			}
			c.mu.Unlock()
		case <-c.stop:
			return
		}

	}
}
