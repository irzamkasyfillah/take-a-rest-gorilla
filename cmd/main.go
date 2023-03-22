package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	api "github.com/irzam/my-app/api/routes"
	"github.com/irzam/my-app/lib/database"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Connect to the database
	database.Connect()

	// Register API routes
	api.RegisterRoutes(r)

	// Start server
	port := os.Getenv("PORT")
	log.Printf("Starting server on port %s...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
