package lru_cache

import (
	"testing"
	// "fmt"
	"reflect"
)

// Create
func TestCreate(t *testing.T) {
	//test
	cache := CreateCache(1)

	//assertion
	if cache.size != 1 {
		t.Error("Cache size incorrect, expected to be 1")
	}
	if reflect.TypeOf(cache.dictionary).Kind() != reflect.Map {
		t.Error("Cache dictionary should be type of Map")
	}
	if len(cache.dictionary) != 0 {
		t.Error("should initialize dictionary empty map")
	}
}

// GetValueForKey
func TestGetValueForKey(t *testing.T) {
	//setup
	cache := CreateCache(3)
	cache.StoreValueForKey("abcd", 1)
	cache.StoreValueForKey("xyz", 2)

	//get
	result, _ := cache.GetValueForKey(2)

	//assertion
	if result != "xyz" {
		t.Error("didnt return expected value")
	}
	if cache.dictionary[2].age != 1 {
		t.Error("didnt reset the age")
	}
}

func TestGetValueForKeyToRemoveLeastRecentlyUsedItems(t *testing.T) {
	//setup
	cache := CreateCache(2)
	cache.StoreValueForKey("a", 1)
	cache.StoreValueForKey("b", 2)
	cache.GetValueForKey(1)
	cache.StoreValueForKey("c", 3)

	//get
	result, error := cache.GetValueForKey(2)

	//assertion
	if error == nil && result != "" {
		t.Error("Expected to return an error")
	}
}

//StoreValueForKey
func TestStoreValueForKey(t *testing.T) {
	//setup
	cache := CreateCache(1)

	//test
	cache.StoreValueForKey("abcd", 1)

	//assertion
	if len(cache.dictionary) != 1 {
		t.Error("didnt store item successfully")
	}
	if cache.dictionary[1].value != "abcd" {
		t.Error("correct value wasn't stored")
	}
}

//StoreValueForKey
func TestDiscardLeastUsed(t *testing.T) {
	//setup
	cache := CreateCache(1)
	cache.StoreValueForKey("abcd", 1)

	//test
	cache.DiscardLeastUsed()

	//assertion
	if len(cache.dictionary) != 0 {
		t.Error("didnt delete least recently used item")
	}
}

//AgeCacheItems
func TestAgeCacheItems(t *testing.T) {
	//setup
	cache := CreateCache(1)
	cache.StoreValueForKey("abcd", 1)

	//test
	cache.AgeCacheItems();

	//assertion
	if cache.dictionary[1].age != 2 {
		t.Error("wasn't able to age cache items")
	}
}