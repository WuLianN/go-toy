package api

import (
	"github.com/WuLianN/go-toy/global"
	"github.com/WuLianN/go-toy/internal/model"
	"github.com/WuLianN/go-toy/internal/service"
	"github.com/WuLianN/go-toy/pkg/app"
	"github.com/WuLianN/go-toy/pkg/convert"
	"github.com/WuLianN/go-toy/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type MenuApi struct{}

// @Summary 获取角色菜单
// @Accept json
// @Produce json
// @Tags 菜单
// @Success 200 {string} string "ok"
// @Router /getMenuList [get]
func (m *MenuApi) GetRoleMenu(c *gin.Context) {
	response := app.NewResponse(c)
	userIdStr := c.Query("user_id")
	var userId uint32

	if userIdStr != "" {
		userId = convert.StrTo(userIdStr).MustUInt32()
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

// @Summary 添加子菜单
// @Param parent_id body number true "父级菜单id"
// @Param name body string true "名称"
// @Accept json
// @Produce json
// @Tags 菜单
// @Success 200 {string} string "ok"
// @Router /addMenuItem [post]
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

// @Summary 删除菜单
// @Param id body number true "菜单id"
// @Accept json
// @Produce json
// @Tags 菜单
// @Success 200 {string} string "ok"
// @Router /deleteMenuItem [post]
func (m *MenuApi) DeleteMenuItem(c *gin.Context) {
	requestBody := model.DeleteMenuItem{}
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

	err2 := svc.DeleteMenuItem(requestBody, userId)

	if err2 != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
	})
}

// @Summary 更新菜单
// @Param id body number true "菜单id"
// @Param name body string false "菜单名称"
// @Param icon body string false "图标"
// @Accept json
// @Produce json
// @Tags 菜单
// @Success 200 {string} string "ok"
// @Router /updateMenuItem [post]
func (m *MenuApi) UpdateMenuItem(c *gin.Context) {
	requestBody := model.UpdateMenuItem{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &requestBody)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateMenuItem(&requestBody)

	if err != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
	})
}
