package controllers

import (
	"gin-api/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"db"
)

func GetUsers(context *gin.Context){
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, users)
}