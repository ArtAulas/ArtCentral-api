package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	service "example/artcentral-api/service"
)

func GetAllUsers(c *gin.Context){
	users := service.FetchAllUsers()
	c.IndentedJSON(http.StatusOK, users)
}