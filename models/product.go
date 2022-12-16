package models

import "github.com/shopspring/decimal"

type Product struct {
	ID          int32  `json:"id"`
	CompanyID   int32  `json:"company_id"`
	FactoryID   int32  `json:"factory_id"`
	CatalogID   int32  `json:"catalog_id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Price       string `json:"price"`
	Reference   string `json:"reference"`
	Description string `json:"description"`
	ImageURL    string `json:"image"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
}

type Products struct {
	Products []Product `json:"products"`
}

type CreateItem struct {
	CompanyID   int32           `json:"company_id"`
	FactoryID   int32           `json:"factory_id"`
	CatalogID   int32           `json:"catalog_id"`
	Name        string          `json:"name"`
	Code        string          `json:"code"`
	Price       decimal.Decimal `json:"price"`
	Reference   string          `json:"reference"`
	Description string          `json:"description"`
	ImageURL    string          `json:"image"`
}

type UpdateItem struct {
	ID          int32           `json:"id"`
	CompanyID   int32           `json:"company_id"`
	FactoryID   int32           `json:"factory_id"`
	CatalogID   int32           `json:"catalog_id"`
	Name        string          `json:"name"`
	Code        string          `json:"code"`
	Price       decimal.Decimal `json:"price"`
	Reference   string          `json:"reference"`
	Description string          `json:"description"`
	ImageURL    string          `json:"image"`
}

type DeleteItem struct {
	ID int32 `json:"id"`
}
