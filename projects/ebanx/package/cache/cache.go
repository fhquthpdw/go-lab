package cache

import (
	"errors"
	"sync"
	"time"
)

func NewCache(gcDuration int) *Cache {
	c := &Cache{
		items:      make(map[string]*CacheItem),
		gcDuration: gcDuration,
	}
	go c.GC()
	return c
}

type CacheItem struct {
	val       interface{}
	createdAt time.Time
	life      time.Duration // zero means always live
}

func (i *CacheItem) isExpire() bool {
	if i.life == 0 {
		return false
	}
	return time.Now().Sub(i.createdAt) > i.life
}

type Cache struct {
	sync.RWMutex
	items      map[string]*CacheItem
	gcDuration int
}

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

func (c *Cache) Set(key string, value interface{}, life time.Duration) error {
	c.Lock()
	defer c.Unlock()
	c.items[key] = &CacheItem{
		val:       value,
		createdAt: time.Now(),
		life:      life * time.Second,
	}
	return nil
}

func (c *Cache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()
	if _, ok := c.items[key]; !ok {
		return errors.New("key not exists")
	}
	delete(c.items, key)
	return nil
}

func (c *Cache) List() map[string]interface{} {
	r := make(map[string]interface{})
	for k, v := range c.items {
		r[k] = v.val
	}
	return r
}

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

func (c *Cache) DeleteExpired(keys []string) {
	c.Lock()
	defer c.Unlock()
	for _, key := range keys {
		delete(c.items, key)
	}
}

func (c *Cache) GC() {
	if c.gcDuration == 0 {
		return
	}

	for {
		select {
		case <-time.After(time.Duration(c.gcDuration) * time.Second):
			if keys := c.GetExpiredKeys(); len(keys) != 0 {
				for _, key := range keys {
					c.Delete(key)
				}
			}
		}
	}
}
