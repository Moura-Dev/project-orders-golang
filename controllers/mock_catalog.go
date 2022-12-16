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
					Name:        "Catalogo Andaluz",
					Description: "Trapo",
					CreateAt:    "2020-01-01 00:00:00",
					UpdateAt:    "2020-01-01 00:00:00",
				},
				{
					ID:          2,
					CompanyID:   2,
					Name:        "Catalogo Loth",
					Description: "Algodão",
					CreateAt:    "2020-01-01 00:00:00",
					UpdateAt:    "2020-01-01 00:00:00",
				},
				{
					ID:          3,
					CompanyID:   3,
					Name:        "Catalogo Cleiber",
					Description: "Prego Telheiro",
					CreateAt:    "2020-01-01 00:00:00",
					UpdateAt:    "2020-01-01 00:00:00",
				},
			},
		},
	})
}
