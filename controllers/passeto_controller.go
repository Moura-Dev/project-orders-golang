package controllers

import (
	"base-project-api/models"
	"base-project-api/repository"
	"base-project-api/services/token"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type IToken interface {
}

type Token struct {
	tokenMaker token.Maker
}

func NewToken() IToken {
	return Token{}
}

func (t Token) RenewAccessToken(ctx *gin.Context) {
	var req models.RenewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	refreshPayload, err := t.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	session, err := repository.GetSessionByID(refreshPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if session.IsBlocked {
		err := fmt.Errorf("blocked session")
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	if session.Email != refreshPayload.Email {
		err := fmt.Errorf("incorrect session user")
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	if session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("mismatched session token")
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	if time.Now().After(session.ExpiresAt) {
		err := fmt.Errorf("expired session")
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	accessTokenDurationStr := os.Getenv("AccessTokenDuration")
	accessTokenDuration, _ := time.ParseDuration(accessTokenDurationStr)
	accessToken, accessPayload, err := t.tokenMaker.CreateToken(
		refreshPayload.Email,
		accessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rsp := models.RenewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, rsp)
}
