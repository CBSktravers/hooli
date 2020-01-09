package handlers

//handle errors return status of work
import (
	"context"
	"log"

	"github.com/CBSktravers/hooli/profile/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createClient() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	return client
}

func createMongoProfile(profile models.Profile) {
	// Establish client to mongodabase
	client := createClient()

	// Get a handle for your collection
	collection := client.Database("mongodb").Collection("profile")

	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), profile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted a single document: ", insertResult.InsertedID)

	//Close client
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB closed.")
}

func getMongoProfile(profile models.Profile) models.Profile {
	// Establish client to mongodabase
	client := createClient()

	// Get a handle for your collection
	collection := client.Database("mongodb").Collection("profile")

	var result models.Profile

	// Create filter to find item in database
	//filter := bson.D{{"name", profile.Name}}
	var filter bson.D
	filter = append(filter, bson.E{"department", profile.Department})
	filter = append(filter, bson.E{"name", profile.Name})
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found a single document: %+v\n", result)

	return result
}

func getAllMongoProfiles() []*models.Profile {
	// Establish client to mongodabase
	client := createClient()

	// Get a handle for your collection
	collection := client.Database("mongodb").Collection("profile")

	findOptions := options.Find()
	//findOptions.SetLimit(2)

	// Finding multiple documents returns a cursor
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	var results []*models.Profile

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem models.Profile
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	log.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results
}
func getMongoProfilesByDepartment(profile models.Profile) []*models.Profile {
	// Establish client to mongodabase
	client := createClient()

	// Get a handle for your collection
	collection := client.Database("mongodb").Collection("profile")

	findOptions := options.Find()
	filter := bson.D{{"department", profile.Department}}

	// Finding multiple documents returns a cursor
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	var results []*models.Profile

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem models.Profile
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	// based on return change logs
	log.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results
}

func deleteMongoProfile(profile models.Profile) {
	// Establish client to mongodabase
	client := createClient()

	// Get a handle for your collection
	collection := client.Database("mongodb").Collection("profile")

	var filter bson.D
	filter = append(filter, bson.E{"department", profile.Department})
	filter = append(filter, bson.E{"name", profile.Name})

	// Delete all the documents in the collection
	deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	// Close the connection once no longer needed
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connection to MongoDB closed.")
	}
}

//updateProfile
//addProfiles
//deleteProfile
//deleteProfiles
//getProfiles
