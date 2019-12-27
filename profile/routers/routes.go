package routers

import (
	"net/http"
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
		"CreateProfile", // Name
		"POST",          // HTTP method
		"/profile",      // Route pattern
		CreateProfile,   // Endpoint function
	},
	Route{
		"UpdateProfile", // Name
		"POST",          // HTTP method
		"/profile",      // Route pattern
		UpdateProfile,   // Endpoint function
	},
	Route{
		"DeleteProfile", // Name
		"POST",          // HTTP method
		"/profile",      // Route pattern
		DeleteProfile,   // Endpoint function
	},
	Route{
		"Getrofile", // Name
		"GET",       // HTTP method
		"/profile",  // Route pattern
		GetProfile,  // Endpoint function
	},
	Route{
		"GetProfiles", // Name
		"GET",         // HTTP method
		"/profiles",   // Route pattern
		GetProfiles,   // Endpoint function
	},
}
