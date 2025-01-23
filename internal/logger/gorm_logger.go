package logger

import (
	"context"
	"fmt"
	"shadowCloud/internal/global"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

// 自定义Gorm日志
type GormLogger struct {
	Logger logger.LogLevel
}

// LogMode 实现 gorm.Logger 接口的方法
func (g *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	g.Logger = level
	return g
}

// Warn 实现 gorm.Logger 接口的方法
func (g *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if g.Logger == logger.Warn {
		global.Logger.Warn(msg, zap.Any("data", data))
	}
}

// Info 实现 gorm.Logger 接口的方法
func (g *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if g.Logger == logger.Info {
		global.Logger.Info(msg, zap.Any("data", data))
	}
}

// Error 实现 gorm.Logger 接口的方法
func (g *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if g.Logger == logger.Error {
		global.Logger.Error(msg, zap.Any("data", data))
	}
}

// Trace 实现 gorm.Logger 接口的方法
func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	//根据日志级别决定是否记录 SQL 日志
	if g.Logger == logger.Info {
		//计算执行时间
		elapsed := time.Since(begin)
		//获取 SQL 和影响的行数
		sql, rows := fc()
		// println("[", elapsed.Nanoseconds()/1e6, "ms]", sql, "影响行数:", rows)
		global.Logger.Info(sql, zap.Any("data", map[string]interface{}{
			"time":  fmt.Sprintf("%d ms", elapsed.Nanoseconds()/1e6),
			"rows":  rows,
			"error": err,
		}))
	}
}
