package main

import (
	"context"
	"log"
	"time"
)

func main() {
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

	mongoClient, err := connectToMongoDB()

	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	db := mongoClient.Database("logs")

	err = consume("logs", ch, db)
	if err != nil {
		log.Println(err)
	}
}
