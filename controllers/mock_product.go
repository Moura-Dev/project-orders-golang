package controllers

import (
	"base-project-api/models"
	"github.com/gin-gonic/gin"
)

func (mc *MockController) GetProductByID(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": models.Product{
			ID:          1,
			Name:        "teste",
			Description: "teste",
			Price:       "1000",
			CompanyID:   1,
			FactoryID:   1,
			CatalogID:   1,
			Code:        "MV123456789",
			Reference:   "123456789",
			ImageURL:    "teste.com.br",
			CreateAt:    "2020-01-01 00:00:00",
			UpdateAt:    "2020-01-01 00:00:00",
		},
	})
}

func (mc *MockController) GetAllProducts(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": models.Products{
			Products: []models.Product{
				{
					ID:          1,
					Name:        "Caneta Azul",
					Description: "Caneta",
					Price:       "100",
					CompanyID:   1,
					FactoryID:   1,
					CatalogID:   1,
					Code:        "MVC1203",
					Reference:   "",
					ImageURL:    "caneta.com.br",
					CreateAt:    "2020-01-01 00:00:00",
					UpdateAt:    "2020-01-01 00:00:00",
				},
				{
					ID:          2,
					Name:        "livro black",
					Description: "livro black",
					Price:       "10",
					CompanyID:   1,
					FactoryID:   1,
					CatalogID:   1,
					Code:        "MVL12",
					Reference:   "123456789",
					ImageURL:    "teste.com.br",
					CreateAt:    "2020-01-01 00:00:00",
					UpdateAt:    "2020-01-01 00:00:00",
				},
				{
					ID:          3,
					Name:        "Mouse gamer",
					Description: "Mouse",
					Price:       "100",
					CompanyID:   1,
					FactoryID:   1,
					CatalogID:   1,
					Code:        "MVM12",
					Reference:   "123456789",
					ImageURL:    "teste.com.br",
					CreateAt:    "2020-01-01 00:00:00",
					UpdateAt:    "2020-01-01 00:00:00",
				},
			},
		},
	})
}
