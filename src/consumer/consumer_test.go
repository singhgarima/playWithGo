package consumer

import (
"testing"
"time"
"fmt"
"amqp_utilities"
)

func delaySecond(n time.Duration) {
	time.Sleep(n * time.Second)
}


func TestYo(t *testing.T) {
	messages := make(chan amqp_utilities.Message)
	for i := 0; i < 5; i++ {
		fmt.Println("loop", i)

		ConsumeData(messages)
		fmt.Println("loop end", i)
		delaySecond(2)
	}
	
	for j := range messages {
		fmt.Println(j)
	}
	t.Error("Just Like that")
}