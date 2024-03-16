package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/WuLianN/go-toy/global"
	"github.com/WuLianN/go-toy/internal/model"
	"github.com/WuLianN/go-toy/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

type FileListRequest struct {
	UserId   uint32 `json:"user_id"`
	Page     int    `json:"page" default:"1"`
	PageSize int    `json:"page_size" default:"10"`
	Order    string `json:"order" default:"desc"`
	Keyword  string `json:"keyword"`
}

type DeleteFileRequest struct {
	UserId uint32   `json:"user_id"`
	Ids    []uint32 `json:"ids" binding:"required"`
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName

	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported.")
	}

	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}

	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}

	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}

	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName

	// 存储fileName至mysql
	svc.dao.CreateUploadRecord(fileName, accessUrl)

	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}

func (svc *Service) GetUploadFileList(req *FileListRequest) ([]model.UploadRecord, error) {
	var records []model.UploadRecord
	var err error

	// 检查user_id
	bool := svc.dao.IsAdmin(req.UserId)

	if !bool {
		return nil, errors.New("user_id is not valid.")
	}

	if records, err = svc.dao.QueryUploadRecordList(req.Page, req.PageSize, req.Order, req.Keyword); err != nil {
		return nil, err
	}
	return records, nil
}

func (svc *Service) DeleteUploadFile(userId uint32, Ids []uint32) error {
	var delList []model.UploadRecord
	var err error

	// 检查user_id
	bool := svc.dao.IsAdmin(userId)

	if !bool {
		return errors.New("user_id is not valid.")
	}

	// 批量删除
	if delList, err = svc.dao.DeleteUploadRecord(Ids); err != nil {
		return err
	}

	// 物理删除
	if len(delList) > 0 {
		for _, record := range delList {
			if err = deleteImage(record.Name); err != nil {
				return err
			}
		}
	}

	return nil
}

func deleteImage(name string) error {
	// 获取图片在服务器上的完整路径
	fullPath := filepath.Join(global.AppSetting.UploadSavePath, name)

	// 删除图片
	err := os.Remove(fullPath)
	if err != nil {
		return fmt.Errorf("failed to delete image: %w", err)
	}

	return nil
}
