package main

import "github.com/gin-gonic/gin"
import "net/http"
import "gin-api/routes"

func main(){
	// routes
	route := gin.Default()
	//database connections
	database.ConnectDatabase()

	routes.SetupRoutes(router)

	err := route.Run(":8080")
	if err != nil {
		panic(err)
	}
}