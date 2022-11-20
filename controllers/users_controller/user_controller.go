package users_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/moura-dev/project-orders-golang/models"
	"github.com/moura-dev/project-orders-golang/repository/user_repository"
	"github.com/moura-dev/project-orders-golang/services"
)

func Get(ctx *gin.Context) {

	userId := services.GetUserIdFromContext(ctx)

	user, err := user_repository.Get(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func Create(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = services.SHA256ENCODER(user.Password)

	user, err := user_repository.Create(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = ""

	ctx.JSON(http.StatusOK, user)
}

func Delete(ctx *gin.Context) {

	userId := services.GetUserIdFromContext(ctx)

	err := user_repository.Delete(userId)
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
			"error": err.Error(),
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
