package event

import (
	"fmt"
	Rabbit "github.com/arshia-rgh/rabbit-helper-go/rabbit_helper"
	"log"
	"os"
)

func Consume(routingKey string, callback func(data any)) {
	rabbitMQURL := fmt.Sprintf("amqp://%v:%v@%v:%v/",
		os.Getenv("RABBITMQ_USERNAME"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)

	rabbit := Rabbit.New(rabbitMQURL)
	defer rabbit.Close()

	err := rabbit.Consume(routingKey, callback)

	if err != nil {
		log.Panic(err)
	}
}
