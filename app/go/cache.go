package main

import "sync"

type Cache struct {
	iconHashCache sync.Map // user_id -> iconHash
	themeCache    sync.Map // user_id -> *ThemeModel
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

func (c *Cache) getTheme(userID int64) *ThemeModel {
	if v, ok := c.themeCache.Load(userID); ok {
		return v.(*ThemeModel)
	}
	return nil
}

func (c *Cache) setTheme(userID int64, theme *ThemeModel) {
	c.themeCache.Store(userID, theme)
}
