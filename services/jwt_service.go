package services

import (
	"fmt"
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
