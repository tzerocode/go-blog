package global

import (
	"first-gin/pkg/logger"
	"first-gin/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
