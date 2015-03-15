package lru_cache

import (
	"errors"
	// "fmt"
)

type CacheStore struct {
	size       int
	dictionary map[int]*CacheItem
}

func CreateCache(size int) CacheStore {
	dictionary := make(map[int]*CacheItem, size)
	cacheStore := CacheStore{size, dictionary}
	return cacheStore
}

func (c *CacheStore) GetValueForKey(key int) (string, error) {
	item, present := c.dictionary[key]
	if present {
		c.dictionary[key].ResetAge()
		c.AgeCacheItems()
		return item.value, nil
	}
	return "", errors.New("Cannot find value for key")
}

func (c *CacheStore) StoreValueForKey(value string, key int) {
	if len(c.dictionary) < c.size {
		c.dictionary[key] = &CacheItem{value, 0}
		c.AgeCacheItems()
	} else {
		c.DiscardLeastUsed()
		c.StoreValueForKey(value, key)
	}
}

func (c *CacheStore) AgeCacheItems() {
	for _, item := range c.dictionary {
		item.IncreaseAge()
    }
}

func (c *CacheStore) DiscardLeastUsed() {
	var maxAge int = 0
	var deletionCandidate int
	for key, value := range c.dictionary {
        if (value.age >= maxAge) {
        	maxAge = value.age
        	deletionCandidate = key
        }
    }
    delete(c.dictionary, deletionCandidate)
}