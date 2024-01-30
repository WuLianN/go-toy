package api

import (
	"github.com/WuLianN/go-toy/global"
	"github.com/WuLianN/go-toy/internal/service"
	"github.com/WuLianN/go-toy/pkg/app"
	"github.com/WuLianN/go-toy/pkg/convert"
	"github.com/WuLianN/go-toy/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type DraftApi struct{}

func (d *DraftApi) GetDraft(c *gin.Context) {
	response := app.NewResponse(c)
	idStr := c.Query("id")
	if idStr == "" {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails("id不能为空"))
		return
	}

	id := convert.StrTo(idStr).MustUInt32()

	svc := service.New(c.Request.Context())
	result, err := svc.GetDraft(id)

	if err != nil {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": "无这篇文章",
			"type":    "info",
		})
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type":    "success",
		"result":  result,
	})
}

func (d *DraftApi) CreateDraft(c *gin.Context) {
	response := app.NewResponse(c)
	var userId uint32
	token := GetToken(c)
	err, tokenInfo := GetTokenInfo(token)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	userId = tokenInfo.UserId

	svc := service.New(c.Request.Context())
	draftId := svc.CreateDraft(userId)

	if draftId == 0 {
		response.ToErrorResponse(errcode.ServerError)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type":    "success",
		"result": gin.H{
			"draft_id": draftId,
		},
	})
}

func (d *DraftApi) SaveDraft(c *gin.Context) {
	param := service.SaveRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateDraft(param)

	if err != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type":    "success",
	})
}

func (d *DraftApi) DeleteDraft(c *gin.Context) {
	param := service.DeleteRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteDraft(param)

	if err != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type":    "success",
	})
}
