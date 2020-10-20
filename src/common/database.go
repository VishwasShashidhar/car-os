package common

import (
	"fmt"
	"github.com/VishwasShashidhar/car-os/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strconv"
)

var db *gorm.DB

// getDatabaseUrl ... Constructs the database url
func getDatabaseUrl() string {
	port, err := strconv.Atoi(GetConfigurationForKey("DB_PORT"))
	if err != nil {
		panic("port is not a number. panicking...")
	}

	return fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%d host=%s",
		GetConfigurationForKey("DB_USER"),
		GetConfigurationForKey("DB_PASSWORD"),
		GetConfigurationForKey("DB_DATABASE"),
		port,
		GetConfigurationForKey("DB_HOST"),
	)
}

func ConnectDatabase() {
	var err error
	db, err = gorm.Open(postgres.Open(getDatabaseUrl()), &gorm.Config{})
	if err != nil {
		log.Fatal("failed connection to database")
	}
	db.AutoMigrate(&models.User{})
}
