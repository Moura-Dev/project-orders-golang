package controllers

import (
	"base-project-api/models"
	"base-project-api/repository"
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

	address, err := repository.CreateAddress(address)
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

	address, err := repository.UpdateAddress(address)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, address)
}

func GetAddresses(ctx *gin.Context) {
	var address models.Address
	//user id
	userID := ctx.Param("id")
	userIDIdInt, err := strconv.Atoi(userID)

	address, err = repository.GetAddress(userIDIdInt)
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
	err = repository.DeleteAddress(addressIDInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Address deleted successfully"})
}
