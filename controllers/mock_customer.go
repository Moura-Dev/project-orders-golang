package controllers

import (
	"base-project-api/models"
	"github.com/gin-gonic/gin"
)

type Customer struct {
	ID          int32  `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	FantasyName string `json:"fantasy_name"`
	CNPJ        string `json:"cnpj"`
	IE          string `json:"ie"`
	Phone       string `json:"phone"`
	Website     string `json:"website"`
	LogoURL     string `json:"logo"`
	Street      string `json:"street"`
	Number      string `json:"number"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zipcode     string `json:"zipcode"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
}

func (mc *MockController) GetCustomerByID(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": models.Company{
			ID:          1,
			Email:       "teste@test.com.br",
			Name:        "teste",
			FantasyName: "teste",
			CPFCNPJ:     "123456789",
			IE:          "123456789",
			Phone:       "123456789",
			Website:     "teste.com.br",
			LogoURL:     "teste.com.br",
			Street:      "rua adele",
			Number:      "123",
			City:        "sao paulo",
			State:       "sp",
			Zipcode:     "123456789",
			CreateAt:    "2020-01-01 00:00:00",
			UpdateAt:    "2020-01-01 00:00:00",
		},
	})
}

func (mc *MockController) GetAllCustomers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": models.Customers{
			Customers: []models.Customer{
				{ID: 1,
					Email:       "teste@test.com.br",
					Name:        "teste",
					FantasyName: "teste",
					CPFCNPJ:     "123456789",
					IE:          "123456789",
					Phone:       "123456789",
					Website:     "teste.com.br",
					LogoURL:     "teste.com.br",
					Street:      "rua adele",
					Number:      "123",
					City:        "sao paulo",
					State:       "sp",
					Zipcode:     "123456789",
					CreateAt:    "2020-01-01 00:00:00",
					UpdateAt:    "2020-01-01 00:00:00",
				},
				{ID: 2,
					Email:       "teste@test.com.br",
					Name:        "teste",
					FantasyName: "teste",
					CPFCNPJ:     "123456789",
					IE:          "123456789",
					Phone:       "123456789",
					Website:     "teste.com.br",
					LogoURL:     "teste.com.br",
					Street:      "rua adele 210",
					Number:      "123",
					City:        "sao paulo",
					State:       "sp",
					Zipcode:     "123456789",
					CreateAt:    "2020-01-01 00:00:00",
					UpdateAt:    "2020-01-01 00:00:00",
				},
				{
					ID:          3,
					Email:       "teste@test.com.br",
					Name:        "teste",
					FantasyName: "teste",
					CPFCNPJ:     "123456789",
					IE:          "123456789",
					Phone:       "123456789",
					Website:     "teste.com.br",
					LogoURL:     "teste.com.br",
					Street:      "rua adele 210",
					Number:      "123",
					City:        "sao paulo",
					State:       "sp",
					Zipcode:     "123456789",
					CreateAt:    "2020-01-01 00:00:00",
					UpdateAt:    "2020-01-01 00:00:00"},
			},
		},
	})
}
