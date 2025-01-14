package main

import "github.com/gin-gonic/gin"
import "net/http"
import "gin-api/routes"

func main(){
// Initialize the database connection
db.InitDB()

// Create a new Gin router
r := gin.Default()

// Set up the routes
routes.SetupRoutes(r)

// Start the server
r.Run(":8080")
}