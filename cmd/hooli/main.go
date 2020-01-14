package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CBSktravers/hooli/pkg/profile"
	"github.com/CBSktravers/hooli/pkg/profile/handlers"
	profileRepo "github.com/CBSktravers/hooli/pkg/profile/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var appName = "profile service"

func main() {
	fmt.Printf("Starting %v\n", appName)
	StartWebServer("8080")
}

//StartWebServer starts the profile server
func StartWebServer(port string) {
	// Establish client to mongodabase
	client := createClient()

	// Get a handle for your collection
	collection := client.Database("mongodb").Collection("profile")
	profileSvc := profile.NewDefaultService(profileRepo.NewMongo(collection))

	logger := log.New(os.Stdout, "hooli ", log.LstdFlags|log.Lshortfile)
	logger.Println("Starting metadata HTTP service at " + port)

	h := handlers.NewHandlers(logger, profileSvc)

	mux := http.NewServeMux()

	h.SetupRoutes(mux)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
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
