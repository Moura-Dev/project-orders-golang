package controllers

import (
	"base-project-api/models"
	"base-project-api/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repository.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func GetUserInfo(ctx *gin.Context) {
	var user models.User
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err = repository.GetUserInfo(userIdInt, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func CreateSeller(ctx *gin.Context) {
	var seller models.Seller

	if err := ctx.ShouldBindJSON(&seller); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seller, err := repository.CreateSeller(seller)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, seller)
}

func GetSellerById(ctx *gin.Context) {
	var seller models.Seller
	sellerID := ctx.Param("id")
	sellerIdInt, err := strconv.Atoi(sellerID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	seller, err = repository.GetSellerById(sellerIdInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, seller)
}

func GetAllSellers(ctx *gin.Context) {
	sellers, err := repository.GetAllSellers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sellers)
}

func DeleteSeller(ctx *gin.Context) {

	sellerID := ctx.Param("id")
	sellerIdInt, err := strconv.Atoi(sellerID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = repository.DeleteSellerByID(sellerIdInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Deleted Successfully")
}
