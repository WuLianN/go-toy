package global

import (
	"github.com/WuLianN/go-toy/pkg/logger"
	"github.com/WuLianN/go-toy/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	JWTSetting      *setting.JWTSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
