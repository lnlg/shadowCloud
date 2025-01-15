package bootstrap

import (
	appEVE "shadowCloud/app/event"
	"shadowCloud/app/task"
	"shadowCloud/internal/config"
	"shadowCloud/internal/crontab"
	"shadowCloud/internal/event"
	"shadowCloud/internal/global"
	"shadowCloud/internal/logger"
	"shadowCloud/internal/mysql"
	"shadowCloud/internal/redis"
	"shadowCloud/internal/validator"
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
	// 初始化redis
	global.Rdb = redis.New()
	// 初始化验证器
	global.Validator = validator.InitValidator()

	// 初始化crontab
	global.Crontab = crontab.Init()
	// 添加任务
	global.Crontab.AddTask(task.Tasks()...)
	// 启动crontab
	global.Crontab.Start()

	// 初始化事件调度器
	global.Event = event.New()
	// 注册事件
	appEVE.RegisterAppEvent()
}
