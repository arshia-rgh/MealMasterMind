package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type Log struct {
	Name      string    `bson:"name"`
	Level     string    `bson:"level"`
	Data      string    `bson:"data"`
	CreatedAt time.Time `bson:"created_at"`
}

func insertLog(db *mongo.Database, logData Log) error {
	collection := db.Collection("logs")

	_, err := collection.InsertOne(context.TODO(), logData)
	if err != nil {
		log.Println("failed to insert into logs", err)
		return err
	}

	return nil
}
func connectToMongoDB() (*mongo.Client, error) {
	dbUser := os.Getenv("MONGO_USER")
	dbPassword := os.Getenv("MONGO_PASSWORD")
	dbHost := os.Getenv("MONGO_HOST")
	dbPort := os.Getenv("MONGO_PORT")
	dbName := os.Getenv("MONGO_DB")
	var uri string
	if dbUser != "" {
		uri = fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", dbUser, dbPassword, dbHost, dbPort, dbName)
	} else {
		uri = fmt.Sprintf("mongodb://%v:%v/%v", dbHost, dbPort, dbName)
	}

	clientOptions := options.Client().ApplyURI(uri)

	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB !")

	return mongoClient, nil
}
