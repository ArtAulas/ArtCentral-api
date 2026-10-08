package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	c "example/artcentral-api/controller"
	"example/artcentral-api/utils"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	utils.SetDB()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
	}))
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", getHome)

	users := e.Group("/Users")
	users.GET("/", c.GetAllUsers)
	users.POST("/", c.AddUser)
	users.POST("/login", c.Login)

	port := os.Getenv("API_PORT")
	if err := e.Start("localhost:" + port); err != nil {
		log.Fatal(err)
	}
}

func getHome(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello! This is my API!"})
}
