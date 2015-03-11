package lru_cache

import (
	"testing"
	// "fmt"
	"reflect"
)

func TestCreate(t *testing.T) {
	cache := CreateCache(64)

	if cache.size != 64 {
		t.Error("Cache size incorrect, expected to be 64")
	}
	if reflect.TypeOf(cache.dictionary).Kind() != reflect.Map {
		t.Error("Cache dictionary should be type of Map")
	}
	if len(cache.dictionary) != 0 {
		t.Error("should initialize dictionary empty map")
	}
}

func TestGetValueForKey(t *testing.T) {
	
}

func TestStoreValueForKey(t *testing.T) {
	cache := CreateCache(1024)
	StoreValueForKey("abcd", 1)

	if len(cache.dictionary) != 1 {
		t.Error("didnt store item successfully")
	}
	if cache.dictionary[1].value != "abcd" {
		t.Error("correct value wasn't stored")
	}
}