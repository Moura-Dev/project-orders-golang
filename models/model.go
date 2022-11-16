package models

import "time"

type Address struct {
	Id          int32     `db:"id"`
	UserId      int32     `db:"user_id" json:"user_id"`
	AddressName string    `db:"address_name" json:"address_name"`
	Street      string    `db:"street" json:"street"`
	Number      string    `db:"number" json:"number"`
	ZipCode     string    `db:"zip_code" json:"zip_code"`
	City        string    `db:"city" json:"city"`
	State       string    `db:"state" json:"state"`
	District    string    `db:"district" json:"district"`
	Reference   string    `db:"reference" json:"reference"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type Phones struct {
	Id             int32     `db:"id" json:"id"`
	FirstPhone     string    `db:"first_phone" json:"first_phone"`
	SecondaryPhone string    `db:"secondary_phone" json:"secondary_phone"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}
