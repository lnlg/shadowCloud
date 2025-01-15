package logger

import (
	"os"
	"shadowCloud/internal/config"
	"shadowCloud/internal/global"
	"shadowCloud/internal/tool"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	options []zap.Option
	conf    config.Logger
)

// 初始化日志
func New() (logger *zap.Logger, err error) {
	// 获取配置文件
	conf = global.Config.Logger

	// 创建日志存放目录
	rootDir := tool.GetRootDir()
	logDir := rootDir + conf.FilePath
	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return
	}

	loggerConf := genConfig()

	// 设置时间格式
	loggerConf.EncoderConfig = genEncodeConfig()

	// 设置日志写入器
	writer, err := genWriteSyncer("")
	if err != nil {
		return nil, err
	}

	// 创建日志核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(loggerConf.EncoderConfig),
		writer,
		loggerConf.Level,
	)

	// 添加调用位置
	options = append(options, zap.AddCaller())

	// 创建日志
	logger = zap.New(core, options...)

	return logger, nil
}

// 生成WriteSyncer 日志写入器
func genWriteSyncer(name string) (writeSyncer zapcore.WriteSyncer, err error) {
	// 创建日志存放目录
	rootDir := tool.GetRootDir()
	logDir := rootDir + conf.FilePath
	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return
	}
	filename := ""
	if name == "" {
		filename = logDir + conf.FileName
	} else {
		filename = logDir + name + ".log"
	}
	// 创建日志写入器
	lumberJack := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    conf.MaxSize, // megabytes
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge, //days
	}

	// 设置日志写入器
	writeSyncer = zapcore.AddSync(lumberJack)

	return
}

// 生成配置
func genConfig() (config zap.Config) {
	// 生成配置
	config = zap.NewProductionConfig()

	// 生成编码配置
	config.EncoderConfig = genEncodeConfig()

	// 设置日志级别
	config.Level = zap.NewAtomicLevelAt(getLevel())

	return

}

// 生成编码配置
func genEncodeConfig() (c zapcore.EncoderConfig) {
	// 生成编码配置
	c = zap.NewProductionEncoderConfig()

	// 设置时间格式
	c.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05"))
	}

	// 设置日志级别
	c.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(strings.ToUpper(l.String()))
	}

	c.TimeKey = "time" // 设置时间字段名

	return
}

// 配置文件的level转换为zapcore的level
func getLevel() (level zapcore.Level) {
	switch conf.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	return
}
