package models

import "time"

type Company struct {
	ID          int32     `db:"id"`
	UserID      int32     `db:"user_id" json:"user_id"`
	Name        string    `db:"name" json:"name"`
	CNPJ        string    `db:"cnpj" json:"cnpj"`
	FantasyName string    `db:"fantasy_name" json:"fantasy_name"`
	IE          string    `db:"ie" json:"ie"`
	Logo        string    `db:"logo" json:"logo"`
	Phones      Phones    `json:"phones"`
	Addresses   []Address `json:"address"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	// TODO: fix this type to receive from body and parse to body
	//CompanyType CompanyType `json:"company_type"`
}

type Costumer struct {
}

type ShippingCompany struct {
}

type Factory struct {
}

type CompanyType struct {
	ShippingCompany ShippingCompany
	CUSTOMER        Costumer
	FACTORY         Factory
}
