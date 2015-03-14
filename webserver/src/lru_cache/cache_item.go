package lru_cache

type CacheItem struct { 
	value string
	age   int
}

func (item *CacheItem) ResetAge() {
	item.age = 0
}

func (item *CacheItem) IncreaseAge() {
	item.age = item.age + 1
}