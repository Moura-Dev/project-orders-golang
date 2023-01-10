package models

type Portage struct {
	ID          int32  `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	FantasyName string `json:"fantasy_name"`
	CPFCNPJ     string `json:"cpfcnpj"`
	IE          string `json:"ie"`
	Phone       string `json:"phone"`
	Website     string `json:"website"`
	LogoURL     string `json:"logo"`
	Street      string `json:"street"`
	Number      string `json:"number"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zipcode     string `json:"zipcode"`
	IsActive    bool   `json:"is_active"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
}

type Portages struct {
	Portages []Company `json:"portages"`
}

type CreatePortage struct {
	Emai        string `json:"email"`
	Name        string `json:"name"`
	FantasyName string `json:"fantasy_name"`
	CPFCNPJ     string `json:"cpfcnpj"`
	IE          string `json:"ie"`
	Phone       string `json:"phone"`
	Website     string `json:"website"`
	LogoURL     string `json:"logo"`
	Steet       string `json:"steet"`
	Number      string `json:"number"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zipcode     string `json:"zipcode"`
}

type UpdatePortage struct {
	ID          int32  `json:"id"`
	Emai        string `json:"email"`
	Name        string `json:"name"`
	FantasyName string `json:"fantasy_name"`
	CPFCNPJ     string `json:"cpfcnpj"`
	IE          string `json:"ie"`
	Phone       string `json:"phone"`
	Website     string `json:"website"`
	LogoURL     string `json:"logo"`
	Steet       string `json:"steet"`
	Number      string `json:"number"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zipcode     string `json:"zipcode"`
}

type DeletePortage struct {
	ID int32 `json:"id"`
}
