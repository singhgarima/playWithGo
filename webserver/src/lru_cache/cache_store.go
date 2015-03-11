package lru_cache

type CacheItem struct { 
	value string
	age   int
}

type CacheStore struct {
	size       int
	dictionary map[int]CacheItem
}

// func (c CacheStore) GetValueForKey(key int) string {
// 	return c.dictionary[key].value
// }

// func (c CacheStore) StoreValueForKey(value string, key int) {
// 	fmt.Println(c)
// 	c.dictionary[key] = CacheItem{value, 0}
// }
