package handlers

import (
	"context"
	"log"

	"github.com/CBSktravers/hooli/profile/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createSession() *mongo.Client {
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

func addProfile() []models.Profile {
	var profiles []models.Profile
	session := createSession()
	// Get a handle for your collection
	collection := session.Database("mongodb").Collection("profile")
	// Some dummy data to add to the Database
	test := models.Profile{"test_name", "test", map[string]string{
		"provider_id": "test.com",
		"provider":    "TEST",
	}}

	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), test)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted a single document: ", insertResult.InsertedID)

	return profiles
}

//updateProfile
//addProfiles
//deleteProfile
//deleteProfiles
//getProfile
//getProfiles
