package jwtToken

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaim struct {
	Email  string `json:"email"`
	UserID string `json:"user_id"`
}

type AuthClaim struct {
	UserClaim
	jwt.RegisteredClaims
}

func GenerateToken(userClaim UserClaim, secret string, expiresIN time.Duration) (string, error) {
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthClaim{
		UserClaim: userClaim,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIN)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	token, err := claim.SignedString([]byte(secret))
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(token)
	return token, nil
}

func ParseToken(tokenString string, secret string) (*UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*AuthClaim)
	if ok && token.Valid {
		return nil, errors.New("invalid token")
	}
	return &claims.UserClaim, nil
}
