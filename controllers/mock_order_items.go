package controllers

import (
	models "base-project-api/models"
	"github.com/gin-gonic/gin"
)

func (mc *MockController) GetOrderItemByID(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": models.OrderItem{
			ID:        1,
			OrderID:   1,
			ProductID: 1,
			Quantity:  1,
			Price:     1,
			CreateAt:  "2020-01-01 00:00:00",
			UpdateAt:  "2020-01-01 00:00:00",
		},
	})
}

func (mc *MockController) GetAllOrdersItems(ctx *gin.Context) {
	ctx.JSON(
		200,
		gin.H{
			"data": models.OrderItems{
				OrderItems: []models.OrderItem{
					{ID: 1,
						OrderID:   1,
						ProductID: 1,
						Quantity:  1,
						Price:     1,
						CreateAt:  "2020-01-01 00:00:00",
						UpdateAt:  "2020-01-01 00:00:00",
					},
					{ID: 2,
						OrderID:   2,
						ProductID: 2,
						Quantity:  2,
						Price:     2,
						CreateAt:  "2020-01-01 00:00:00",
						UpdateAt:  "2020-01-01 00:00:00",
					},
					{
						ID:        3,
						OrderID:   3,
						ProductID: 3,
						Quantity:  3,
						Price:     13,
						CreateAt:  "2020-01-01 00:00:00",
						UpdateAt:  "2020-01-01 00:00:00",
					},
				},
			},
		},
	)

}
