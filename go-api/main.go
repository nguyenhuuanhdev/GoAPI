package main

import (
	"go-api/database"
	"go-api/middleware"
	"go-api/models"
	"go-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.Use(middleware.CORSMiddleware())

    database.ConnectDB()

    database.DB.AutoMigrate(&models.User{}, &models.Account{}) // ✅ thêm Account

    routes.UserRoutes(r)
    routes.AuthRoutes(r) // ✅ thêm dòng này

    r.Run(":8080")
}
