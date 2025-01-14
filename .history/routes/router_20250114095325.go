package routes

import (
	"github.com/gin-gonic/gin"
	"gin-api/controllers"
	"net/http"
)

func SetupRoutes(router *gin.Engine) {
	// test connection
	router.GET("/ping", func(context *gin.Context){
		context.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"status" : "success",
			"message" : "success connected",
		})
	})

	// Define route for user actions
  router.Group("/users")
	{
    router.GET("/", controllers.GetUsers)
    router.GET("/:userId", controllers.GetUsersById)
    router.POST("/", controllers.CreateUser)
    router.PUT("/:userId", controllers.UpdateUser)
	}
}
