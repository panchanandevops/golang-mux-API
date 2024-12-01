package main

import (
	"log"
	"net/http"

	"go-crud-app/db"
	"go-crud-app/router"

	_ "github.com/lib/pq"
)

func main() {
	// Initialize database
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Setup routes
	r := router.SetupRouter(database)

	// Start the server
	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
