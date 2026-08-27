package service

import (
	"database/sql"
	m "example/artcentral-api/models"
	utils "example/artcentral-api/utils"
	"fmt"
	"log"
)

func FetchAllUsers() ([]m.Users, error){
	db := utils.DB
	var users []m.Users

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var user m.Users
		if err := rows.Scan(
				&user.ID, 
				&user.Name, 
				&user.Email,
				&user.Password,
				&user.BirthDate,
				&user.Role,
				&user.CreatedAt,
				&user.UpdatedAt,
			); err != nil {
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

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return 0, fmt.Errorf("Erro ao criptografar senha: %e", err)
	}

	result, err := db.Exec(
		"INSERT INTO users(name, email, password, birthDate, role) VALUES(?,?,?,?,?)",
		user.Name, 
		user.Email,
		password,
		user.BirthDate,
		user.Role,
	)
	if err != nil {
		return 0, fmt.Errorf("Erro ao inserir usuário: %e", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Erro ao inserir usuário: %e", err)
	}
	return id, nil
}

func FetchUserByEmail(email string) (m.Users, error){
	db := utils.DB
	var user m.Users

	row := db.QueryRow("SELECT * FROM users WHERE email = ?", email)
	if err := row.Scan(			
			&user.ID, 
			&user.Name, 
			&user.Email,
			&user.Password,
			&user.BirthDate,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			if err == sql.ErrNoRows{
				return user, fmt.Errorf("Sem usuário com email: %v", email)
			}
			return user, fmt.Errorf("Email: %v, Erro: %v", email, err)
		}
	return user, nil
}