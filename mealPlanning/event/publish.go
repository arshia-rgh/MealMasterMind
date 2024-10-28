package event

import (
	"fmt"
	"log"
	"os"

	Rabbit "github.com/arshia-rgh/rabbit-helper-go/rabbit_helper"
)

func Publish(routingKey string, data any) {
	rabbitMQURL := fmt.Sprintf("amqp://%v:%v@%v:%v/",
		os.Getenv("RABBITMQ_USERNAME"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)

	rabbit := Rabbit.New(rabbitMQURL)
	defer rabbit.Close()

	go func() {
		err := rabbit.Publish(routingKey, data)
		if err != nil {
			log.Printf("error publishing mssage to the %v queue, message: %v", routingKey, data)

		}
	}()

}
