package bootstrap

import (
	"shadowCloud/internal/config"
	"shadowCloud/internal/global"
	"shadowCloud/internal/logger"
	"shadowCloud/internal/mysql"
)

func init() {
	var err error
	// 初始化配置文件
	global.Config = config.GetConfig()
	// 初始化日志
	if global.Logger, err = logger.New(); err != nil {
		panic(err)
	}
	// 初始化数据库
	global.Db = mysql.New()
}
