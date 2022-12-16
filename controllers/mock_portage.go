package controllers

import (
	models "base-project-api/models"
	"github.com/gin-gonic/gin"
)

func (mc *MockController) GetPortageByID(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": models.Factory{
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

func (mc *MockController) GetAllPortages(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": models.Factories{
			Factories: []models.Factory{
				{ID: 1,
					Email:       "teste@test.com.br",
					Name:        "Transpordora IkaroShow",
					FantasyName: "LTDA Transporte",
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
					Name:        "Rafits Transporte",
					FantasyName: "Rafits LTDA",
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
					Name:        "Correios",
					FantasyName: "Correios LTDA",
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
