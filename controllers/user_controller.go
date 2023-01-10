package controllers

import (
	"base-project-api/models"
	"base-project-api/repository"
	"base-project-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

//func CreateUser(ctx *gin.Context) {
//	var user models.User
//	user.Password = services.SHA256ENCODER(user.Password)
//
//	if err := ctx.ShouldBindJSON(&user); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	user.Password = services.SHA256ENCODER(user.Password)
//
//	user, err := repository.CreateUser(user)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	user.Password = ""
//
//	ctx.JSON(http.StatusOK, user)
//}
//
//func GetUserInfo(ctx *gin.Context) {
//	token := ctx.Request.Header.Get("authorization")
//	token = token[7:]
//	userId, err := services.NewJWTService().GetUserIdFromToken(token)
//	if err != nil {
//		ctx.JSON(401, gin.H{
//			"error": err.Error(),
//		})
//		return
//	}
//	userIdInt, err := strconv.Atoi(userId)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	user, err := repository.GetUserInfo(userIdInt)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	user.Password = ""
//
//	ctx.JSON(http.StatusOK, user)
//}
//
//func (c *Controller) CreateSeller(ctx *gin.Context) {
//	var seller models.Seller
//
//	if err := ctx.ShouldBindJSON(&seller); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	token := ctx.Request.Header.Get("authorization")
//	token = token[7:]
//	userId, err := services.NewJWTService().GetUserIdFromToken(token)
//	if err != nil {
//		ctx.JSON(401, gin.H{
//			"error": err.Error(),
//		})
//		return
//	}
//	userIdInt, err := strconv.Atoi(userId)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	seller.UserID = int32(userIdInt)
//
//	seller, err = repository.CreateSeller(seller)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, seller)
//}
//
//func GetSellerById(ctx *gin.Context) {
//	var seller models.Seller
//
//	sellerID := ctx.Param("id")
//	sellerIdInt, err := strconv.Atoi(sellerID)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	seller, err = repository.GetSellerById(sellerIdInt)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, seller)
//}
//
//func GetAllSellers(ctx *gin.Context) {
//	token := ctx.Request.Header.Get("authorization")
//	token = token[7:]
//	userId, err := services.NewJWTService().GetUserIdFromToken(token)
//	if err != nil {
//		ctx.JSON(401, gin.H{
//			"error": err.Error(),
//		})
//		return
//	}
//	userIdInt, err := strconv.Atoi(userId)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	sellers, err := repository.GetAllSellers(userIdInt)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, sellers)
//}
//
//func DeleteSeller(ctx *gin.Context) {
//
//	sellerID := ctx.Param("id")
//	sellerIdInt, err := strconv.Atoi(sellerID)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = repository.DeleteSellerByID(sellerIdInt)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, "Deleted Successfully")
//}

func (t Token) LoginUser(ctx *gin.Context) { // Get User in db
	data := models.LoginUserRequest{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := repository.GetUser(data.Email)
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
	TokenDurationStr := os.Getenv("AccessTokenDuration")
	accessTokenDuration, _ := time.ParseDuration(TokenDurationStr)
	accessToken, accessPayload, _ := t.tokenMaker.CreateToken(
		user.Email,
		accessTokenDuration,
	)

	refreshTokenDurationStr := os.Getenv("AccessTokenDuration")
	refreshTokenDuration, _ := time.ParseDuration(refreshTokenDurationStr)
	refreshToken, refreshPayload, err := t.tokenMaker.CreateToken(
		user.Email,
		refreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//create ID           uuid.UUID
	//	Email        string
	//	RefreshToken string
	//	UserAgent    string
	//	ClientIp     string
	//	IsBlocked    bool
	//	ExpiresAt    time.Time
	session := models.CreateSessionParams{
		Email:        user.Email,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    time.Now().Add(time.Hour * 24),
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rsp := models.LoginUserResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  models.UserResponse{ID: user.ID, Email: user.Email},
	}
	ctx.JSON(http.StatusOK, rsp)
}
