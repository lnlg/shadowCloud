package global

import (
	"shadowCloud/internal/config"
	"shadowCloud/internal/crontab"
	"shadowCloud/internal/event"

	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config    *config.Config
	Logger    *zap.Logger
	Db        *gorm.DB
	Rdb       *redis.Client
	Validator *validator.Validate // validator 验证器
	Crontab   *crontab.Crontab    // crontab 定时任务
	Event     *event.Dispatcher   // event 事件调度器
)
