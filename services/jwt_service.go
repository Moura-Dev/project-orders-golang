package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtService struct {
	secret string
	issure string
}

type Claim struct {
	jwt.StandardClaims
}

func NewJWTService() *jwtService {
	return &jwtService{
		secret: "secret-key",
		issure: "orders-api",
	}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	userIDStr := strconv.Itoa(userID)

	claim := Claim{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    s.issure,
			Id:        userIDStr,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenEncode, err := token.SignedString([]byte(s.secret))

	if err != nil {
		return "", err
	}
	return tokenEncode, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.secret), nil
	})

	return err == nil
}

func (s *jwtService) GetUserIdFromToken(t string) (string, error) {
	tokenString := t
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})

	if token.Valid {
		return claims["jti"].(string), nil
	}
	return "", err
}

func GetUserIdFromContext(ctx *gin.Context) (int, error) {
	token := ctx.Request.Header.Get("authorization")
	token = token[7:]
	userId, err := NewJWTService().GetUserIdFromToken(token)
	if err != nil {
		ctx.JSON(401, gin.H{
			"error": err.Error(),
		})
		return 0, nil
	}
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return 0, err
	}
	return userIdInt, err
}
