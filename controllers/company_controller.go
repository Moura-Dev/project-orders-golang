package controllers

import (
	"base-project-api/models"
	"base-project-api/repository"
	"base-project-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateCompany(ctx *gin.Context) {
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
	company, err = repository.CreateCompany(company)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, company)
}

func UpdateCompany(ctx *gin.Context) {
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

	company, err = repository.UpdateCompany(company)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, company)
}

func DeleteCompany(ctx *gin.Context) {

	companyID := ctx.Param("id")
	companyIDInt, err := strconv.Atoi(companyID)
	if err != nil {
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

	err = repository.DeleteCompany(userIdInt, companyIDInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Company deleted successfully")
}

func GetAllCompany(ctx *gin.Context) {
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
	companies, err := repository.GetAllCompany(userIdInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, companies)
}
