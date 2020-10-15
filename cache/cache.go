package cache

import (
	"github.com/l552121229/golang-tools/hash"
	"github.com/l552121229/golang-tools/hash/fnv64a"
)

var minShards = 1024

// Cache Cache
type Cache struct {
	shards []*Shard
	hash   fnv64a.Hash
}

// NewCache 新建缓存
func NewCache() *Cache {
	cache := &Cache{
		hash:   hash.GetFnv64a(),
		shards: make([]*Shard, minShards),
	}
	for i := 0; i < minShards; i++ {
		cache.shards[i] = initNewShard()
	}

	return cache
}

// getShard 获取Shard存储
func (c *Cache) getShard(hashedKey uint64) (shard *Shard) {
	return c.shards[hashedKey&uint64(minShards-1)]
}

// Set Set
func (c *Cache) Set(key string, value []byte) {
	hashedKey := c.hash.Sum64(key)
	shard := c.getShard(hashedKey)
	shard.set(hashedKey, value)
}

// Get Get
func (c *Cache) Get(key string) ([]byte, error) {
	hashedKey := c.hash.Sum64(key)
	shard := c.getShard(hashedKey)
	return shard.get(key, hashedKey)
}

// Del Del
func (c *Cache) Del(key string) bool {
	hashedKey := c.hash.Sum64(key)
	shard := c.getShard(hashedKey)
	return shard.del(hashedKey)
}

// GetShard GetShard
func (c *Cache) GetShard(key string) *Shard {
	return c.getShard(c.hash.Sum64(key))
}
