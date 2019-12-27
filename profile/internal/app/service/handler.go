package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	name       string             `json:"name,omitempty" bson:"name,omitempty"`
	department string             `json:"department,omitempty" bson:"department,omitempty"`
}

func CreateProfile(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var profile Profile
	_ = json.NewDecoder(request.Body).Decode(&profile)
	collection := client.Database("hooli").Collection("profiles")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, profile)
	json.NewEncoder(response).Encode(result)
}
func GetProfile(response http.ResponseWriter, request *http.Request) {

}
func GetProfiles(response http.ResponseWriter, request *http.Request) {

}
