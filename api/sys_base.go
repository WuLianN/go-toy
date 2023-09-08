package api

import (
	"github.com/gin-gonic/gin"
	"github.com/WuLianN/go-toy/internal/service"
	"github.com/WuLianN/go-toy/pkg/app"
	"github.com/WuLianN/go-toy/global"
	"github.com/WuLianN/go-toy/pkg/errcode"
)

type BaseApi struct {}

// @Summary 测试Ping
// @Produce json
// @Success 1 {string} string "成功"
// @Router /api/ping [get]
func (b *BaseApi) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// @Summary 登录
// @Produce json
// @Tags user
// @Param user_name body string true "用户名" 
// @Param password body string true "密码" 
// @Success 1 {string} string "成功"
// @Failure 0 {string} string "失败"
// @Router /api/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	param := service.UserRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	loginStatus := svc.CheckLogin(&param)

	// 登录失败 - 账号/密码错误
	if loginStatus != true {
		response.ToResponse(gin.H{
			"code": errcode.Fail.Code(),
			"msg": "用户名不存在或者密码错误",
		})
		return
	}

	token, err := app.GenerateToken(param.UserName, "")
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"code": errcode.Success.Code(),
		"msg": errcode.Success.Msg(),
		"data": gin.H{ "token": token },
	})
}