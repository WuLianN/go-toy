package api

import (
	"strconv"

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
