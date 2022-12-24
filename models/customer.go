package models

type Customer struct {
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

type Customers struct {
	Customers []Customer `json:"customers"`
}

type CreateCustomer struct {
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
}

type UpdateCustumer struct {
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
}

type DeleteCustumer struct {
	ID int32 `json:"id"`
}
