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

func GetUsersById(context *gin.Context){
	userId := context.Param("userId") 
	context.JSON(userId)

	// var user userModels.User
	// if err := db.DB.First(&userId).Error; err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": "User Not Found"})
	// 	return
	// }
	// context.JSON(http.StatusOK, user)
}

func CreateUser(context *gin.Context){
	var users userModels.User

	if err := context.ShouldBindJSON(&users); err != nil{
		 context.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"status":  "failed",
				"message": "Invalid input data",
				"errors":  err.Error(),
		})
		return
	}

	if err := db.DB.Create(&users).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code" : 500,
			"status" : "failed",
			"message" : "failed create user",
		})
		return
	}
	
	context.JSON(http.StatusCreated, users)
}