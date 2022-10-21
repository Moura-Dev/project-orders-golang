package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type StatusType struct {
	Quote     string
	Approved  string
	Cancelled string
}

type CompanyType struct {
	PORTAGE  string
	CUSTOMER string
	FACTORY  string
}

type ShippingType struct {
	Cif string
	Fob string
}
type User struct {
	ID        int32  `db:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Email     string
	Login     string
	Password  string `db:"password" json:"password"`
	Active    bool   `db:"active" json:"active"`
}

type Seller struct {
	ID     int32  `db:"id" json:"id"`
	UserID int32  `db:"user_id" json:"user_id"`
	Name   string `db:"name"`
	TaxID  string `db:"tax_id" json:"tax_id"`
}

type Company struct {
	ID          int32  `db:"id"`
	UserID      int32  `db:"user_id" json:"user_id"`
	Name        string `db:"name" json:"name"`
	Cnpj        string `db:"cnpj" json:"cnpj"`
	FantasyName string `db:"fantasy_name" json:"fantasy_name"`
	IE          string `db:"ie" json:"ie"`
	Fone        string `db:"fone" json:"fone"`
	Fone2       string `db:"fone2" json:"fone2"`
	Logo        string `db:"logo" json:"logo"`
	CompanyType string `db:"company_type" json:"company_type"`
	AddressID   int32  `db:"address_id" json:"address_id"`
}

type Address struct {
	ID         int32 `db:"id"`
	Street     string
	Number     string
	ZipCode    string `db:"zip_code"`
	City       string
	State      string
	District   string
	Reference  string
	AdressName string
	UserID     int32 `db:"user_id" json:"user_id"`
}

type Product struct {
	ID          int32           `db:"id" json:"id"`
	Name        string          `db:"name" json:"name"`
	Description string          `db:"description" json:"description"`
	Cod         string          `db:"cod" json:"cod"`
	Price       decimal.Decimal `db:"price" json:"price"`
	Ipi         decimal.Decimal `db:"ipi" json:"ipi"`
	Active      bool
	UserID      int32 `db:"user_id" json:"user_id"`
	Company_ID  int32 `db:"company_id" json:"company_id"`
}

type Order struct {
	ID         int32        `db:"id"`
	Status     StatusType   `db:"status_type"`
	Shipping   ShippingType `db:"shipping"`
	UserID     int32        `db:"user_id"`
	PortageID  int32        `db:"portage_id"`
	CustomerID int32        `db:"customer_id"`
	EmployerID int32        `db:"employer_id"`
	Obervation string       `db:"observation"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type OrderItem struct {
	OrderID      int32 `db:"order_id"`
	ProductID    int32 `db:"product_id"`
	Quantity     int32
	PriceDecimal decimal.Decimal
	Discount     decimal.Decimal
}
