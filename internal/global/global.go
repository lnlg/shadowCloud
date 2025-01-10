package global

import (
	"shadowCloud/internal/config"

	"go.uber.org/zap"
)

var (
	Config *config.Config
	Logger *zap.Logger
)
