package main

import (
	"fmt"

	"github.com/CBSktravers/hooli/metadataservice/internal/app/service"
)

var appName = "metadataservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
}
