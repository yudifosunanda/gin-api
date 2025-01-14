package controllers

import (
	"gin-api/models"
	"gin-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context){
	users :=  services.GetAllUsers()
	context.JSON(http.StatusOK, gin.H({
		"data" :users
	}))
}