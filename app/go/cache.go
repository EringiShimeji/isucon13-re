package main

import "sync"

type Cache struct {
	idIconHash   sync.Map // user_id -> iconHash
	nameIconHash sync.Map // username -> iconHash
	theme        sync.Map // user_id -> *ThemeModel
}

var cache Cache

func (c *Cache) getIconHashById(userID int64) (string, bool) {
	if v, ok := c.idIconHash.Load(userID); ok {
		return v.(string), true
	}
	return "", false
}

func (c *Cache) setIconHashWithId(userID int64, hash string) {
	c.idIconHash.Store(userID, hash)
}

func (c *Cache) getIconHashByName(username string) (string, bool) {
	if v, ok := c.nameIconHash.Load(username); ok {
		return v.(string), true
	}
	return "", false
}

func (c *Cache) setIconHashWithName(username string, hash string) {
	c.nameIconHash.Store(username, hash)
}

func (c *Cache) getTheme(userID int64) *ThemeModel {
	if v, ok := c.theme.Load(userID); ok {
		return v.(*ThemeModel)
	}
	return nil
}

func (c *Cache) setTheme(userID int64, theme *ThemeModel) {
	c.theme.Store(userID, theme)
}
