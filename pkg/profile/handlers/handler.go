package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CBSktravers/hooli/pkg/profile"
	"github.com/CBSktravers/hooli/pkg/profile/models"
)

// ProfileResource manages endpoints for profile
type ProfileResource struct {
	profile.Service
}

// Create Endpoint that creates a profile
func (r ProfileResource) Create(w http.ResponseWriter, req *http.Request) {
	// Log users request
	// check permissons now or early?
	log.Println("Create Profile called by user:")

	decoder := json.NewDecoder(req.Body)
	var profile models.Profile
	err := decoder.Decode(&profile)

	// Varify all input is there and correctly formated

	//handle error better
	if err != nil {
		panic(err)
	}

	// call create profile
	r.Service.Create(&profile)
	// return and log response
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Profile Created"))
}
