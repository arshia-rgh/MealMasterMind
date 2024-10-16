package event

import (
	"encoding/json"
	"log"
)

func Consume(replyId string) map[string]interface{} {
	conn, err := connect()
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		replyId,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer for queue %s: %v", replyId, err)
	}

	for msg := range msgs {
		if msg.CorrelationId == replyId {
			var response map[string]interface{}
			_ = json.Unmarshal(msg.Body, &response)

			return response
		}

	}

	return nil
}
