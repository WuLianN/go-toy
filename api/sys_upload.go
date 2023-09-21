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

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

// @Summary 上传文件
// @Accept application/form-data
// @Produce json
// @Tags 基建
// @Param type body string true "类型" 
// @Param file body string true "文件" 
// @Router /upload/file [post]
func (u Upload) UploadFile(c *gin.Context) {
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
		"file_access_url": fileInfo.AccessUrl,
	})
}