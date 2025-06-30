package routes

import (
	"gin-api/auth"
	"gin-api/controllers"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func SetupRoutes(router *gin.Engine) {

	// get csrf token
	router.GET("/csrf-token", func(c *gin.Context) {
		token := csrf.GetToken(c)
		c.JSON(200, gin.H{"csrf_token": token})
	})

	// jwt autentication
	// Login route (returns token)
	router.POST("/login", auth.AuthMiddleware.LoginHandler)

	// JWT refresh route
	router.GET("/refresh", auth.AuthMiddleware.RefreshHandler)

	// Protected group
	authGroup := router.Group("/api")

	authGroup.Use(auth.AuthMiddleware.MiddlewareFunc())
	{
		// Define route for user actions
		authGroup.GET("/users", controllers.GetUsers)
		authGroup.GET("/users/:userId", controllers.GetUsersById)
		authGroup.POST("/users/add", controllers.CreateUser)
		authGroup.PUT("/users/:userId", controllers.UpdateUser)
		authGroup.DELETE("/users/:userId", controllers.DeleteUser)
		authGroup.PUT("/update-password/:userId", controllers.UpdatePassword)
	}
}
