package handlers

import (
	"log"
	"net/http"

	"github.com/CBSktravers/hooli/pkg/profile"
	"github.com/CBSktravers/hooli/pkg/profile/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

// ProfileResource manages endpoints for profile
type Handlers struct {
	logger  *log.Logger
	service profile.Service
	context buffalo.Context
}

// Create Endpoint that creates a profile

func (h Handlers) Create(c buffalo.Context) error {

	// Log users request
	// check permissons now or early?
	log.Println("Create Profile called by user:")

	//decoder := json.NewDecoder(r.Body)
	//var profile models.Profile
	//err := decoder.Decode(&profile)

	// Varify all input is there and correctly formated

	//handle error better
	//if err != nil {
	//	panic(err)
	//}

	var profile models.Profile

	// call create profile
	respObj := h.service.Create(profile)
	// return and log response
	return c.Render(http.StatusCreated, render.JSON(respObj))
}
