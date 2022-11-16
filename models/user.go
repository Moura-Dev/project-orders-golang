package models

import "time"

type User struct {
	Id        int32     `db:"id" json:"ID"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	CPF       string    `json:"cpf" json:"cpf"`
	Phones    Phones    `db:"phones" json:"phones"`
	Email     string    `db:"email" json:"email"`
	Login     string    `db:"login" json:"login"`
	Password  string    `db:"password" json:"password"`
	Active    bool      `db:"active" json:"active"`
	Addresses []Address `db:"addresses" json:"addresses"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
