package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDataBaseClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin2:admin2@ds119020.mlab.com:19020/metrondb"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Creating Connection")

	// Create Connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Checking Connection")

	// Check the connections
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("you're already connected to Metron DB")
	return client

}
