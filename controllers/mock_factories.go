package controllers

import (
	models "base-project-api/models"
	"github.com/gin-gonic/gin"
)

func (mc *MockController) GetFactoryByID(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": models.Factory{
			ID:          1,
			Email:       "Andaluz@test.com.br",
			Name:        "Andaluz",
			FantasyName: "Andaluz",
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
			Commission:  "5",
			CreateAt:    "2020-01-01 00:00:00",
			UpdateAt:    "2020-01-01 00:00:00",
		},
	})
}

func (mc *MockController) GetAllFactories(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": models.Factories{
			Factories: []models.Factory{
				{ID: 1,
					Email:       "andaluz@test.com.br",
					Name:        "Andaluz",
					FantasyName: "Andaluz",
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
					Commission:  "5",
					CreateAt:    "2020-01-01 00:00:00",
					UpdateAt:    "2020-01-01 00:00:00",
				},
				{ID: 2,
					Email:       "Cleiber@test.com.br",
					Name:        "Cleiber",
					FantasyName: "Cleiber",
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
					Commission:  "5",
					CreateAt:    "2020-01-01 00:00:00",
					UpdateAt:    "2020-01-01 00:00:00",
				},
				{
					ID:          3,
					Email:       "LLMonteiro@test.com.br",
					Name:        "LLMonteiro",
					FantasyName: "LLMonteiro",
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
					Commission:  "5",
					CreateAt:    "2020-01-01 00:00:00",
					UpdateAt:    "2020-01-01 00:00:00"},
			},
		},
	})
}
