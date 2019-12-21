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

	err := http.ListenAndServe(":"+port, mux) // Goroutine will block here

	if err != nil {
		logger.Println("An error occured starting HTTP listener at port " + port)
		logger.Println("Error: " + err.Error())
	}
}

//StartWebServer starts a webserver at port
func StartWebServerOld(port string) {
	logger := log.New(os.Stdout, "hooli ", log.LstdFlags|log.Lshortfile)
	logger.Println("Starting metadata HTTP service at " + port)

	r := NewRouter()    // NEW
	http.Handle("/", r) // NEW

	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

	if err != nil {
		logger.Println("An error occured starting HTTP listener at port " + port)
		logger.Println("Error: " + err.Error())
	}
}
