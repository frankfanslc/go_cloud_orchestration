package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/streadway/amqp"
)

func main()  {
	fmt.Println("Consuming from RabbitMQ ...")
	time.Sleep(10 * time.Second)

	// Dial connection to the RabbitMQ broker
	conn, err := amqp.Dial(rabbitMQhost())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Open a communication channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a communication channel to RabbitMQ")
	defer ch.Close()

	// Declare the queue
	q, err := ch.QueueDeclare(
		queue(), false, false,
		false, false, nil)

	failOnError(err, "Failed to declare the queue")

	// Register and get a message consumer
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
		)
	failOnError(err, "failed to register a consumer")

	forever := make(chan bool)

	go func() {
		// Processing the messages
		for msg := range msgs {
			log.Printf("Received the '%s' message", msg.Body)
		}
	}()
	log.Printf("Waiting for messages ...")
	<-forever
}

func rabbitMQhost() string {
	brokerAddr := os.Getenv("BROKER_ADDR")
	if len(brokerAddr) == 0 {
		brokerAddr = "amqp://guest:guest@localhost:5672/"
	}
	return brokerAddr
}

func queue() string {
	queue := os.Getenv("QUEUE")
	if len(queue) == 0 {
		queue = "default-queue"
	}
	return queue
}

func failOnError(err error, message string) {
	if err != nil {
		fmt.Println(message)
	}
}