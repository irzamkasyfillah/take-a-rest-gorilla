package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// initializes the database connection.
func Connect() {
	// Get environment variables
	dbHost, dbPort, dbUser, dbPass, dbName := GetConfig()

	var err error
	// Connect to the database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Connected to database")
}

// GetDB returns the database connection.
func GetDB() *gorm.DB {
	return db
}
