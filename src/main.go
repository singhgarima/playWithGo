package main 

import (
"server"
"time"
"producer"
"consumer"
"math/rand"
"lru_cache"
"fmt"
"sync"
"amqp_utilities"
"strconv"
)

var cache lru_cache.CacheStore

func delaySecond(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func main() {
	var waitGroup sync.WaitGroup
    waitGroup.Add(4)

    messages := make(chan amqp_utilities.Message)

	rand.Seed( time.Now().UTC().UnixNano())
	CreateCache()
	
	go func() {
		defer waitGroup.Done()
		fmt.Println("Starting http server")
		server.Start(&cache)
	}()

	go func() {
		defer waitGroup.Done()
		fmt.Println("Consuming Messages")
		consumer.ConsumeData(messages)
	}()

	go func() {
		defer waitGroup.Done()
		fmt.Println("Producing Data")
		DataBurst()
	}()

	go func() {
		defer waitGroup.Done()
		fmt.Println("Converting message to cache")
		for msg := range messages {
			key, _ := strconv.Atoi(msg.Key)
			cache.StoreValueForKey(msg.Body, key)
			fmt.Println(key)
		}
	}()

	fmt.Println("Waiting To Finish")
    waitGroup.Wait()

    fmt.Println("\nTerminating Program")
}

func CreateCache () {
	cache = lru_cache.CreateCache(10)
	cache.StoreValueForKey("some value", 1)
}

func DataBurst () {
	interval := rand.Intn(10)
	data := rand.Intn(100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < data; j++ {
			producer.PublishRandomData()
		}
		delaySecond(interval)
	}	
}
