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
	// Log users request
	// check permissons now or early?
	log.Println("Create Profile called by user:")

	decoder := json.NewDecoder(r.Body)
	var profile models.Profile
	err := decoder.Decode(&profile)
	log.Println(profile.Name)
	// Varify all input is there and correctly formated

	//handle error better
	if err != nil {
		panic(err)
	}

	// call create profile
	h.service.Create(profile)
	// return and log response
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Profile Created"))
}
