package server

import (
	"log"
	"net/http"
	"os"
	"time"
)

//NewServer returns a fully configured http.Server
func NewServer(mux *http.ServeMux, listenPort string) (*http.Server, error) {
	// will add error message as this function grows
	logger := log.New(os.Stdout, "hooli ", log.LstdFlags|log.Lshortfile)
	logger.Println("server starting")

	srv := &http.Server{
		Handler:      mux,
		Addr:         ":" + listenPort,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	logger.Println("server created")
	return srv, nil
}
