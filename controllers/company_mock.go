package controllers

import "github.com/gin-gonic/gin"

type Company struct {
	ID          int32  `db:"id"`
	UserID      int32  `db:"user_id" json:"user_id"`
	Name        string `db:"name" json:"name"`
	Cnpj        string `db:"cnpj" json:"cnpj"`
	FantasyName string `db:"fantasy_name" json:"fantasy_name"`
	IE          string `db:"ie" json:"ie"`
	Fone        string `db:"fone" json:"fone"`
	Fone2       string `db:"fone2" json:"fone2"`
	Logo        string `db:"logo" json:"logo"`
	CompanyType string `db:"company_type" json:"company_type"`
	AddressID   int32  `db:"address_id" json:"address_id"`
}

func (mc *MockController) GetAllCompany(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": []Company{
			{
				ID:          1,
				UserID:      1,
				Name:        "Company Name",
				Cnpj:        "123456789",
				FantasyName: "Fantasy Name",
				IE:          "123456789",
				Fone:        "123456789",
				Fone2:       "123456789",
				Logo:        "Logo",
				CompanyType: "Company Type",
				AddressID:   1,
			},
			{
				ID:          2,
				UserID:      2,
				Name:        "Company Name 2",
				Cnpj:        "123456789",
				FantasyName: "Fantasy Name 2",
				IE:          "123456789",
				Fone:        "123456789",
				Fone2:       "123456789",
				Logo:        "Logo 2",
				CompanyType: "Company Type 2",
				AddressID:   2,
			},
		},
	})
}
