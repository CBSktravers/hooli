package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CBSktravers/hooli/pkg/profile"
	"github.com/CBSktravers/hooli/pkg/profile/driver"
	"github.com/CBSktravers/hooli/pkg/profile/handlers"
	profileRepo "github.com/CBSktravers/hooli/pkg/profile/repository"
)

var appName = "profile service"

func main() {
	fmt.Printf("Starting %v\n", appName)
	StartWebServer("8080")
}

//StartWebServer starts the profile server
func StartWebServer(port string) {

	// Get a handle for your collection
	collection := driver.CreateClient()
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
