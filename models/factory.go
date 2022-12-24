package models

type Factory struct {
	ID          int32  `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	FantasyName string `json:"fantasy_name"`
	CPFCNPJ     string `json:"cpfcnpj"`
	IE          string `json:"ie"`
	Phone       string `json:"phone"`
	Website     string `json:"website"`
	LogoURL     string `json:"logo"`
	Street      string `json:"steet"`
	Number      string `json:"number"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zipcode     string `json:"zipcode"`
	Commission  string `json:"commission"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
}

type Factories struct {
	Factories []Factory `json:"factories"`
}

type CreateFactory struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	FantasyName string `json:"fantasy_name"`
	CPFCNPJ     string `json:"cpfnpj"`
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

type UpdateFactory struct {
	ID          int32  `json:"id"`
	Email       string `json:"email"`
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

type DeleteFactory struct {
	ID int32 `json:"id"`
}
