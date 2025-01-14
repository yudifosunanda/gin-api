package routes

import (
	"github.com/gin-gonic/gin"
	"gin-api/controllers"
	"net/http"
)

func SetupRoutes(router *gin.Engine) {
	// test connection
	router.GET("/pingx", func(context *gin.Context){
		context.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"status" : "success",
			"message" : "success connected",
		})
	})

	// Define route for user actions
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:userId", controllers.CreateUser)
	router.POST("/users", controllers.CreateUser)
}
