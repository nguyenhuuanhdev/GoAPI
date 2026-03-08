package controllers

import (
	"go-api/database"
	"go-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "chào mừng đến với API của tôi",
	})
}

func GetUsers(c *gin.Context) {

	var users []models.User

	database.DB.Find(&users)

	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {

	id := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "không tìm thấy người dùng",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {

	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&newUser)

	c.JSON(http.StatusCreated, gin.H{
		"message": "đã tạo người dùng",
		"user":    newUser,
	})
}

func UpdateUser(c *gin.Context) {

	id := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "không tìm thấy người dùng",
		})
		return
	}

	var input models.User

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Name = input.Name
	user.Age = input.Age

	database.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "đã cập nhật người dùng",
		"user":    user,
	})
}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "không tìm thấy người dùng",
		})
		return
	}

	database.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "xóa user thành công",
	})
}