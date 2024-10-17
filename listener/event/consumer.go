package main

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func consumeAndReplyQueue(queue string, ch *amqp.Channel) {
	msgs, err := ch.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer for queue %s: %v", queue, err)
	}

	for msg := range msgs {
		log.Printf("Received a message from queue %s: %s", queue, msg.Body)

		var url, method string
		switch queue {
		case "login":
			url, method = "http://auth-service/api/login/", "POST"
		case "register":
			//

		default:
			//
		}

		resPayload, err := requestToService(url, msg, method)
		if err != nil {
			log.Printf("Error making request to service: %v", err)
			continue
		}

		jsonResponse, _ := json.Marshal(resPayload)

		err = ch.Publish(
			"",
			msg.ReplyTo,
			false,
			false,
			amqp.Publishing{
				ContentType:   "application/json",
				CorrelationId: msg.CorrelationId,
				Body:          jsonResponse,
			},
		)

		if err != nil {
			log.Printf("Error publishing response: %v", err)

		}

	}
}
