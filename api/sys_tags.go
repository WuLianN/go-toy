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

type TagApi struct{}

func (t *TagApi) GetTagList(c *gin.Context) {
	ids := c.Query("ids")
	menuIdStr := c.Query("menu_id")

	menuId := convert.StrTo(menuIdStr).MustUInt32()

	response := app.NewResponse(c)
	token := GetToken(c)
	err, tokenInfo := GetTokenInfo(token)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	userId := tokenInfo.UserId

	svc := service.New(c.Request.Context())
	tagList, err2 := svc.GetTagList(userId, ids, menuId)

	if err2 != nil {
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"result":  tagList,
	})
}

func (t *TagApi) CreateTag(c *gin.Context) {
	requestBody := service.CreateTagRequest{}
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
	requestBody.UserId = tokenInfo.UserId

	svc := service.New(c.Request.Context())

	tagList, err2 := svc.QueryTag(&requestBody)

	if len(tagList) > 0 && err2 == nil {
		response.ToResponse(gin.H{
			"code":    errcode.Warning.Code(),
			"message": "标签已存在",
		})
		return
	}

	tagId, err3 := svc.CreateTag(&requestBody)

	if err3 != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"result": gin.H{
			"id": tagId,
		},
	})
}

func (t *TagApi) DeleteTag(c *gin.Context) {
	requestBody := service.DeleteTagRequest{}
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
	requestBody.UserId = tokenInfo.UserId

	svc := service.New(c.Request.Context())

	err2 := svc.DeleteTag(&requestBody)

	if err2 != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
	})
}

func (t *TagApi) UpdateTag(c *gin.Context) {
	requestBody := service.UpdateTagRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &requestBody)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())

	err2 := svc.UpdateTag(&requestBody)

	if err2 != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
	})
}

func (t *TagApi) BindTag2Menu(c *gin.Context) {
	requestBody := model.MenuTags{}
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

	if len(requestBody.Tags) == 0 {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails("tags", "tags不能为空"))
		return
	}

	svc := service.New(c.Request.Context())
	err2 := svc.BindTag2Menu(&requestBody, tokenInfo.UserId)

	if err2 != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
	})
}
