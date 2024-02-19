package api

import (
	"github.com/WuLianN/go-toy/global"
	dao "github.com/WuLianN/go-toy/internal/dao"
	"github.com/WuLianN/go-toy/internal/service"
	"github.com/WuLianN/go-toy/pkg/app"
	"github.com/WuLianN/go-toy/pkg/convert"
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
	var userId uint32
	userIdStr := c.Query("id")
	response := app.NewResponse(c)

	if userIdStr == "" {
		var token string
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("Authorization")
		}

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
		userId = tokenInfo.UserId
	} else {
		userId = convert.StrTo(userIdStr).MustUInt32()
	}

	svc := service.New(c.Request.Context())

	userInfo, err2 := svc.GetUserInfo(userId)

	if err2 != nil {
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
		"result":  userInfo,
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

	if !isSystemUser {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	// 验证旧密码是否正确
	isRightOldPasword := service.ComparePassword(param.OldPassword, userInfo.Password)

	if !isRightOldPasword {
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

func (u *UserApi) UpdateUserInfo(c *gin.Context) {
	requestBody := service.UserInfoRequest{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &requestBody)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	if requestBody.Id == 0 {
		token := GetToken(c)
		err, tokenInfo := GetTokenInfo(token)
		if err != nil {
			response.ToErrorResponse(err)
			return
		}
		requestBody.Id = tokenInfo.UserId
	}

	svc := service.New(c.Request.Context())
	userInfo, err := svc.UpdateUserInfo(&requestBody)

	if err != nil {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": errcode.Fail.Msg(),
		})
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type":    "success",
		"result": gin.H{
			"id":        requestBody.Id,
			"user_name": userInfo.UserName,
			"avatar":    userInfo.Avatar,
		},
	})
}

func (u *UserApi) BingUser(c *gin.Context) {
	requestBody := service.UserRequest{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &requestBody)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	var userId uint32
	token := GetToken(c)
	err, tokenInfo := GetTokenInfo(token)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	userId = tokenInfo.UserId

	svc := service.New(c.Request.Context())
	err2 := svc.BindUser(userId, &requestBody)

	if err2 != nil {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": err2.Error(),
		})
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type":    "success",
	})
}

func (u *UserApi) UnbindUser(c *gin.Context) {
	requestBody := service.UserIdRequest{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &requestBody)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	var userId uint32
	token := GetToken(c)
	err, tokenInfo := GetTokenInfo(token)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	userId = tokenInfo.UserId

	svc := service.New(c.Request.Context())
	err2 := svc.UnbindUser(userId, requestBody.Id)

	if err2 != nil {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": err2.Error(),
		})
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type":    "success",
	})
}

func (u *UserApi) GetBindedUserList(c *gin.Context) {
	userIdStr := c.Query("user_id")

	response := app.NewResponse(c)

	if userIdStr == "" {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": errcode.Fail.Msg(),
		})
		return
	}

	userId := convert.StrTo(userIdStr).MustUInt32()

	svc := service.New(c.Request.Context())
	list, err := svc.GetBindedUserList(userId)

	if err != nil {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": errcode.Fail.Msg(),
		})
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type":    "success",
		"result":  list,
	})
}

func (u *UserApi) ChangeAccount(c *gin.Context) {
	requestBody := service.UserIdRequest{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &requestBody)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	var userId uint32
	token := GetToken(c)
	err, tokenInfo := GetTokenInfo(token)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	userId = tokenInfo.UserId

	svc := service.New(c.Request.Context())
	status := svc.CheckBindedUser(userId, requestBody.Id)

	if !status {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": "该用户未绑定",
		})
		return
	}

	changeStatus, userInfo := svc.ChangeAccount(requestBody.Id)

	if !changeStatus {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": "该用户不存在",
		})
		return
	}

	newUserToken, err2 := app.GenerateToken(userInfo.Id, userInfo.UserName)

	if err2 != nil {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": "切换失败",
		})
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type":    "success",
		"result": gin.H{
			"token":     newUserToken,
			"user_name": userInfo.UserName,
			"avatar":    userInfo.Avatar,
			"id":        userInfo.Id,
		},
	})
}
