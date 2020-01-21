package driver

import (
	"log"
	"os"
)

// Configuration struct
type Configuration struct {
	Server, MongoDBHost, DBUser, DBPwd, Database string
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

// ConfigProfile profile database connection
func createConfig() Configuration {
	AppConfig := DefaultConfiguration()
	loadConfigFromEnvironment(&AppConfig)
	return AppConfig
}
