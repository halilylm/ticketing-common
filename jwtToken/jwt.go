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

type JWTToken struct {
	ExpiresIn time.Duration
	Secret    string
}

func NewJWTToken(expires time.Duration, secret string) *JWTToken {
	return &JWTToken{
		ExpiresIn: expires,
		Secret:    secret,
	}
}

func (j *JWTToken) GenerateToken(userClaim UserClaim) (string, error) {
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthClaim{
		UserClaim: userClaim,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.ExpiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	token, err := claim.SignedString([]byte("super_secret"))
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(token)
	return token, nil
}

func (j *JWTToken) ParseToken(tokenString string) (*UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
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
