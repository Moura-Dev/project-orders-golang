package controllers

import (
	"base-project-api/models"
	"base-project-api/repository"
	"base-project-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(ctx *gin.Context) {
	var user models.User
	user.Password = services.SHA256ENCODER(user.Password)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = services.SHA256ENCODER(user.Password)

	user, err := repository.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = ""

	ctx.JSON(http.StatusOK, user)
}

func GetUserInfo(ctx *gin.Context) {
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

	user, err := repository.GetUserInfo(userIdInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = ""

	ctx.JSON(http.StatusOK, user)
}

func CreateSeller(ctx *gin.Context) {
	var seller models.Seller

	if err := ctx.ShouldBindJSON(&seller); err != nil {
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
	seller.UserID = int32(userIdInt)

	seller, err = repository.CreateSeller(seller)
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
	sellers, err := repository.GetAllSellers(userIdInt)
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

func Login(ctx *gin.Context) { // Get User in db
	data := models.User{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := repository.GetUser(data.Login)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	if user.Password != services.SHA256ENCODER(data.Password) {
		ctx.JSON(401, gin.H{
			"error": "invalid credentials",
		})
		return
	}
	token, err := services.NewJWTService().GenerateToken(int(user.ID))
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"access_token": token,
	})
}
