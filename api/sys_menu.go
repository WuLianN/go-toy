package api

import (
	"github.com/gin-gonic/gin"
	"github.com/WuLianN/go-toy/internal/service"
	"github.com/WuLianN/go-toy/pkg/app"
	"github.com/WuLianN/go-toy/pkg/errcode"
)

type MenuApi struct {}

func (m *MenuApi) GetRoleMenu(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	list := svc.GetMenuList()
	if list != nil {
		response.ToResponse(gin.H{
			"code": errcode.Success.Code(),
			"message": errcode.Success.Msg(),
			"type": "success",
			"result": list,
		})
	} else {
		response.ToResponse(gin.H{
			"code": errcode.Fail.Code(),
			"message": errcode.Fail.Msg(),
		})
	}
}