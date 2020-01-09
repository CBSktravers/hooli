package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CBSktravers/hooli/profile/models"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Profile called by user:")
	decoder := json.NewDecoder(r.Body)
	var profile models.Profile
	err := decoder.Decode(&profile)

	//handle error better
	if err != nil {
		panic(err)
	}

	result := getMongoProfile(profile)
	j, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(j))
}

// Handler for HTTP Get - "/profiles"
// Returns all profiles documents
func GetProfiles(w http.ResponseWriter, r *http.Request) {
	//Only want one param might no want to pass a profile
	log.Println("Get Profiles called by user:")
	decoder := json.NewDecoder(r.Body)
	var profile models.Profile
	err := decoder.Decode(&profile)

	//handle error better
	if err != nil {
		panic(err)
	}

	results := getMongoProfilesByDepartment(profile)
	j, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(j))
}
func GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	log.Println("Get all Profiles called by user:")
	decoder := json.NewDecoder(r.Body)
	var profile models.Profile
	err := decoder.Decode(&profile)

	//handle error better
	if err != nil {
		panic(err)
	}

	results := getAllMongoProfiles()
	j, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(j))
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

	//handle error better
	if err != nil {
		panic(err)
	}

	createMongoProfile(profile)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Profile Created"))
}
