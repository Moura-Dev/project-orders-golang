package models

import "time"

type Costumer struct {
	Id        int32     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Telephone string    `db:"telephone" json:"telephone"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
}
