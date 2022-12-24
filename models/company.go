package models

type Company struct {
	ID          int32  `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
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
	Plan        int32  `json:"plan"`
	IsActive    bool   `json:"is_active"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
}

type Companies struct {
	Companies []Company `json:"companies"`
}

type CreateCompany struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
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
}
