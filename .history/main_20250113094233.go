package main

import "github.com/gin-gonic/gin"
import "net/http"

func main(){
	route := gin.Default()
	route.GET("/ping", func(context *gin.Context){
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"status":  "success",
			"message": "success connected",
		})
	})

	err := route.Run(":8080")
	if err != nil {
		panic(err)
	}
}


// package main

// import (
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// )

// func main() {
// 	route := gin.Default()
// 	route.GET("/ping", func(context *gin.Context) {
// 		context.JSON(http.StatusOK, gin.H{
// 			"code":    200,
// 			"status":  "success",
// 			"message": "success connected",
// 		})
// 	})

// 	err := route.Run(":8080")
// 	if err != nil {
// 		panic(err)
// 	}
// }