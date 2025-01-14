package routes

import (
	"github.com/gin-gonic/gin"
	"gin-api/controllers"
)

func SetupRoutes(router *gin.Engine) {
	// test connection
	route.GET("/pingx", func(context *gin.Context){
		context.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"status" : "success",
			"message" : "success connected",
		})
	})

	// Define route for user actions
	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)
}
