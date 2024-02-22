package db

import (
	"fmt"

	"github.com/WuLianN/go-toy/global"
	otgorm "github.com/WuLianN/go-toy/pkg/opentracing-gorm"
	"github.com/WuLianN/go-toy/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {

	}

	otgorm.AddGormCallbacks(db)

	return db, nil
}
