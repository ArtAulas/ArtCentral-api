package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	service "example/artcentral-api/service"
	m "example/artcentral-api/models"
)

func GetAllUsers(c *gin.Context){
	users, err := service.FetchAllUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"erro":err})
	}
	c.IndentedJSON(http.StatusOK, users)
}

func AddUser(c *gin.Context){
	var newUser m.Users

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"erro":"Body Inválido"})
	}

	id, err := service.AddUser(newUser)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"erro":err})
	}
	newUser.ID = int(id)
	c.IndentedJSON(http.StatusOK, gin.H{"Usuário Cadastrado":newUser})
}