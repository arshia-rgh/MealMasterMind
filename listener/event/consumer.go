package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func consume(routingKeys []string, conn *amqp.Connection) error {
	ch, err := conn.Channel()

	if err != nil {
		return err
	}
	defer ch.Close()

	for _, v := range routingKeys {
		_, err := ch.QueueDeclare(
			v,
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return err
		}
	}

	for _, v := range routingKeys {
		msgs, err := ch.Consume(
			v,
			"",
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Println(err)
		}

		go func(routingKey string) {
			for d := range msgs {
				switch routingKey {
				case "send-mail":
					//
				case "log":
					//

				}
				d.Ack(false)
			}
		}(v)

	}
	log.Println("waiting for messages ")
	select {}

}
