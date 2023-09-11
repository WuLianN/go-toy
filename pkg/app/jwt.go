package app

import (
	"time"

	"github.com/WuLianN/go-toy/global"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId uint `json:"id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

func GetJWTSecret() string {
	return global.JWTSetting.Secret
}

func GenerateToken(UserId uint, UserName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		UserId:    UserId,
		UserName:  UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt: jwt.NewNumericDate(nowTime),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(GetJWTSecret()))

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetJWTSecret()), nil
	})

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
