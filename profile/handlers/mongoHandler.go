package handlers

import (
	"context"
	"log"

	"github.com/CBSktravers/hooli/profile/models"
	"go.mongodb.org/mongo-driver/bson"
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

func createMongoProfile(profile models.Profile) {
	// Establish client to mongodabase
	client := createClient()

	// Get a handle for your collection
	collection := client.Database("mongodb").Collection("profile")

	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), profile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted a single document: ", insertResult.InsertedID)

	//Close client
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB closed.")
}

func getMongoProfile(profile models.Profile) {
	// Establish client to mongodabase
	client := createClient()

	// Get a handle for your collection
	collection := client.Database("mongodb").Collection("profile")

	var result models.Profile

	// Create filter to find item in database
	filter := bson.D{{"name", profile.Name}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Found a single document: %+v\n", result)
}

//updateProfile
//addProfiles
//deleteProfile
//deleteProfiles
//getProfile
//getProfiles
