package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// SayHello ... Says hello to the user
func SayHello() string {
	return "Hello, cars!"
}

// GetAppName ... Gets the app name from the configuration
func GetAppName() string {
	appName, present := os.LookupEnv("APP_NAME")
	if !present {
		log.Println("app name not present in the configuration, resorting to default value!")
		appName = "CarOS"
	}
	return appName
}

func loadConfiguration() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed loading environment configuration: %s", err.Error())
	}
}

func main() {
	loadConfiguration()
	fmt.Printf(GetAppName())
}
