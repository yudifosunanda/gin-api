package routes

import (
	"github.com/gin-gonic/gin"
	"gin-api/controllers"
)

// SetupRoutes sets up the application's routes
func SetupRoutes(r *gin.Engine) {
	userController := &controllers.UserController{}

	r.GET("/users", usersController.GetUsers)
	r.POST("/users", usersController.CreateUser)
}
