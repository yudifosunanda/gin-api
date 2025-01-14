package controllers

import (
	"gin-api/models"
	"gin-api/db"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context){
	var users []moduserModelsels.User
	if err := db.DB.Find(&users).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	context.JSON(http.StatusOK, users)
}