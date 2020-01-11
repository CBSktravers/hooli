package main

import (
	"fmt"

	"github.com/CBSktravers/hooli/pkg/profile/driver"
)

var appName = "profile service"

func main() {
	fmt.Printf("Starting %v\n", appName)
	driver.StartWebServer("8080")
}
