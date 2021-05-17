package global

import (
	"github.com/SuperArnold/GO_Blog/pkg/logger"
	"github.com/SuperArnold/GO_Blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	Logger          *logger.Logger
)
