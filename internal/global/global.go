package global

import (
	"shadowCloud/internal/config"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Logger *zap.Logger
	Db     *gorm.DB
	Rdb    *redis.Client
)
