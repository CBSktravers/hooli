package repository

import (
	"context"
	"log"

	"github.com/CBSktravers/hooli/pkg/profile/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// Mongo is an implementation of the profile repository backed by a MongoDB
type Mongo struct {
	db *mongo.Collection
}

// NewMongo creates and returns a pointer to Mongo
func NewMongo(db *mongo.Collection) *Mongo {
	return &Mongo{db: db}
}

// Create add a profile to mongodb db
func (m *Mongo) Create(profile *models.Profile) error {
	// Insert a single document into db
	insertResult, err := m.db.InsertOne(context.TODO(), profile)
	// Check if error occured adding to database
	if err != nil {
		log.Println(err)
	}
	log.Println("Inserted a single document: ", insertResult.InsertedID)
	//Look into connection pooling/closing connections
	return err
}
