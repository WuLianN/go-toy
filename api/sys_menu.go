package api

import (
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
	token := GetToken(c)
	err, tokenInfo := GetTokenInfo(token)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	svc := service.New(c.Request.Context())
	list := svc.GetMenuList(tokenInfo.UserId)
	if list != nil {
		response.ToResponse(gin.H{
			"code":    errcode.Success.Code(),
			"message": errcode.Success.Msg(),
			"type":    "success",
			"result":  list,
		})
	} else {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": errcode.Fail.Msg(),
		})
	}
}
