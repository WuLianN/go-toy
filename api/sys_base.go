package api

import (
	"github.com/WuLianN/go-toy/global"
	"github.com/WuLianN/go-toy/internal/service"
	"github.com/WuLianN/go-toy/pkg/app"
	"github.com/WuLianN/go-toy/pkg/convert"
	"github.com/WuLianN/go-toy/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

// @Summary 登录
// @Accept json
// @Produce json
// @Tags user
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {string} string "ok"
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
	if !loginStatus {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
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

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type":    "success",
		"result": gin.H{
			"token":      token,
			"user_name":  userInfo.UserName,
			"avatar":     userInfo.Avatar,
			"id":         userInfo.Id,
			"is_privacy": userInfo.IsPrivacy,
		},
	})
}

// @Summary 注册
// @Accept json
// @Produce json
// @Tags user
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {string} string "ok"
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
	bool, userId, err := svc.CheckRegister(&param)

	if bool {
		token, err := app.GenerateToken(userId, param.UserName)
		if err != nil {
			global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
			response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
			return
		}

		response.ToResponse(gin.H{
			"code":    errcode.Success.Code(),
			"message": "注册成功",
			"result": gin.H{
				"token":     token,
				"user_name": param.UserName,
				"id":        userId,
			},
		})
	} else if !bool && err == nil {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": "用户已注册",
		})
	} else {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": "注册失败",
		})
	}
}

// @Summary 推荐列表
// @Param user_id query int false "用户id"
// @Param page query int false "页码"
// @Param page_size query int false "页大小"
// @Tags 通用业务
// @Success 200 {string} string "ok"
// @Router /getRecommendList [get]
func (b *BaseApi) GetRecommendList(c *gin.Context) {
	response := app.NewResponse(c)

	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")
	userIdStr := c.Query("user_id")
	tagIdStr := c.Query("tag_id")

	if pageStr == "" {
		pageStr = "1"
	}
	if pageSizeStr == "" {
		pageSizeStr = "10"
	}
	if tagIdStr == "" {
		tagIdStr = "0"
	}

	svc := service.New(c.Request.Context())

	var userId uint32
	var isSelf uint8

	if userIdStr != "" {
		userId = convert.StrTo(userIdStr).MustUInt32()

		bool := svc.IsPrivacyUser(userId)
		if bool {
			response.ToResponse(gin.H{
				"code":    errcode.Success.Code(),
				"message": errcode.Success.Msg(),
				"type":    "success",
				"result":  make([]int, 0),
			})
			return
		}
	} else {
		token := GetToken(c)
		err, tokenInfo := GetTokenInfo(token)
		if err != nil {
			response.ToErrorResponse(err)
			return
		}
		userId = tokenInfo.UserId
		isSelf = 1
	}

	page := convert.StrTo(pageStr).MustInt()
	pageSize := convert.StrTo(pageSizeStr).MustInt()
	tagId := convert.StrTo(tagIdStr).MustUInt32()

	list, err := svc.GetRecommendList(userId, page, pageSize, tagId, isSelf)

	if err != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	if len(list) > 0 {
		response.ToResponse(gin.H{
			"code":    errcode.Success.Code(),
			"message": errcode.Success.Msg(),
			"type":    "success",
			"result":  list,
		})
	} else {
		response.ToResponse(gin.H{
			"code":    errcode.Success.Code(),
			"message": errcode.Success.Msg(),
			"type":    "success",
			"result":  make([]int, 0),
		})
	}
}
