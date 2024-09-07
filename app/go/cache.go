package main

import "sync"

type Cache struct {
	iconHashCache sync.Map // user_id -> iconHash
}

var cache Cache

func (c *Cache) getIconHash(userID int64) []byte {
	if v, ok := c.iconHashCache.Load(userID); ok {
		return v.([]byte)
	}
	return nil
}

func (c *Cache) setIconHash(userID int64, image []byte) {
	c.iconHashCache.Store(userID, image)
}
