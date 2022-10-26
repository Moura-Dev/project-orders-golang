package controllers

import (
	"base-project-api/models"
	"base-project-api/repository"
	"base-project-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateAddress(ctx *gin.Context) {
	var address models.Address

	if err := ctx.ShouldBindJSON(&address); err != nil {
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
	address.UserID = int32(userIdInt)
	address, err = repository.CreateAddress(address)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, address)
}

func UpdateAddress(ctx *gin.Context) {
	var address models.Address

	if err := ctx.ShouldBindJSON(&address); err != nil {
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
	address.UserID = int32(userIdInt)

	address, err = repository.UpdateAddress(address)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, address)
}

func GetAllAddress(ctx *gin.Context) {
	var address []models.Address
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

	address, err = repository.GetAllAddress(userIdInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, address)
}

func DeleteAddress(ctx *gin.Context) {
	addressID := ctx.Param("id")
	addressIDInt, err := strconv.Atoi(addressID)
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
	err = repository.DeleteAddress(userIdInt, addressIDInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Address deleted successfully"})
}
