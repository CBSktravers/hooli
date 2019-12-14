package main

import (
	"log"
	"net/http"
	"github.com/CBSktravers/hooli/server"
)

var (
	HooliCertFile    = "./certs/local.localhost.cert" //os.Getenv("Hooli_CERT_FILE")
	HooliKeyFile     = "./certs/local.localhost.key"  //os.Getenv("Hooli_KEY_FILE")
	HooliServiceAddr = ":8080"                        //os.Getenv("Hooli_SERVICE_ADDR")
)

const message = "Hello Wolrd"

//const HooliServiceAddr = ":8080"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})

	srv := server.New(mux, HooliServiceAddr)

	err := srv.ListenAndServeTLS(HooliCertFile, HooliKeyFile)
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

