package service

import (
	m "example/artcentral-api/models"
)

// MOCK
var users = []m.Users{
	{ID: 1, Name: "John", Email: "john@gmail.com"},
	{ID: 2, Name: "Jane", Email: "jane@gmail.com"},
	{ID: 3, Name: "Jamie", Email: "jamie@gmail.com"},
}

func FetchAllUsers() ([]m.Users){
	return users
}