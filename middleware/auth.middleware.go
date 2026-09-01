package middleware

import (
	"example/artcentral-api/utils"
	
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")

		if authHeader == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error":"Authorization header is missing"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return 
		}

		authToken := strings.Split(authHeader, " ")
		if len(authToken) != 2 || authToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid Token Format"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return 
		}

		tokenString := authToken[1]

		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error":err})
			c.AbortWithStatus(http.StatusUnauthorized)
			return 
		}

		c.Set("user",claims)
		c.Next()
	}
}