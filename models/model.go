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
	ID          int32  `db:"id"`
	Street      string `db:"street" json:"street"`
	Number      string `db:"number" json:"number"`
	ZipCode     string `db:"zip_code" json:"zip_code"`
	City        string `db:"city" json:"city"`
	State       string `db:"state" json:"state"`
	District    string `db:"district" json:"district"`
	Reference   string `db:"reference" json:"reference"`
	AddressName string `db:"addressname" json:"address_name"`
	UserID      int32  `db:"user_id" json:"user_id"`
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
	ID         int32     `db:"id" json:"id"`
	Status     string    `db:"status" json:"status"`
	Shipping   string    `db:"shipping" json:"shipping"`
	UserID     int32     `db:"user_id" json:"user_id"`
	PortageID  int32     `db:"portage_id" json:"portage_id"`
	CustomerID int32     `db:"customer_id" json:"customer_id"`
	Purchaser  string    `db:"purchaser" json:"purchaser"`
	Obervation string    `db:"observation" json:"observation"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}

type OrderItem struct {
	OrderID      int32           `db:"order_id" json:"order_id"`
	ProductID    int32           `db:"product_id" json:"product_id"`
	Quantity     int32           `db:"quantity" json:"quantity"`
	PriceDecimal decimal.Decimal `db:"price" json:"price"`
	Discount     decimal.Decimal `db:"discount" json:"discount"`
}
