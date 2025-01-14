package routes

import (
	"github.com/gin-gonic/gin"
	"gin-api/controllers"
)

func SetupRoutes(router *gin.Engine) {
	// Define route for user actions
	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)
}
