package company_controller

import (
	"github.com/moura-dev/project-orders-golang/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/moura-dev/project-orders-golang/repository/company_repository"
	"github.com/moura-dev/project-orders-golang/services"
)

func Create(ctx *gin.Context) {
	var company models.Company

	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := services.GetUserIdFromContext(ctx)

	company, err := company_repository.Create(userId, company)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, company)
}

func Update(ctx *gin.Context) {
	var company models.Company

	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company, err := company_repository.Update(company)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, company)
}

func Delete(ctx *gin.Context) {
	userId := services.GetUserIdFromContext(ctx)

	strCompanyId := ctx.Param("id")
	companyId, err := strconv.Atoi(strCompanyId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = company_repository.Delete(userId, companyId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Company deleted successfully")
}

func Get(ctx *gin.Context) {
	userId := services.GetUserIdFromContext(ctx)

	companies, err := company_repository.Get(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, companies)
}
