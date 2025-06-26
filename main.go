package main

import "github.com/gin-gonic/gin"
import "gin-api/routes"
import "gin-api/db"
import "github.com/utrack/gin-csrf"
import "github.com/gin-contrib/sessions"
import "github.com/gin-contrib/sessions/cookie"
import "os"
import "fmt"
import "github.com/joho/godotenv"
import "gin-api/auth"

func main(){
	// Initialize the database connection
	db.InitDB()

	
	// jwt
	auth.InitJWT()
	
	// Create a new Gin router
	r := gin.Default()

	err := godotenv.Load()

	if err != nil{
		fmt.Println("no env found")
	}

	fmt.Println(os.Getenv("COOKIE_SECRET_KEY"))

	// Initialize session store (required by csrf)
	store := cookie.NewStore([]byte(os.Getenv("COOKIE_SECRET_KEY")))
	r.Use(sessions.Sessions("mysession", store))

	// Apply CSRF middleware
	r.Use(csrf.Middleware(csrf.Options{
		Secret: os.Getenv("CSRF_SECRET_KEY"), // should be 32 bytes
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