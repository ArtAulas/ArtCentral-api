package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	c "example/artcentral-api/controller"
)

func main(){
	router := gin.Default()
	router.GET("/" , getHome)

	users := router.Group("/Users")
	users.GET("/", c.GetAllUsers)

	router.Run("localhost:8080")
}

func getHome(c *gin.Context){
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello! This is my API!"})
}