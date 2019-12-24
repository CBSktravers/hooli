package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CBSktravers/hooli/profile/dbclient"
	"github.com/gorilla/mux"
)

var DBClient dbclient.IBoltClient

func GetProfile(w http.ResponseWriter, r *http.Request) {

	// Read the 'profileId' path parameter from the mux map
	var profileId = mux.Vars(r)["profileId"]

	// Read the profile struct BoltDB
	profile, err := DBClient.QueryProfile(profileId)

	// If err, return a 404
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(profile)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
