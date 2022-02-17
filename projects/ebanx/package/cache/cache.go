package cache

import (
	"errors"
	"sync"
	"time"
)

// init a cache instance
func NewCache(gcDuration int) *Cache {
	c := &Cache{
		items:      make(map[string]*CacheItem),
		gcDuration: gcDuration,
	}
	go c.GC()
	return c
}

// Cache value item
type CacheItem struct {
	val       interface{}
	createdAt time.Time
	life      time.Duration // zero means always live
}

// check is expired
func (i *CacheItem) isExpire() bool {
	if i.life == 0 {
		return false
	}
	return time.Now().Sub(i.createdAt) > i.life
}

// Cache
type Cache struct {
	sync.RWMutex
	items      map[string]*CacheItem
	gcDuration int
}

// get cache value
// return nil if expired or not exists
func (c *Cache) Get(key string) interface{} {
	c.RLock()
	defer c.RUnlock()
	if value, ok := c.items[key]; ok {
		if value.isExpire() == false {
			return value.val
		}
	}
	return nil
}

// set cache value
func (c *Cache) Set(key string, value interface{}, life time.Duration) error {
	c.Lock()
	defer c.Unlock()
	c.items[key] = &CacheItem{
		val:       value,
		createdAt: time.Now(),
		life:      time.Duration(life) * time.Second,
	}
	return nil
}

// delete cache value
func (c *Cache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()
	if _, ok := c.items[key]; !ok {
		return errors.New("key not exists")
	}
	delete(c.items, key)
	return nil
}

// get all cache items
func (c *Cache) Items() map[string]*CacheItem {
	return c.items
}

// get all expired keys for GC
func (c *Cache) GetExpiredKeys() (keys []string) {
	c.RLock()
	defer c.RUnlock()
	for key, item := range c.items {
		if item.isExpire() {
			keys = append(keys, key)
		}
	}
	return keys
}

// delete all expired keys for GC
func (c *Cache) DeleteExpired(keys []string) {
	c.Lock()
	defer c.Unlock()
	for _, key := range keys {
		delete(c.items, key)
	}
}

// GC every body knows
func (c *Cache) GC() {
	for {
		select {
		case <-time.After(time.Duration(c.gcDuration) * time.Second):
			if keys := c.GetExpiredKeys(); len(keys) != 0 {
				for _, key := range keys {
					c.Delete(key)
				}
			}
			//fmt.Println("GC")
		}
	}
}
