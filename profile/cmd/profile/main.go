package main

import (
	"fmt"

	"github.com/CBSktravers/hooli/profile/dbclient"
	"github.com/CBSktravers/hooli/profile/internal/app/service"
)

var appName = "profileservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("8080")
}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
