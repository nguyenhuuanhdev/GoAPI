package routes

import (
	"go-api/controllers"
	"go-api/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
    // Public - không cần token
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    // Protected - cần token
    protected := r.Group("/api")
    protected.Use(middleware.AuthMiddleware())
    {
        protected.GET("/users", controllers.GetUsers)
        protected.POST("/users", controllers.CreateUser)
        protected.PUT("/users/:id", controllers.UpdateUser)
        protected.DELETE("/users/:id", controllers.DeleteUser)
    }
}