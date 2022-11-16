package users_controller

import (
	"base-project-api/models"
	"base-project-api/repository/user_repository"
	"base-project-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Get(ctx *gin.Context) {

	userId, err := services.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := user_repository.Get(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = ""

	ctx.JSON(http.StatusOK, user)
}

func Create(ctx *gin.Context) {
	var user models.User
	user.Password = services.SHA256ENCODER(user.Password)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = services.SHA256ENCODER(user.Password)

	user, err := user_repository.Create(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = ""

	ctx.JSON(http.StatusOK, user)
}

func Delete(ctx *gin.Context) {

	sellerID := ctx.Param("id")
	sellerIdInt, err := strconv.Atoi(sellerID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user_repository.Delete(sellerIdInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Deleted Successfully")
}

func Login(ctx *gin.Context) { // Get User in db
	var data models.User

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := user_repository.Login(data.Login)
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
	token, err := services.NewJWTService().GenerateToken(int(user.Id))
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
