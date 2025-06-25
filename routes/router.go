package routes

import (
	"github.com/gin-gonic/gin"
	"gin-api/controllers"
  "github.com/utrack/gin-csrf"

)

func SetupRoutes(router *gin.Engine) {

	// get csrf token
	router.GET("/csrf-token", func(c *gin.Context) {
		token := csrf.GetToken(c)
		c.JSON(200, gin.H{"csrf_token": token})
	})

	// Define route for user actions	
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:userId", controllers.GetUsersById)
	router.POST("/users/add", controllers.CreateUser)
	router.PUT("/users/:userId", controllers.UpdateUser)
	router.DELETE("/users/:userId", controllers.DeleteUser)
}
