package api

import (
	"errors"

	"github.com/WuLianN/go-toy/pkg/app"
	"github.com/WuLianN/go-toy/pkg/errcode"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type TokenInfo struct {
	UserId   uint32
	UserName string
}

func GetTokenInfo(token string) (*errcode.Error, *TokenInfo) {
	ecode := errcode.Success
	claims, err := app.ParseToken(token)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			ecode = errcode.UnauthorizedTokenTimeout
		} else {
			ecode = errcode.UnauthorizedTokenError
		}
	}

	if ecode != errcode.Success {
		return ecode, nil
	}

	tokenInfo := TokenInfo{UserId: claims.UserId, UserName: claims.UserName}

	return nil, &tokenInfo
}

func GetToken(c *gin.Context) string {
	var token string
	if s, exist := c.GetQuery("token"); exist {
		token = s
	} else {
		token = c.GetHeader("Authorization")
	}

	return token
}
