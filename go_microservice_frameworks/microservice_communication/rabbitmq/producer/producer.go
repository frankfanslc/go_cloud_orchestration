package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/streadway/amqp"
)

func main()  {
	fmt.Println("Producing to RabbitMQ ...")
	time.Sleep(10*time.Second)

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

	msgCount := 0

	// Get signal of finish
	doneCh := make(chan struct{})

	go func() {
		for {
			msgCount++
			message := fmt.Sprintf("Message #%d", msgCount)

			// Pubhlish the message
			err := ch.Publish("",
				q.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:		 []byte(message),
				})

			log.Printf("'%s' message successfully published in RabbitMQ", message)
			failOnError(err, "Failed to publish the '" + message + "' message")

			time.Sleep(5 * time.Second)
		}
	}()

	<-doneCh
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
