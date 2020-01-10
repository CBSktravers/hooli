package repository

import (
	"context"
	"log"

	"github.com/CBSktravers/hooli/profile/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createClient() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	return client
}

// Create add a profile to mongodb
func Create(profile models.Profile) error {
	// Establish client to mongodabase
	client := createClient()

	// Get a handle for your collection
	collection := client.Database("mongodb").Collection("profile")

	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), profile)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Inserted a single document: ", insertResult.InsertedID)

	//Close client
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Connection to MongoDB closed.")

	return nil
}
