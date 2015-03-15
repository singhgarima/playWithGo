package consumer

import (
	"github.com/streadway/amqp"
	"amqp_utilities"
	"fmt"
	"encoding/json"
)

var (
	connection *amqp.Connection
	channel *amqp.Channel
)

func ConsumeData (messages chan amqp_utilities.Message){
	var err error
	connection, err = amqp.Dial("amqp://guest:guest@localhost:5672")
	defer connection.Close()

	if (err == nil) {
		channel, err = connection.Channel()
		defer channel.Close()

		if (err == nil) {
			queue, err := amqp_utilities.CreateQueue(channel, "QUEUE")
			
			if ( err == nil ) {
				GetDeliveries(channel, queue, messages)
			} else {
				panic(fmt.Sprintf("%s", err))
			}
		} else {
			panic(fmt.Sprintf("%s", err))
		}
	} else {
		panic(fmt.Sprintf("%s", err))
	}
}

func GetDeliveries (channel *amqp.Channel, queue *amqp.Queue, messages chan amqp_utilities.Message) {
	deliveries, err := channel.Consume(queue.Name, "", true, false, false, false, nil)

	if ( err == nil ) {
		for d := range deliveries {
			var m amqp_utilities.Message
			json.Unmarshal(d.Body, &m)
			messages <- m
		}
	} else {
		fmt.Println("Consume: Error raised %s", err)
	}
}