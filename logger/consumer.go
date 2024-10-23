package main

import (
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math"
	"os"
	"time"
)

func consume(routingKey string, ch *amqp.Channel, db *mongo.Database) error {
	queue, err := ch.QueueDeclare(
		routingKey,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for {
		msgs, err := ch.Consume(
			queue.Name,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Println(err)
			continue
		}

		for msg := range msgs {
			var logData Log
			err := json.Unmarshal(msg.Body, &logData)
			if err != nil {
				log.Printf("error unmarshalling the message: %v\n", err)
			}
			go func() {
				err := insertLog(db, logData)
				if err != nil {
					log.Printf("error inserting log: %v\n", err)
				}
			}()
		}
	}
}

func connect() (*amqp.Connection, error) {
	var connection *amqp.Connection
	var counts int64
	var backOff = 1 * time.Second

	rabbitMQURL := fmt.Sprintf("amqp://%v:%v@%v:%v/",
		os.Getenv("RABBITMQ_USERNAME"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)

	for {
		c, err := amqp.Dial(rabbitMQURL)
		if err != nil {
			log.Println("rabbitmq not yet ready...!")
			counts++
		} else {
			log.Println("connected to RabbitMQ")
			connection = c
			break
		}

		if counts > 5 {
			log.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Printf("backing off for %v seconds\n", backOff)
		time.Sleep(backOff)

	}

	return connection, nil
}
