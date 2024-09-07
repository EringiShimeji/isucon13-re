package main

import "sync"

type Cache struct {
	idIconHash    sync.Map // user_id -> iconHash
	nameIconHash  sync.Map // username -> iconHash
	theme         sync.Map // user_id -> *ThemeModel
	idIconImage   sync.Map // user_id -> image
	nameIconImage sync.Map // username -> image
	livestream    sync.Map // id -> LivestreamModel

	iconID int64
}

var cache Cache

func getOrInsertMap[V interface{}](m *sync.Map, key any, f func() (V, error)) (V, error) {
	if v, ok := m.Load(key); ok {
		return v.(V), nil
	}
	v, err := f()
	if err == nil {
		m.Store(key, v)
	}
	return v, err
}

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

func (c *Cache) getIconImageById(userID int64) []byte {
	if v, ok := c.nameIconImage.Load(userID); ok {
		return v.([]byte)
	}
	return nil
}

func (c *Cache) setIconImageById(userID int64, image []byte) {
	c.nameIconImage.Store(userID, image)
	c.iconID++
}

func (c *Cache) getIconImageByName(username string) []byte {
	if v, ok := c.nameIconImage.Load(username); ok {
		return v.([]byte)
	}
	return nil
}

func (c *Cache) setIconImageByName(username string, image []byte) {
	c.nameIconImage.Store(username, image)
	c.iconID++
}
