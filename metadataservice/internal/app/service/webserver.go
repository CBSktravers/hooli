package service

import (
	"log"
	"net/http"
	"os"
)

//StartWebServer starts a webserver at port
func StartWebServer(port string) {
	logger := log.New(os.Stdout, "hooli ", log.LstdFlags|log.Lshortfile)
	logger.Println("Starting metadata HTTP service at " + port)

	h := NewHandlers(logger)

	mux := http.NewServeMux()

	h.SetupRoutes(mux)

	err := http.ListenAndServe(":"+port, mux)

	if err != nil {
		logger.Println("An error occured starting HTTP listener at port " + port)
		logger.Println("Error: " + err.Error())
	}
}
