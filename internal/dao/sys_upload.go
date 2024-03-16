package dao

import (
	"time"

	"github.com/WuLianN/go-toy/internal/model"
)

func (d *Dao) CreateUploadRecord(filename, accessUrl string) error {
	loc, err := time.LoadLocation("Asia/Shanghai")

	if err != nil {
		loc = time.FixedZone("CST", 8*3600) // 替换上海时间
	}

	uploadRecord := model.UploadRecord{
		Name:      filename,
		CreatedAt: time.Now().In(loc).Format(time.DateTime),
		AccessUrl: accessUrl,
	}

	err1 := d.engine.Table("upload_files").Create(&uploadRecord).Error

	return err1
}

func (d *Dao) DeleteUploadRecord(idList []uint32) ([]model.UploadRecord, error) {
	var uploadRecordList []model.UploadRecord
	var err error

	// 先查询要删除的记录
	err = d.engine.Table("upload_files").Where("id in ?", idList).Find(&uploadRecordList).Error
	if err != nil {
		return nil, err
	}

	// 然后执行删除操作
	err = d.engine.Table("upload_files").Where("id in ?", idList).Delete(&uploadRecordList).Error

	return uploadRecordList, err
}

func (d *Dao) QueryUploadRecordList(page int, pageSize int, order string, keyword string) ([]model.UploadRecord, error) {
	var uploadRecordList []model.UploadRecord
	var err error

	if order == "asc" {
		order = "asc"
	} else {
		order = "desc"
	}

	orderStr := "created_at " + order

	if keyword != "" {
		err = d.engine.Table("upload_files").Where("name like ?", "%"+keyword+"%").Limit(pageSize).Offset((page - 1) * pageSize).Order(orderStr).Find(&uploadRecordList).Error
	} else {
		err = d.engine.Table("upload_files").Limit(pageSize).Offset((page - 1) * pageSize).Order(orderStr).Find(&uploadRecordList).Error
	}

	return uploadRecordList, err
}
