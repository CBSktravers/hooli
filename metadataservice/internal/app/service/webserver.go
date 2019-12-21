package service

import (
	"log"
	"net/http"
	"os"
)

//StartWebServer starts a webserver at port
func StartWebServer(port string) {

	logger := log.New(os.Stdout, "hooli ", log.LstdFlags|log.Lshortfile)

	logger.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

	if err != nil {
		logger.Println("An error occured starting HTTP listener at port " + port)
		logger.Println("Error: " + err.Error())
	}
}
