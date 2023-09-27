package api

import (
	"github.com/gin-gonic/gin"
	"github.com/WuLianN/go-toy/internal/service"
	"github.com/WuLianN/go-toy/pkg/app"
	"github.com/WuLianN/go-toy/global"
	"github.com/WuLianN/go-toy/pkg/errcode"
	"github.com/WuLianN/go-toy/internal/model"
)

type BaseApi struct {}

// @Summary 测试Ping
// @Produce json
// @Success 0 {string} string "ok"
// @Router /api/ping [get]
func (b *BaseApi) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// @Summary 登录
// @Accept json
// @Produce json
// @Tags user
// @Param username body string true "用户名" 
// @Param password body string true "密码" 
// @Success 0 {string} string "ok"
// @Failure 1 {string} string "fail"
// @Router /login [post]
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
	loginStatus, userInfo := svc.CheckLogin(&param)

	// 登录失败 - 账号/密码错误
	if loginStatus != true {
		response.ToResponse(gin.H{
			"code": errcode.Fail.Code(),
			"message": "用户名不存在或者密码错误",
		})
		return
	}

	token, err := app.GenerateToken(userInfo.Id, userInfo.UserName)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	// 获取角色权限
	var roles []model.Role
	if userInfo.Id != 0 {
		roles = svc.GetRoleList(userInfo.Id)
	}
	
	response.ToResponse(gin.H{
		"code": errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type": "success",
		"result": gin.H {
			"desc": "manager",
			"token": token,
			"roles": roles,
			"username": userInfo.UserName,
			"realName": "Vben Admin",
			"userId": userInfo.Id,
		},
	})
}

// @Summary 注册
// @Accept json
// @Produce json
// @Tags user
// @Param username body string true "用户名" 
// @Param password body string true "密码" 
// @Success 0 {string} string "ok"
// @Failure 1 {string} string "fail"
// @Router /register [post]
func (b *BaseApi) Register(c *gin.Context) {
	param := service.UserRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	bool, err := svc.CheckRegister(&param)

	if bool == true {
		response.ToResponse(gin.H{
			"code": errcode.Success.Code(),
			"message": "注册成功",
		})
	} else if bool == false && err == nil {
		response.ToResponse(gin.H{
			"code": errcode.Fail.Code(),
			"message": "用户已注册",
		})
	} else {
		response.ToResponse(gin.H{
			"code": errcode.Fail.Code(),
			"message": "注册失败",
		})
	}
}