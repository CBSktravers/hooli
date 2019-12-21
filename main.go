package main

import (
	"log"
	"net/http"
	"os"
)

var message = "Hello Hooli!!!"

func main() {
	logger := log.New(os.Stdout, "hooli ", log.LstdFlags|log.Lshortfile)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})
	logger.Println("server starting")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
