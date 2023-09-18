package api

import (
	"github.com/gin-gonic/gin"
	"github.com/WuLianN/go-toy/internal/service"
	"github.com/WuLianN/go-toy/global"
	"github.com/WuLianN/go-toy/pkg/app"
	"github.com/WuLianN/go-toy/pkg/errcode"
	dao "github.com/WuLianN/go-toy/internal/dao"
	"github.com/WuLianN/go-toy/internal/model"
)

type UserApi struct {}

// @Summary 获取用户信息
// @Accept json
// @Produce json
// @Tags user
// @Success 0 {string} string "ok"
// @Failure 1 {string} string "fail"
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
			"code": errcode.Fail.Code(),
			"msg": "无token",
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
			"code": errcode.Fail.Code(),
			"msg": "用户ID错误",
		})
		return
	}

	svc := service.New(c.Request.Context())

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
			"roles": roles,
			"username": userInfo.UserName,
			"realName": "Vben Admin",
			"userId": userInfo.Id,
		},
	})
}

// @Summary 修改密码
// @Param oldPassword body string true "旧密码"
// @Param newPassword body string true "新密码"
// @Accept json
// @Produce json
// @Tags user
// @Success 0 {string} string "ok"
// @Failure 1 {string} string "fail"
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
			"code": errcode.Fail.Code(),
			"message": "当前密码错误",
		})
		return 
	}

	svc := service.New(c.Request.Context())
	isSuccessful := svc.ChangePassword(userInfo.Id, param.NewPassword)

	if isSuccessful {
		response.ToResponse(gin.H{
			"code": errcode.Success.Code(),
			"message": "密码已更换",
		})
		return 
	} else {
		response.ToResponse(gin.H{
			"code": errcode.Fail.Code(),
			"message": "密码更换失败!",
		})
		return 
	}
}