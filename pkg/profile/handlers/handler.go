package profile

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CBSktravers/hooli/pkg/profile/models"
)

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	// Log users request
	// check permissons now or early?
	log.Println("Create Profile called by user:")

	decoder := json.NewDecoder(r.Body)
	var profile models.Profile
	err := decoder.Decode(&profile)

	// Varify all input is there and correctly formated

	//handle error better
	if err != nil {
		panic(err)
	}

	// call create profile
	// return and log response
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Profile Created"))
}
