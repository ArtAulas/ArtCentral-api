package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	c "example/artcentral-api/controller"
	"example/artcentral-api/middleware"
	"example/artcentral-api/utils"

	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main(){
	utils.SetDB()

	router := gin.Default()

	router.GET("/", middleware.AuthMiddleware(), getHome)

	users := router.Group("/Users")
	users.GET("/", c.GetAllUsers)
	users.POST("/", c.AddUser)
	users.POST("/login", c.Login)

	port := os.Getenv("API_PORT")
	router.Run("localhost:"+port)
}

func getHome(c *gin.Context){
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello! This is my API!"})
	user, _ := c.Get("user")
	c.IndentedJSON(http.StatusOK, gin.H{"user": user})
}