package driver

import (
	"log"
	"net/http"

	"github.com/CBSktravers/hooli/pkg/profile/routers"
)

//StartWebServer starts the profile server
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
