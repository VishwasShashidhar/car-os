package common

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var defaultValueSet map[string]string

// LoadConfiguration ... Loads the configuration from an environment
func LoadConfiguration() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed loading environment configuration: %s", err.Error())
	}
}

// Init ... Initializes default value set for keys
func Init() {
	defaultValueSet = make(map[string]string)
	defaultValueSet["APP_NAME"] = "CarOS"
	defaultValueSet["APP_PORT"] = "8080"
}

func getDefaultValueForKey(key string) string {
	return defaultValueSet[key]
}

// GetConfigurationForKey ... Gets value for a key.
// Resorts to default value if key not found in the environment.
// If there's no default value available, an empty string is returned
func GetConfigurationForKey(key string) string {
	value, present := os.LookupEnv(key)
	if !present {
		log.Println("key not present in the environment, resorting to default value!")
		return getDefaultValueForKey(key)
	}
	return value
}
