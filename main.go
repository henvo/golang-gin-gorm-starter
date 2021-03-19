package main

import (
	"github.com/henvo/golang-gin-gorm-starter/models"
	"github.com/henvo/golang-gin-gorm-starter/routes"
)

// Entrypoint for app.
func main() {
	// Load the routes
	r := routes.SetupRouter()

	// Initialize database
	models.SetupDatabase()

	// Start the HTTP API
	r.Run()
}
