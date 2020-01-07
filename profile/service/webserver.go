package service

import (
	"log"
	"net/http"

	"github.com/CBSktravers/hooli/profile/routers"
)

func StartWebServer(port string) {
	r := routers.NewRouter()
	http.Handle("/", r)
	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
