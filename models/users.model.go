package models

type Users struct {
	ID int
	Name string
	Email string
	Password string
	BirthDate string
	Role string
	CreatedAt string
	UpdatedAt string
}

type Login struct {
	Email string
	Password string
}