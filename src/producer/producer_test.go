package producer

import (
"testing"
"time"
"math/rand"
)

func delaySecond(n time.Duration) {
	time.Sleep(n * time.Second)
}


func TestYo(t *testing.T) {
	rand.Seed( time.Now().UTC().UnixNano())
	for i := 0; i < 5; i++ {
		PublishRandomData()
		delaySecond(2)
	}
	t.Error("OMH")
}