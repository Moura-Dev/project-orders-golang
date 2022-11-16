package company_controller

import (
	"base-project-api/models"
	"base-project-api/repository/company_repository"
	"base-project-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Create(ctx *gin.Context) {
	var company models.Company

	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := ctx.Request.Header.Get("authorization")
	token = token[7:]
	userId, err := services.NewJWTService().GetUserIdFromToken(token)
	if err != nil {
		ctx.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	company.UserID = int32(userIdInt)
	company, err = company_repository.Create(company)
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

	token := ctx.Request.Header.Get("authorization")
	token = token[7:]
	userId, err := services.NewJWTService().GetUserIdFromToken(token)
	if err != nil {
		ctx.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	company.UserID = int32(userIdInt)

	company, err = company_repository.Update(company)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, company)
}

func Delete(ctx *gin.Context) {
	userId, err := services.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
	token := ctx.Request.Header.Get("authorization")
	token = token[7:]
	userId, err := services.NewJWTService().GetUserIdFromToken(token)
	if err != nil {
		ctx.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	companies, err := company_repository.Get(userIdInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, companies)
}
