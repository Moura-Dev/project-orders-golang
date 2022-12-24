package models

type Order struct {
	ID          int32  `json:"id"`
	CompanyID   int32  `json:"company_id"`
	FactoryID   int32  `json:"factory_id"`
	CustomerID  int32  `json:"customer_id"`
	UrlPdf      string `json:"url_pdf"`
	ExpireOrder string `json:"expire_order"`
	IsActive    bool   `json:"is_active"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
}

type Orders struct {
	Orders []Order `json:"orders"`
}

type CreateOrder struct {
	CompanyID  int32 `json:"company_id"`
	FactoryID  int32 `json:"factory_id"`
	CustomerID int32 `json:"customer_id"`
}

type UpdateOrder struct {
	ID          int32  `json:"id"`
	CompanyID   int32  `json:"company_id"`
	FactoryID   int32  `json:"factory_id"`
	CustomerID  int32  `json:"customer_id"`
	UrlPdf      string `json:"url_pdf"`
	ExpireOrder string `json:"expire_order"`
}

type DeleteOrder struct {
	ID int32 `json:"id"`
}
