package ygjwt

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("youguan_netdisk")

type Claims struct {
	UserId string
	jwt.StandardClaims
}

// 根据唯一识别号 indentify 发放 token
func ReleseToken(indentify string) (string, error) {
	// 7 天过期时间
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := Claims{
		UserId: indentify,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "youguan",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Error(err)
		return "", err
	}
	return tokenString, nil
}

// 解析 token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := new(Claims)
	if token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	}); err != nil {
		log.Error(err)
		return nil, nil, err
	} else {
		return token, claims, nil
	}
}
