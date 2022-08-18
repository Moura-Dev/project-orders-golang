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
	Portage  string
	Customer string
	Factory  string
}

type ShippingType struct {
	Cif string
	Fob string
}
type User struct {
	ID        int32  `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
	Login     string
	Password  string
	Active    bool
}

type Seller struct {
	UserID int32 `db:"user_id"`
	Name   string
	TaxID  string `db:"tax_id`
}

type Company struct {
	ID          int32 `db:"id"`
	UserID      int32 `db:"user_id"`
	Name        string
	Cnpj        string
	FantasyName string `db:"fantasy_name"`
	IE          string `db:"IE`
	Fone        string
	Fone2       string
	Logo        string
	CompanyType CompanyType `db:"company_type`
	AdressID    int32
}

type Adress struct {
	ID         int32 `db:"id"`
	Street     string
	Number     string
	ZipCode    string `db:"zip_code"`
	City       string
	State      string
	District   string
	Reference  string
	AdressName string
}

type Product struct {
	ID          int32 `db:"id"`
	Name        string
	Description string
	Cod         string
	Price       decimal.Decimal
	Ipi         decimal.Decimal
	Active      bool
	Company_ID  int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Order struct {
	ID         int32        `db:"id"`
	Status     StatusType   `db:"status_type"`
	Shipping   ShippingType `db:"shipping"`
	UserID     int32        `db:"user_id"`
	PortageID  int32        `db:"portage_id"`
	CustomerID int32        `db:"customer_id"`
	EmployerID int32        `db:"employer_id"`
	Obervation string
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
