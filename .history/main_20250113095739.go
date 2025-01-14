package main

import "github.com/gin-gonic/gin"
import "net/http"

func main(){
	// routes
	route := gin.Default()
	//database connections
	database.ConnectDatabase()


	err := route.Run(":8080")
	if err != nil {
		panic(err)
	}
}