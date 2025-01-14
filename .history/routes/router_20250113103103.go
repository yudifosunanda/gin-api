package routes

import (
	"github.com/gin-gonic/gin"
	"gin-api/controllers"
)

// SetupRoutes sets up the application's routes
func SetupRoutes(r *gin.Engine) {
	userController := &controllers.UserController{}

	r.GET("/users", userController.GetUsers)
	r.POST("/users", userController.CreateUser)
}
