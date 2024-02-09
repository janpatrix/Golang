package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"engineeringleader.de/listener/event"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	//connect to rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	//listen for messages
	log.Println("Listening and processing RabbitMQ messages...")

	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	//consumer
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}

	//observe queue and process events
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	//wait for rabbit
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			log.Println("Connected to RabbitMQ")
			connection = c
			break
		}
		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}
