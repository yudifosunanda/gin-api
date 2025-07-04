package controllers

import (
	"gin-api/db"
	"gin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(context *gin.Context) {
	var users []models.User
	if err := db.DB.Preload("Roles").Find(&users).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	context.JSON(http.StatusOK, users)
}

func GetUsersById(context *gin.Context) {
	userId := context.Param("userId")

	var user models.User
	if err := db.DB.Preload("Roles").First(&user, userId).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "User Not Found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": user,
	})
}

func CreateUser(context *gin.Context) {
	var users models.User

	if err := context.ShouldBindJSON(&users); err != nil {
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
			"code":    500,
			"status":  "failed",
			"message": "failed create user",
		})
		return
	}

	context.JSON(http.StatusCreated, users)
}

func UpdateUser(context *gin.Context) {
	var users models.User
	userId := context.Param("userId")

	if err := db.DB.First(&users, userId).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "User Not Found"})
		return
	}

	var updatedFields struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := context.ShouldBindJSON(&updatedFields); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"status":  "failed",
			"message": "Invalid input data",
			"errors":  err.Error(),
		})
		return
	}

	// Update only the specified fields
	if err := db.DB.Model(&users).Updates(updatedFields).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": users,
	})

}

func UpdatePassword(context *gin.Context) {
	var users models.User

	userId := context.Param("userId")

	if err := db.DB.First(&users, userId).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "user not found",
			"status":  "failed",
		})
	}

	var updateData struct {
		Password string `json:"password"`
	}

	if err := context.ShouldBindJSON(&updateData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"status":  "failed",
			"message": "Invalid input data",
			"errors":  err.Error(),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateData.Password), bcrypt.DefaultCost)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":   500,
			"status": "failed to hash password",
			"error":  err.Error(),
		})
		return
	}

	if err := db.DB.Model(&users).Update("password", hashedPassword).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"status":  "failed",
			"message": "failed to update password",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"status":  "success",
		"message": "success update password",
	})

}

func DeleteUser(context *gin.Context) {
	var user models.User
	userId := context.Param("userId")

	// find data
	if err := db.DB.First(&user, userId).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": "No Data Found",
		})
		return
	}

	// delete data
	if err := db.DB.Delete(&user, userId).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": "Error on delete",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success delete data",
	})
}
