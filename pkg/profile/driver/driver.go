package driver

import (
	"context"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/CBSktravers/hooli/pkg/profile"
	"github.com/CBSktravers/hooli/pkg/profile/handlers"
	profileRepo "github.com/CBSktravers/hooli/pkg/profile/repository"
)

// Configuration struct
type Configuration struct {
	Server, MongoDBHost, DBUser, DBPwd, Database string
}

//StartWebServer starts the profile server
func StartWebServer(port string) {
	// Set up app config server default values
	AppConfig := DefaultConfiguration()
	// if envrionment values exist overwirite defaule values
	loadConfigFromEnvironment(&AppConfig)
	uri := AppConfig.MongoDBHost + "://" + AppConfig.Server
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

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

	collection := client.Database(AppConfig.Database).Collection(AppConfig.Database)
	// Get a handle for your collection
	profileSvc := profile.NewDefaultService(profileRepo.NewMongo(collection))

	logger := log.New(os.Stdout, "hooli ", log.LstdFlags|log.Lshortfile)
	//Set up http handlers
	logger.Println("Starting metadata HTTP service at " + port)

	h := handlers.NewHandlers(logger, profileSvc)

	mux := http.NewServeMux()

	h.SetupRoutes(mux)

	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}

// DefaultConfiguration create new instance of Something
func DefaultConfiguration() Configuration {
	var AppConfig Configuration
	AppConfig.Server = "localhost:27017"
	AppConfig.MongoDBHost = "mongodb"
	AppConfig.DBUser = ""
	AppConfig.DBPwd = ""
	AppConfig.Database = "profile"
	return AppConfig
}

// Load configuration from environment
func loadConfigFromEnvironment(appConfig *Configuration) {
	server, ok := os.LookupEnv("PROFILE_SERVER")
	if ok {
		appConfig.Server = server
		log.Printf("[INFO]: Server information loaded from env.")
	}

	mongodbHost, ok := os.LookupEnv("PROFILE_MONGODB_HOST")
	if ok {
		appConfig.MongoDBHost = mongodbHost
		log.Printf("[INFO]: MongoDB host information loaded from env.")
	}

	mongodbUser, ok := os.LookupEnv("PROFILE_MONGODB_USER")
	if ok {
		appConfig.DBUser = mongodbUser
		log.Printf("[INFO]: MongoDB user information loaded from env.")
	}

	mongodbPwd, ok := os.LookupEnv("PROFILE_MONGODB_PWD")
	if ok {
		appConfig.DBPwd = mongodbPwd
		log.Printf("[INFO]: MongoDB password information loaded from env.")
	}

	database, ok := os.LookupEnv("PROFILE_MONGODB_DATABASE")
	if ok {
		appConfig.Database = database
		log.Printf("[INFO]: MongoDB database information loaded from env.")
	}
}
