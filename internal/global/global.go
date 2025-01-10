package global

import (
	"shadowCloud/internal/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Logger *zap.Logger
	Db     *gorm.DB
)
