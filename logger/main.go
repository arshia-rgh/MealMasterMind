package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"logger-service/event"
	"time"
)

var DB *mongo.Database

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
	DB = mongoClient.Database("logs")

	event.Consume("logs", callback)

}

func callback(data any) {
	logData, ok := data.(Log)
	if !ok {
		log.Printf("invalid data type: %T", data)
		return
	}

	err := insertLog(DB, logData)

	if err != nil {
		log.Printf("error inserting log : %v", err)
	}
}
