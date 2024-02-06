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

	response := app.NewResponse(c)
	token := GetToken(c)
	err, tokenInfo := GetTokenInfo(token)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	userId := tokenInfo.UserId

	svc := service.New(c.Request.Context())
	tagList, err2 := svc.GetTagList(userId, ids)

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

func (t *TagApi) GetDraftTagList(c *gin.Context) {
	draftIdStr := c.Query("draft_id")
	tagIdStr := c.Query("tag_id")

	draftId := convert.StrTo(draftIdStr).MustUInt32()
	tagId := convert.StrTo(tagIdStr).MustUInt32()

	response := app.NewResponse(c)

	token := GetToken(c)
	err, tokenInfo := GetTokenInfo(token)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	userId := tokenInfo.UserId

	svc := service.New(c.Request.Context())
	tagList, err2 := svc.GetDraftTagList(userId, tagId, draftId)

	if err2 != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"result":  tagList,
	})
}

func (t *TagApi) GetMenuTagList(c *gin.Context) {
	menuIdStr := c.Query("menu_id")
	response := app.NewResponse(c)

	if menuIdStr == "" {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	menuId := convert.StrTo(menuIdStr).MustUInt32()

	svc := service.New(c.Request.Context())
	tagList, err := svc.GetMenuTagList(menuId)

	if err != nil {
		response.ToErrorResponse(errcode.Fail)
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

func (t *TagApi) UnbindTag2Menu(c *gin.Context) {
	requestBody := model.MenuTags{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &requestBody)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	if len(requestBody.Tags) == 0 {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails("tags", "tags不能为空"))
		return
	}

	svc := service.New(c.Request.Context())
	err2 := svc.UnbindTag2Menu(&requestBody)

	if err2 != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
	})
}

func (t *TagApi) BindTag2Draft(c *gin.Context) {
	requestBody := model.DraftTags{}
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
	err2 := svc.BindTag2Draft(&requestBody, tokenInfo.UserId)

	if err2 != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
	})
}

func (t *TagApi) UnbindTag2Draft(c *gin.Context) {
	requestBody := model.DraftTags{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &requestBody)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	if len(requestBody.Tags) == 0 {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails("tags", "tags不能为空"))
		return
	}

	svc := service.New(c.Request.Context())
	err2 := svc.UnbindTag2Draft(&requestBody)

	if err2 != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
	})
}
