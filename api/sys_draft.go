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

// @Summary 获取草稿 [已发布]
// @Param id body uint32 true "草稿id"
// @Tags 草稿
// @Success 200 {string} string "ok"
// @Router /getDraft [get]
func (d *DraftApi) GetDraft(c *gin.Context) {
	response := app.NewResponse(c)
	idStr := c.Query("id")
	if idStr == "" {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails("id不能为空"))
		return
	}

	id := convert.StrTo(idStr).MustUInt32()

	var userId uint32
	token := GetToken(c)

	_, tokenInfo := GetTokenInfo(token)

	if tokenInfo != nil {
		userId = tokenInfo.UserId
	}

	svc := service.New(c.Request.Context())
	result, err := svc.GetDraft(id, userId)

	if err != nil {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": "无这篇文章",
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

// @Summary 获取用户草稿 [token]
// @Param id body uint32 true "草稿id"
// @Tags 草稿
// @Success 200 {string} string "ok"
// @Router /getUserDraft [get]
func (d *DraftApi) GetUserDraft(c *gin.Context) {
	response := app.NewResponse(c)
	idStr := c.Query("id")
	if idStr == "" {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails("id不能为空"))
		return
	}

	id := convert.StrTo(idStr).MustUInt32()

	token := GetToken(c)
	err, tokenInfo := GetTokenInfo(token)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	userId := tokenInfo.UserId

	svc := service.New(c.Request.Context())
	result, err2 := svc.GetUserDraft(id, userId)

	if err2 != nil {
		response.ToResponse(gin.H{
			"code":    errcode.Fail.Code(),
			"message": "文章不存在",
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

// @Summary 创建草稿
// @Tags 草稿
// @Success 200 {object} model.ResponseResult{result=model.CreateDraftResponse} "ok"
// @Router /createDraft [get]
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

// @Summary 保存草稿
// @Param id body uint32 true "草稿id"
// @Param title body string false "标题"
// @Param content body string false "内容"
// @Tags 草稿
// @Accept json
// @Success 200 {object} model.ResponseResult "ok"
// @Router /saveDraft [post]
func (d *DraftApi) SaveDraft(c *gin.Context) {
	param := service.SaveRequest{}
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

	param.UserId = tokenInfo.UserId

	svc := service.New(c.Request.Context())
	err2 := svc.UpdateDraft(param)

	if err2 != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type":    "success",
	})
}

// @Summary 删除草稿
// @Param id body uint32 true "草稿id"
// @Tags 草稿
// @Accept json
// @Success 200 {string} string "ok"
// @Router /deleteDraft [post]
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

// @Summary 发布
// @Param id body uint32 true "草稿id"
// @Tags 草稿
// @Success 200 {string} string "ok"
// @Router /publishDraft [post]
func (d *DraftApi) PublishDraft(c *gin.Context) {
	param := service.PublishRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.PublishDraft(param)

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

// @Summary 获取草稿箱
// @Param page query uint32 false "页数"
// @Param page_size query uint32 false "页码"
// @Tags 草稿
// @Success 200 {string} string "ok"
// @Router /getDraftList [get]
func (d *DraftApi) GetDraftList(c *gin.Context) {
	param := service.DraftListRequest{}
	response := app.NewResponse(c)
	token := GetToken(c)
	err, tokenInfo := GetTokenInfo(token)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")
	if pageStr == "" {
		pageStr = "1"
	}
	param.Page = convert.StrTo(pageStr).MustInt()

	if pageSizeStr == "" {
		pageSizeStr = "10"
	}
	param.PageSize = convert.StrTo(pageSizeStr).MustInt()

	param.UserId = tokenInfo.UserId
	param.Status = convert.StrTo(c.Query("status")).MustUInt32()

	svc := service.New(c.Request.Context())
	list, err2 := svc.GetDraftList(&param)

	if err2 != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type":    "success",
		"result":  list,
	})
}

// @Summary 搜索文章
func (d *DraftApi) SearchDrafts(c *gin.Context) {
	response := app.NewResponse(c)

	keyword := c.Query("keyword")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")
	userIdStr := c.Query("user_id")

	if pageStr == "" {
		pageStr = "1"
	}
	if pageSizeStr == "" {
		pageSizeStr = "10"
	}

	svc := service.New(c.Request.Context())

	var userId uint32

	if userIdStr != "" {
		userId = convert.StrTo(userIdStr).MustUInt32()
		// 检查是否是私密账号
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
	}

	page := convert.StrTo(pageStr).MustInt()
	pageSize := convert.StrTo(pageSizeStr).MustInt()

	list, err := svc.SearchDrafts(userId, keyword, page, pageSize)

	if err != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type":    "success",
		"result":  list,
	})
}
