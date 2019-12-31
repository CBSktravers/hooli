package service

import (
	"log"
	"net/http"

	"github.com/CBSktravers/hooli/profile/common"
	"github.com/CBSktravers/hooli/profile/routers"
)

func StartWebServer(port string) {
	//set up database config
	common.InitConfig()
	r := routers.NewRouter()
	http.Handle("/", r)
	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
