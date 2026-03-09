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
    accountID, _ := c.Get("id")
    var users []models.User
    database.DB.Where("account_id = ?", accountID).Find(&users)
    c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
    accountID, _ := c.Get("id")
    id := c.Param("id")
    var user models.User

    
    if err := database.DB.Where("id = ? AND account_id = ?", id, accountID).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "không tìm thấy hoặc không có quyền"})
        return
    }

    c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
    accountID, _ := c.Get("id")
    var newUser models.User
    if err := c.BindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newUser.AccountID = uint(accountID.(float64))
    database.DB.Create(&newUser)
    c.JSON(http.StatusCreated, gin.H{"message": "đã tạo user", "user": newUser})
}

func UpdateUser(c *gin.Context) {
    accountID, _ := c.Get("id")
    id := c.Param("id")
    var user models.User

    if err := database.DB.Where("id = ? AND account_id = ?", id, accountID).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "không tìm thấy hoặc không có quyền"})
        return
    }

    var input models.User
    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user.Name = input.Name
    user.Age = input.Age
    database.DB.Save(&user)
    c.JSON(http.StatusOK, gin.H{"message": "đã cập nhật", "user": user})
}

func DeleteUser(c *gin.Context) {
    accountID, _ := c.Get("id")
    id := c.Param("id")
    var user models.User

    if err := database.DB.Where("id = ? AND account_id = ?", id, accountID).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "không tìm thấy hoặc không có quyền"})
        return
    }

    database.DB.Delete(&user)
    c.JSON(http.StatusOK, gin.H{"message": "xóa thành công"})
}