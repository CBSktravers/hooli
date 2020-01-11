package repository

import (
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
	// use buffalo validator to validate input
	log.Println("Mongo Create called")
	return nil
}
