package dbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/CBSktravers/hooli/profile/internal/app/model"
	"github.com/boltdb/bolt"
)

type IBoltClient interface {
	OpenBoltDb()
	QueryProfile(profileId string) (model.Profile, error)
	Seed()
}

// Real implementation
type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("profiles.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Start seeding profiles
func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedProfiles()
}

// Creates an "ProfileBucket" in our BoltDB. It will overwrite any existing bucket of the same name.
func (bc *BoltClient) initializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("ProfileBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

// Seed (n) make-believe profile objects into the AcountBucket bucket.
func (bc *BoltClient) seedProfiles() {

	total := 100
	for i := 0; i < total; i++ {

		// Generate a key 10000 or larger
		key := strconv.Itoa(10000 + i)

		// Create an instance of our Profile struct
		acc := model.Profile{
			Id:   key,
			Name: "Person_" + strconv.Itoa(i),
		}

		// Serialize the struct to JSON
		jsonBytes, _ := json.Marshal(acc)

		// Write the data to the ProfileBucket
		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("ProfileBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Seeded %v fake profiless...\n", total)
}

func (bc *BoltClient) QueryProfile(profileId string) (model.Profile, error) {
	// Allocate an empty Profile instance we'll let json.Unmarhal populate for us in a bit.
	profile := model.Profile{}

	// Read an object from the bucket using boltDB.View
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Read the bucket from the DB
		b := tx.Bucket([]byte("ProfileBucket"))

		// Read the value identified by our profileId supplied as []byte
		profileBytes := b.Get([]byte(profileId))
		if profileBytes == nil {
			return fmt.Errorf("No profile found for " + profileId)
		}
		// Unmarshal the returned bytes into the profile struct we created at
		// the top of the function
		json.Unmarshal(profileBytes, &profile)

		// Return nil to indicate nothing went wrong, e.g no error
		return nil
	})
	// If there were an error, return the error
	if err != nil {
		return model.Profile{}, err
	}
	// Return the Profile struct and nil as error.
	return profile, nil
}
