package main

import "github.com/gin-gonic/gin"
import "gin-api/routes"
import "gin-api/db"

func main(){
	// Initialize the database connection
	db.InitDB()

	// Create a new Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)
	for _, r := range router.Routes() {
    fmt.Println(r.Method, r.Path)
}
	// Start the server
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}