package controller

import (
	"database/sql"
	"errors"
	m "example/artcentral-api/models"
	service "example/artcentral-api/service"
	"example/artcentral-api/utils"

	"net/http"

	"github.com/labstack/echo/v5"
)

func GetAllUsers(c *echo.Context) error {
	users, err := service.FetchAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"erro": err})
	}
	return c.JSON(http.StatusOK, users)
}

func AddUser(c *echo.Context) error {
	var newUser m.Users

	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"erro": "Body Inválido"})
	}

	if newUser.Role != "artist" && newUser.Role != "client" {
		return c.JSON(http.StatusBadRequest, map[string]string{"erro": "Role Inválido"})
	}

	id, err := service.AddUser(newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"erro": err})
	}
	newUser.ID = int(id)
	return c.JSON(http.StatusOK, map[string]any{"Usuário Cadastrado": newUser})
}

func Login(c *echo.Context) error {
	var loginData m.Login

	if err := c.Bind(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"erro": "Body Inválido"})
	}

	user, err := service.FetchUserByEmail(loginData.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"erro": "Email ou senha inválidos"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"erro": "Erro interno do servidor"})
	}

	if !utils.VerifyPassword(loginData.Password, user.Password) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"erro": "Email ou senha inválidos"})
	}

	token, err := utils.CreateToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"erro": "Erro na Criação do Token"})
	}
	return c.JSON(http.StatusOK, map[string]any{"User": user, "Token": token})
}
