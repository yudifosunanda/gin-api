package controllers

import (
	"gin-api/models"
	"gin-api/db"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context){
	var users []userModels.User
	if err := db.DB.Find(&users).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	context.JSON(http.StatusOK, users)
}

func CreateUser(context *gin.context){
	var users userModels.User

	if err := context.ShouldBinJSON(&users); err != nil{
		context.JSON(http.StatusBadRequest, gin.H{
			"code" : 500,
			"status" : "failed",
			"message" : "failed input"
		})
	}
}