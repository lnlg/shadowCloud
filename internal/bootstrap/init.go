package bootstrap

import (
	"shadowCloud/internal/config"
	"shadowCloud/internal/global"
)

func init() {
	// 初始化配置文件
	global.Config = config.GetConfig()
}
