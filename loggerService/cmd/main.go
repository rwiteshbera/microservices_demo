package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPORT  = "80"
	rpcPORT  = "5001"
	mongoURL = "mongodb://localhost:27017"
	grpcPORT = "50001"
)

var client *mongo.Client

type Config struct {
}

func main() {
	// Connect to MONGODB
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	client = mongoClient

	// Create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Panic(err)
		}
	}()
}

func connectToMongo() (*mongo.Client, error) {
	// Create the connection option
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// Connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("error connecting : ", err.Error())
		return nil, err
	}
	return c, nil
}
