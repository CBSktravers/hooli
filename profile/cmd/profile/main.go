package main

import (
	"fmt"

	"github.com/CBSktravers/hooli/profile/internal/app/service"
)

var appName = "profileservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("8080")
}
