package lru_cache

import (
	"testing"
)

// IncreaseAge
func TestIncreaseAge(t *testing.T) {
	//setup
	item := CacheItem{"somevalue", 0}

	//test
	item.IncreaseAge()

	//assertion
	if item.age != 1 {
		t.Error("didnt increase Age")
	}

}