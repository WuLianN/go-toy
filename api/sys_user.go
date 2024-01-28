package api

import (
	"github.com/WuLianN/go-toy/global"
	dao "github.com/WuLianN/go-toy/internal/dao"
	"github.com/WuLianN/go-toy/internal/service"
	"github.com/WuLianN/go-toy/pkg/app"
	"github.com/WuLianN/go-toy/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// @Summary 获取用户信息
// @Accept json
// @Produce json
// @Tags user
// @Success 200 {string} string "ok"
// @Router /getUserInfo [get]
func (u *UserApi) GetUserInfo(c *gin.Context) {
	var token string
	if s, exist := c.GetQuery("token"); exist {
		token = s
	} else {
		token = c.GetHeader("Authorization")
	}

	response := app.NewResponse(c)

	if token == "" {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": "无token",
		})
		return
	}

	err, tokenInfo := GetTokenInfo(token)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	loginStatus, userInfo := dao.New(global.DBEngine).IsSystemUser("", tokenInfo.UserId)

	if loginStatus != true {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": "用户ID错误",
		})
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type":    "success",
		"result": gin.H{
			"username": userInfo.UserName,
			"userId":   userInfo.Id,
		},
	})
}

// @Summary 修改密码
// @Param oldPassword body string true "旧密码"
// @Param newPassword body string true "新密码"
// @Accept json
// @Produce json
// @Tags user
// @Success 200 {string} string "ok"
// @Router /changePassword [post]
func (u *UserApi) ChangePassword(c *gin.Context) {
	param := service.ChangePasswordRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	token := GetToken(c)

	err, tokenInfo := GetTokenInfo(token)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	isSystemUser, userInfo := dao.New(global.DBEngine).IsSystemUser("", tokenInfo.UserId)

	if isSystemUser == false {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	// 验证旧密码是否正确
	isRightOldPasword := service.ComparePassword(param.OldPassword, userInfo.Password)

	if isRightOldPasword == false {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": "当前密码错误",
		})
		return
	}

	svc := service.New(c.Request.Context())
	isSuccessful := svc.ChangePassword(userInfo.Id, param.NewPassword)

	if isSuccessful {
		response.ToResponse(gin.H{
			"code":    errcode.Success.Code(),
			"message": "密码已更换",
		})
		return
	} else {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": "密码更换失败!",
		})
		return
	}
}
