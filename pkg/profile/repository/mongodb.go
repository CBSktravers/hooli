package repository

import (
	"context"
	"log"

	"github.com/CBSktravers/hooli/pkg/profile/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// Mongo is an implementation of the profile repository backed by a Mongo DB
type Mongo struct {
	db *mongo.Collection
}

// NewMongo creates and returns a pointer to an Mongo
func NewMongo(db *mongo.Collection) *Mongo {
	return &Mongo{db: db}
}

// Create add a profile to mongodb repo
func (r *Mongo) Create(profile *models.Profile) error {
	// Insert a single document
	//insertResult, err := collection.InsertOne(context.TODO(), profile)
	insertResult, err := r.db.InsertOne(context.TODO(), profile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted a single document: ", insertResult.InsertedID)

	//Close client
	//err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB closed.")

	return nil
}
