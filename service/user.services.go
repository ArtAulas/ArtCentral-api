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
			return nil, fmt.Errorf("Erro ao buscar usuários: %e", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Erro ao buscar usuários: %e", err)
	}

	return users, nil
}

func AddUser(user m.Users) (int64, error){
	db := utils.DB
	result, err := db.Exec("INSERT INTO users(name, email) VALUES(?,?)",user.Name, user.Email)
	if err != nil {
		return 0, fmt.Errorf("Erro ao inserir usuário: %e", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Erro ao inserir usuário: %e", err)
	}
	return id, nil
}