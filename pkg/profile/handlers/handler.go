package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CBSktravers/hooli/pkg/profile"
	"github.com/CBSktravers/hooli/pkg/profile/models"
)

// ProfileResource manages endpoints for profile
type Handlers struct {
	logger  *log.Logger
	service profile.Service
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/create", h.Create)
}

func NewHandlers(logger *log.Logger, service profile.Service) *Handlers {
	return &Handlers{
		logger:  logger,
		service: service,
	}
}

// Create Endpoint that creates a profile
func (h Handlers) Create(w http.ResponseWriter, r *http.Request) {
	// check user credientaion were passed
	keys, ok := r.URL.Query()["UserName"]
	if !ok || len(keys[0]) < 1 {
		// Failed to get user
		log.Println("Url Param 'UserName' is missing")
	}
	user := keys[0]
	log.Println("Create profile called by user:", user)
	// autherization of user request has permission

	decoder := json.NewDecoder(r.Body)
	var profile models.Profile
	err := decoder.Decode(&profile)
	log.Printf("%s wants to create new Profile %s", user, profile)

	//handle error better
	if err != nil {
		// return and log response
		log.Println("Fialed to decode query", err)

	}
	// call create to add profile to database
	err = h.service.Create(profile)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Profile Created"))
}
