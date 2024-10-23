package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type Log struct {
	Name  string `bson:"name"`
	Level string `bson:"level"`
	Data  string `bson:"data"`
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
