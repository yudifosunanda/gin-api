package main

import "github.com/gin-gonic/gin"
import "gin-api/routes"
import "gin-api/db"
import "github.com/utrack/gin-csrf"
import "github.com/gin-contrib/sessions"
import "github.com/gin-contrib/sessions/cookie"

func main(){
	// Initialize the database connection
	db.InitDB()

	// Create a new Gin router
	r := gin.Default()

	// Initialize session store (required by csrf)
	store := cookie.NewStore([]byte("super-secret-session-key"))
	r.Use(sessions.Sessions("mysession", store))

	// Apply CSRF middleware
	r.Use(csrf.Middleware(csrf.Options{
		Secret: "a-very-secret-32-byte-key-here!!", // should be 32 bytes
		ErrorFunc: func(c *gin.Context) {
			c.JSON(403, gin.H{"error": "CSRF token mismatch"})
			c.Abort()
		},
	}))

	// Setup routes
	routes.SetupRoutes(r)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}