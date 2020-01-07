package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here is the Profile"))
}

// Handler for HTTP Get - "/profiles"
// Returns all profiles documents
func GetProfiles(w http.ResponseWriter, r *http.Request) {
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
	addProfile()
	fmt.Println(r)
	json.NewDecoder(r.Body)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Profile Created"))
}
