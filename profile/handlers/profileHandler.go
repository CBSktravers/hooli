package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here is the Profile"))
}

//DELETE HERE
const (
	hosts      = "localhost:27017"
	database   = "db"
	username   = ""
	password   = ""
	collection = "profile"
)

type Profile struct {
	Name       string `json:"name"`
	Department string `json:"department"`
}
type MongoStore struct {
	session *mgo.Session
}

var mongoStore = MongoStore{}

func initialiseMongo() (session *mgo.Session) {

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	return

}

// Handler for HTTP Get - "/profiles"
// Returns all profiles documents
func GetProfiles(w http.ResponseWriter, r *http.Request) {
	//Create MongoDB session
	session := initialiseMongo()
	mongoStore.session = session

	w.Header().Set("Access-Control-Allow-Origin", "*")

	col := mongoStore.session.DB(database).C(collection)

	results := []Profile{}
	col.Find(bson.M{"name": bson.RegEx{"", ""}}).All(&results)
	jsonString, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(jsonString))

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
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Profile Created"))
}
