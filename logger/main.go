package main

import (
	"context"
	"log"
	"time"
)

func main() {
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

}
