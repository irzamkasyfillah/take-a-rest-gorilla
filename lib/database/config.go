package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// func init() {
// 	// Load environment variables
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Println("Error loading .env file")
// 	}
// }

func GetConfig() (host string, port string, user string, pass string, name string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	// Set environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_DATABASE")

	return dbHost, dbPort, dbUser, dbPass, dbName
}
