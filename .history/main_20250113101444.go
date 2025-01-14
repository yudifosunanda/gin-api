package main

import "github.com/gin-gonic/gin"
import "net/http"
import "gin-api/routes"

func main(){
// Initialize the database connection
	db.InitDB()

	// Create a new Gin router
	routes := gin.Default()

	// Set up the routes
	routes.SetupRoutes(routes)

	// Start the server
	r.Run(":8080")
}