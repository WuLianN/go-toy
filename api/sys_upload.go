package api

import (
	"github.com/WuLianN/go-toy/global"
	"github.com/WuLianN/go-toy/internal/service"
	"github.com/WuLianN/go-toy/pkg/app"
	"github.com/WuLianN/go-toy/pkg/convert"
	"github.com/WuLianN/go-toy/pkg/errcode"
	"github.com/WuLianN/go-toy/pkg/upload"
	"github.com/gin-gonic/gin"
)

type UploadApi struct{}

func NewUpload() UploadApi {
	return UploadApi{}
}

// @Summary 上传文件
// @Accept application/form-data
// @Produce json
// @Tags 基建
// @Param type body string true "类型"
// @Param file body string true "文件"
// @Router /upload/file [post]
func (u UploadApi) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()

	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)

	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ERROR_UPLOAD_FILE_FAIL.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type":    "success",
		"result": gin.H{
			"file_access_url": fileInfo.AccessUrl,
		},
	})
}

func (u UploadApi) GetUploadFileList(c *gin.Context) {
	param := service.FileListRequest{}
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
	param.Order = c.Query("order")
	param.Keyword = c.Query("keyword")

	svc := service.New(c.Request.Context())

	list, err2 := svc.GetUploadFileList(&param)

	if err2 != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"result":  list,
	})
}

func (u UploadApi) DeleteUploadFile(c *gin.Context) {
	param := service.DeleteFileRequest{}

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
	err2 := svc.DeleteUploadFile(param.UserId, param.Ids)

	if err2 != nil {
		response.ToErrorResponse(errcode.Fail)
		return
	}

	response.ToResponse(gin.H{
		"code":    errcode.Success.Code(),
		"message": errcode.Success.Msg(),
	})
}
