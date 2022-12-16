package controllers

import (
	"base-project-api/models"
	"github.com/gin-gonic/gin"
)

func (mc *MockController) GetCatalogByID(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": models.Catalog{
			ID:          1,
			CompanyID:   1,
			Name:        "teste",
			Description: "teste",
			CreateAt:    "2020-01-01 00:00:00",
			UpdateAt:    "2020-01-01 00:00:00",
		},
	})
}

func (mc *MockController) GetAllCatalogs(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": models.Catalogs{
			Catalogs: []models.Catalog{
				{
					ID:          1,
					CompanyID:   1,
					Name:        "teste1",
					Description: "teste",
					CreateAt:    "2020-01-01 00:00:00",
					UpdateAt:    "2020-01-01 00:00:00",
				},
				{
					ID:          2,
					CompanyID:   2,
					Name:        "teste2",
					Description: "teste2",
					CreateAt:    "2020-01-01 00:00:00",
					UpdateAt:    "2020-01-01 00:00:00",
				},
				{
					ID:          3,
					CompanyID:   3,
					Name:        "teste3",
					Description: "teste2",
					CreateAt:    "2020-01-01 00:00:00",
					UpdateAt:    "2020-01-01 00:00:00",
				},
			},
		},
	})
}
