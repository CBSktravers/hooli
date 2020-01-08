package handlers

import (
	"context"
	"fmt"
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

func AddProfileMongo() {
	var profiles []models.Profile
	client := createClient()

	// Get a handle for your collection
	collection := client.Database("mongodb").Collection("profile")

	// Some dummy data to add to the Database
	comcast := models.Profile{"Comcast", "CBSVOD", map[string]string{
		"provider_id": "cbs.com",
		"provider":    "comcast",
	}}
	verizon := models.Profile{"Verizon", "CBSVOD", map[string]string{
		"provider_id": "cbs.com",
		"provider":    "verizon",
	}}
	dish := models.Profile{"Dish", "CWVOD", map[string]string{
		"provider_id": "cw.com",
		"provider":    "dish",
	}}
	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), comcast)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Insert multiple documents
	profiles := []interface{}{dish, verizon}

	insertManyResult, err := collection.InsertMany(context.TODO(), profiles)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

}

//updateProfile
//addProfiles
//deleteProfile
//deleteProfiles
//getProfile
//getProfiles
