package main

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"listener/mails"
	"log"
)

type MailRequest struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Link    string `json:"link"`
}

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
					var mailReq MailRequest
					if err := json.Unmarshal(d.Body, &mailReq); err != nil {
						log.Println("Failed to unmarshal JSON", err)
						continue
					}
					if err := mails.RequestToSendMailGRPC(mailReq.Email, mailReq.Subject, mailReq.Link); err != nil {
						log.Println("Failed to send mail:", err)
					}
				case "log":
					//TODO gRPC to the logger service

				}
				d.Ack(false)
			}
		}(v)

	}
	log.Println("waiting for messages ")
	select {}

}
