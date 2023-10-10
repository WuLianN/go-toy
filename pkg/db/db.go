package db

import (
	"fmt"

	"github.com/WuLianN/go-toy/global"
	"github.com/WuLianN/go-toy/pkg/setting"
	otgorm "github.com/WuLianN/go-toy/pkg/opentracing-gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	redis "github.com/redis/go-redis/v9"
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
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		
	})


	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		
	}

	otgorm.AddGormCallbacks(db)

	return db, nil
}

func NewRedisDBEngine(redisDBSetting *setting.RedisDBSettingS) (*redis.Client) {
	rdb := redis.NewClient(&redis.Options{
		Addr:	  redisDBSetting.Addr,
		Password: redisDBSetting.Password,
		DB:		  redisDBSetting.DB,
	})

	return rdb
}
