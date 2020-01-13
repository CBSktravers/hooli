package routers

import (
	"net/http"

	"github.com/CBSktravers/hooli/pkg/profile/handlers"
)

// Defines a single route, e.g. a human readable name, HTTP method and the
// pattern the function that will execute when the route is called.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route

// Initialize our routes
var routes = Routes{

	Route{
		"CreateProfile",          // Name
		"POST",                   // HTTP method
		"/create",                // Route pattern
		handlers.Handlers.Create, // Endpoint function
	},
}
