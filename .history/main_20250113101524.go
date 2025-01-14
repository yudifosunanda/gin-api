package main

import "github.com/gin-gonic/gin"
import "net/http"
import "gin-api/routes"
import "gin-api/db"

func main(){
// Initialize the database connection
	db.InitDB()

	// Create a new Gin router
	route := gin.Default()

	// Set up the routes
	routes.SetupRoutes(route)

	// Start the server
	route.Run(":8080")
}