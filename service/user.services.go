package service

import (
	m "example/artcentral-api/models"
	"example/artcentral-api/utils"
	"log"
	"fmt"
)

func FetchAllUsers() ([]m.Users, error){
	db := utils.DB
	var users []m.Users

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next(){
		var user m.Users
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, fmt.Errorf("Erro ao buscar clientes: %e", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Erro ao buscar clientes: %e", err)
	}

	return users, nil
}