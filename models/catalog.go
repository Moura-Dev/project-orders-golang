package models

type Catalog struct {
	ID          int32  `json:"id"`
	CompanyID   int32  `json:"company_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
}

type Catalogs struct {
	Catalogs []Catalog `json:"catalogs"`
}

type CreateCatalog struct {
	CompanyID   int32  `json:"company_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCatalog struct {
	ID          int32  `json:"id"`
	CompanyID   int32  `json:"company_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteCatalog struct {
	ID int32 `json:"id"`
}
