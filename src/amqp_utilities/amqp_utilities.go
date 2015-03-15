package amqp_utilities

import (
"github.com/streadway/amqp"
"fmt"
)

type Message struct {
    Key string
    Body string
}

func CreateQueue(channel *amqp.Channel, queueName string) (*amqp.Queue, error) {
    queue, err := channel.QueueDeclare(
            queueName, // name of the queue
            true,      // durable
            false,     // delete when usused
            false,     // exclusive
            false,     // noWait
            nil,       // arguments
            )
    if err != nil {
        return nil, fmt.Errorf("Queue Declare: %s", err)
    }

    return &queue, nil
}