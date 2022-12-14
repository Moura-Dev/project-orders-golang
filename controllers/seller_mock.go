package controllers

import "github.com/gin-gonic/gin"

type Seller struct {
	ID     int32  `db:"id" json:"-"`
	UserID int32  `db:"user_id" json:"user_id"`
	Name   string `db:"name" json:"name"`
	TaxID  string `db:"tax_id" json:"tax_id"`
}

func (mc *MockController) CreateSeller(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": Seller{
			ID:     1,
			UserID: 1,
			Name:   "Seller Name",
			TaxID:  "123456789",
		},
	})
}

func (mc *MockController) HellowControllers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func (mc *MockController) GetSellerById(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": Seller{
			ID:     1,
			UserID: 1,
			Name:   "Seller Name",
			TaxID:  "123456789",
		},
	})
}

func (mc *MockController) GetAllSellers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": []Seller{
			{
				ID:     2,
				UserID: 2,
				Name:   "Seller name 1",
				TaxID:  "123198765",
			},
			{
				ID:     1,
				UserID: 1,
				Name:   "Seller Name",
				TaxID:  "123456789",
			},
			{
				ID:     3,
				UserID: 3,
				Name:   "Seller name 2",
				TaxID:  "123456789",
			},
		},
	})
}
