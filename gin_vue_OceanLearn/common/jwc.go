package common

import (
	"OceanLearn/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint

	jwt.StandardClaims
}

func ReleaseToken(user model.Go_gin_User) (string, error) {

	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := Claims{
		UserId: uint(user.Id),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "xz.tech",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//解析token

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
