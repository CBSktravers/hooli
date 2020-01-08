package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CBSktravers/hooli/profile/models"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here is the Profile"))
}

// Handler for HTTP Get - "/profiles"
// Returns all profiles documents
func GetProfiles(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Profiles called by user:")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here are the requested profiles"))
}
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Profile Updated"))
}
func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Profile Deleted"))
}
func CreateProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Profile called by user:")
	decoder := json.NewDecoder(r.Body)
	var profile models.Profile
	err := decoder.Decode(&profile)

	if err != nil {
		panic(err)
	}

	log.Println("User input:", profile.Name)
	//AddProfileMong()
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Profile Created"))
}
