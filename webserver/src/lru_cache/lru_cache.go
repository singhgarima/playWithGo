package lru_cache

var cacheStore *CacheStore

func CreateCache(size int) CacheStore {
	dictionary := make(map[int]CacheItem, size)
	cacheStore := CacheStore{size, dictionary}
	return cacheStore
}

func GetValueForKey(key int) string {
	ResetAge(key)
	return cacheStore.dictionary[key].value
}

func StoreValueForKey(value string, key int) {
	if len(cacheStore.dictionary) < cacheStore.size {
		cacheStore.dictionary[key] = CacheItem{value, 0}
		ResetAge(key)
	} else {
		DiscardLeastUsed()
	}
	
}

func AgeCacheItems(c CacheStore) {
	for _, item := range c.dictionary {
        item.age += 1
    }
}


func ResetAge(key int) {
	item := cacheStore.dictionary[key]
	item.age = 0
}


func DiscardLeastUsed() {
	maxAge := 0
	var deletionCandidate int
	for key, value := range cacheStore.dictionary {
        if (value.age > maxAge) {
        	maxAge = value.age
        	deletionCandidate = key
        }
    }
    delete(cacheStore.dictionary, deletionCandidate)
}