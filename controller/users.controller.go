package controller

import (
	m "example/artcentral-api/models"
	service "example/artcentral-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context){
	users, err := service.FetchAllUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"erro":err})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func AddUser(c *gin.Context){
	var newUser m.Users

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"erro":"Body Inválido"})
		return
	}

	if (newUser.Role != "artist" && newUser.Role != "client") {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"erro":"Role Inválido"})
		return
	}

	id, err := service.AddUser(newUser)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"erro":err})
		return
	}
	newUser.ID = int(id)
	c.IndentedJSON(http.StatusOK, gin.H{"Usuário Cadastrado":newUser})
}