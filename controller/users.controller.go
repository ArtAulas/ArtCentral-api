package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	service "example/artcentral-api/service"
)

func GetAllUsers(c *gin.Context){
	users, err := service.FetchAllUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"erro":err})
	}
	c.IndentedJSON(http.StatusOK, users)
}