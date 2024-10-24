package event

import (
	"fmt"
	Rabbit "github.com/arshia-rgh/rabbit-helper-go/rabbit_helper"
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

	go func() {
		err := rabbit.Consume(routingKey, callback)
		if err != nil {
			return
		}
	}()
}
