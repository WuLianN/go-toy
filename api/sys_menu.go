package api

import (
	"strconv"

	"github.com/WuLianN/go-toy/global"
	"github.com/WuLianN/go-toy/internal/model"
	"github.com/WuLianN/go-toy/internal/service"
	"github.com/WuLianN/go-toy/pkg/app"
	"github.com/WuLianN/go-toy/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type MenuApi struct{}

// @Summary 获取角色菜单
// @Accept json
// @Produce json
// @Tags menu
// @Success 200 {string} string "ok"
// @Router /getMenuList [get]
func (m *MenuApi) GetRoleMenu(c *gin.Context) {
	response := app.NewResponse(c)
	userIdStr := c.Query("user_id")
	var userId uint32

	if userIdStr != "" {
		userIdUint64, _ := strconv.ParseUint(userIdStr, 10, 32)
		userId = uint32(userIdUint64)
	} else {
		token := GetToken(c)
		err, tokenInfo := GetTokenInfo(token)
		if err != nil {
			response.ToErrorResponse(err)
			return
		}
		userId = tokenInfo.UserId
	}

	svc := service.New(c.Request.Context())
	list := svc.GetMenuList(userId)

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type":    "success",
		"result":  list,
	})
}

func (m *MenuApi) AddMenuItem(c *gin.Context) {
	requestBody := model.AddMenuItem{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &requestBody)
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

	userId := tokenInfo.UserId

	svc := service.New(c.Request.Context())

	addMenuItem, err2 := svc.AddMenuItem(requestBody, userId)

	if err2 != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"result":  addMenuItem,
	})
}
