package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/CBSktravers/hooli/pkg/profile"
	profileRepo "github.com/CBSktravers/hooli/pkg/profile/repository"
	"github.com/CBSktravers/hooli/pkg/profile/routers"
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

	//STOP using gorilla follow orignal to pass logger and service
	r := routers.NewRouter()
	http.Handle("/", r)
	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

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
