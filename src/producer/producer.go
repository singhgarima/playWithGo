package producer

import (
	"github.com/streadway/amqp"
	"amqp_utilities"
	"math/rand"
	"strconv"
	"fmt"
	"encoding/json"
)

var (
	connection *amqp.Connection
	channel *amqp.Channel
)

func PublishRandomData () error {
	var err error
	connection, err = amqp.Dial("amqp://guest:guest@localhost:5672")
	defer connection.Close()
	
	if (err == nil) {
		channel, err = connection.Channel()
		defer channel.Close()

		if (err == nil) {
			queue, err := amqp_utilities.CreateQueue(channel, "QUEUE")
			if ( err == nil ) {
				PostDeliveries(channel, queue)
			} else {
				return err
			}
		} else {
			return err
		}
	} else {
		return err
	}
	return nil
}

func PostDeliveries (channel *amqp.Channel, queue *amqp.Queue) {
	msg := amqp_utilities.Message{strconv.Itoa(rand.Intn(1000)), strconv.Itoa(rand.Intn(100000))}
	body, _ := json.Marshal(msg)

	publish_err := channel.Publish("", queue.Name, false, false, amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            body,
			DeliveryMode:    amqp.Transient,
			Priority:        0,
		})
	if (publish_err != nil) {
		fmt.Print("Producer: Failed to publish to Queue")
		fmt.Println(publish_err)
	}
}

