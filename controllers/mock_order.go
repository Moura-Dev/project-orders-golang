package controllers

import (
	models "base-project-api/models"
	"github.com/gin-gonic/gin"
)

func (mc *MockController) GetOrderByID(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": models.Order{
			ID:         1,
			FactoryID:  1,
			CompanyID:  1,
			CustomerID: 1,
			CreateAt:   "2020-01-01 00:00:00",
			UpdateAt:   "2020-01-01 00:00:00",
		},
	})
}

func (mc *MockController) GetAllOrders(ctx *gin.Context) {
	ctx.JSON(
		200,
		gin.H{
			"data": models.Orders{
				[]models.Order{
					{ID: 1,
						FactoryID:  1,
						CompanyID:  1,
						CustomerID: 1,
						CreateAt:   "2020-01-01 00:00:00",
						UpdateAt:   "2020-01-01 00:00:00",
					},
					{ID: 2,
						FactoryID:  2,
						CompanyID:  2,
						CustomerID: 2,
						CreateAt:   "2020-01-01 00:00:00",
						UpdateAt:   "2020-01-01 00:00:00",
					},
					{
						ID:         2,
						FactoryID:  2,
						CompanyID:  2,
						CustomerID: 2,
						CreateAt:   "2020-01-01 00:00:00",
						UpdateAt:   "2020-01-01 00:00:00",
					},
				},
			},
		},
	)

}
