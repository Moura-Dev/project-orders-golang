package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type StatusType struct {
	QUOTE     string
	APPROVED  string
	CANCELLED string
}

type ShippingType struct {
	CIF string
	FOB string
}

type Product struct {
	Id          int32           `db:"id" json:"id"`
	UserId      int32           `db:"user_id" json:"user_id"`
	CompanyId   int32           `db:"company_id" json:"company_id"`
	Name        string          `db:"name" json:"name"`
	Description string          `db:"description" json:"description"`
	COD         string          `db:"cod" json:"cod"`
	Price       decimal.Decimal `db:"price" json:"price"`
	IPI         decimal.Decimal `db:"ipi" json:"ipi"`
	Active      bool            `db:"active" json:"active"`
	CreatedAt   time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time       `db:"updated_at" json:"updated_at"`
}

type Order struct {
	Id                int32           `db:"id" json:"id"`
	Status            string          `db:"status" json:"status"`
	UserId            int32           `db:"user_id" json:"user_id"`
	ShippingCompanyId int32           `db:"shipping_company_id" json:"shipping_company_id"`
	CustomerId        int32           `db:"customer_id" json:"customer_id"`
	Purchaser         string          `db:"purchaser" json:"purchaser"`
	Observation       string          `db:"observation" json:"observation"`
	Shipping          ShippingType    `db:"shipping" json:"shipping"`
	Items             []OrderItem     `db:"items" json:"items"`
	Total             decimal.Decimal `json:"total"`
	CreatedAt         time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time       `db:"updated_at" json:"updated_at"`
}

type OrderItem struct {
	Id        int32           `db:"id" json:"id"`
	OrderId   int32           `db:"orderId" json:"orderId"`
	ProductId int32           `db:"product_id" json:"product_id"`
	Quantity  int32           `db:"quantity" json:"quantity"`
	Price     decimal.Decimal `db:"price" json:"price"`
	Discount  decimal.Decimal `db:"discount" json:"discount"`
	Total     decimal.Decimal `db:"total" json:"total"`
}
