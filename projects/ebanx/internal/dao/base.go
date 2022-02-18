package dao

import "ebanx/package/cache"

var db *cache.Cache

type BaseDao struct {
	db *cache.Cache
}

func (b BaseDao) Create(key string, value interface{}) error {
	return b.db.Set(key, value, 0)
}

func (b BaseDao) Update(key string, value interface{}) error {
	return b.db.Set(key, value, 0)
}

func (b BaseDao) Delete(key string) error {
	return b.db.Delete(key)
}

func init() {
	db = cache.NewCache(0)
}
