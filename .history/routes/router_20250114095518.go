package routes

import (
	"github.com/gin-gonic/gin"
	"gin-api/controllers"
	"net/http"
)

func SetupRoutes(router *gin.Engine) {
	// test connection
	// router.GET("/ping", func(context *gin.Context){
	// 	context.JSON(http.StatusOK, gin.H{
	// 		"code" : 200,
	// 		"status" : "success",
	// 		"message" : "success connected",
	// 	})
	// })

	// Define route for user actions
  userRoutes := router.Group("/users")
	{
    userRoutes.GET("/", controllers.GetUsers)
    userRoutes.GET("/:userId", controllers.GetUsersById)
    userRoutes.POST("/add", controllers.CreateUser)
    userRoutes.PUT("/:userId", controllers.UpdateUser)
	}
}
