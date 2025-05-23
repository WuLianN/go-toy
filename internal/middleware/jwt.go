package middleware

import (
	"github.com/golang-jwt/jwt/v5"

	"errors"

	"github.com/WuLianN/go-toy/pkg/app"
	"github.com/WuLianN/go-toy/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			// token = c.GetHeader("token") // 也可以使用自定义header
			token = c.GetHeader("Authorization")
		}

		if token == "" {
			ecode = errcode.UnauthorizedAuthNotExist
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				if errors.Is(err, jwt.ErrTokenExpired) {
					ecode = errcode.UnauthorizedTokenTimeout
				} else {
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Next()
	}

}
