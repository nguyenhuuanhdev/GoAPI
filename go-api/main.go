package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: "1", Name: "Anh", Age: 18},
	{ID: "2", Name: "Vu", Age: 20},
}

func createUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	users = append(users, newUser)
	c.JSON(http.StatusCreated, gin.H{
		"message": "đã tạo người dùng",
		"user":    newUser,
	})

}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser User

	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	for i, u := range users {
		if u.ID == id {
			users[i].Name = updatedUser.Name
			users[i].Age = updatedUser.Age

			c.JSON(http.StatusOK, gin.H{
				"message": "đã cập nhật người dùng",
				"user":    users[i],
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "cập nhật thông tin thất bại",
	})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "xóa user thành công",
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "xóa user thất bại",
	})
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Go API",
	})
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, testID := range users {
		if testID.ID == id {
			c.JSON(http.StatusOK, testID)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Không thấy id",
	})
}

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/", getHome)
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUserByID)
	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	r.Run(":8080")
}
