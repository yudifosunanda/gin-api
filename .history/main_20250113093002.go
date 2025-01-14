package main

import "fmt"
import "github.com/gin-gonic/gin"
import "net/http"

func main(){
	route := gin.default()
	route.GET("/ping"), func(context, *gin.Context){
		context.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"status" : "success",
			"message" : "success connected"
		})
	}
}