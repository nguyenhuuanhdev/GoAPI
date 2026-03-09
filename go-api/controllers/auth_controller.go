package controllers

import (
	"go-api/database"
	"go-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret_key_cua_ban") 

func Register(c *gin.Context) {
    var input models.LoginInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    account := models.Account{
        Username: input.Username,
        Password: input.Password, 
    }

    if err := database.DB.Create(&account).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "username đã tồn tại"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "đăng ký thành công"})
}

func Login(c *gin.Context) {
    var input models.LoginInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var account models.Account
    if err := database.DB.Where("username = ?", input.Username).First(&account).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "sai username hoặc password"})
        return
    }

    // ⚠️ so sánh thẳng vì chưa hash, sau này dùng bcrypt
    if account.Password != input.Password {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "sai username hoặc password"})
        return
    }

    // Tạo JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":       account.ID,
        "username": account.Username,
        "exp":      time.Now().Add(24 * time.Hour).Unix(),
    })

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "không tạo được token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "đăng nhập thành công",
        "token":   tokenString,
    })
}