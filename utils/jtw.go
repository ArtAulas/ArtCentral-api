package utils

import (
 "fmt"
 "github.com/golang-jwt/jwt/v5"
 "time"
 "os"

 m "example/artcentral-api/models"
)

var secret_key = []byte(os.Getenv("JWT_SECRET"))

func CreateToken(user m.Users) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":user.ID,
			"name":user.Name,
			"email":user.Email,
			"role":user.Role,
			"exp":time.Now().Add(time.Hour*24).Unix(),
		})
	tokenString, err := token.SignedString(secret_key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims,error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		return secret_key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Invalid Token")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64){
		return nil, fmt.Errorf("ExpiredToken")
	}
	return claims, err
}