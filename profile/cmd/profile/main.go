package main

import (
	"fmt"

	"github.com/CBSktravers/hooli/profile/service"
)

var appName = "profileservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("8080")
}
