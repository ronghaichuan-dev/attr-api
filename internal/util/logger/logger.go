package logger

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// customLevelEncoder 自定义Level编码器，使用[]包起来
func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + strings.ToUpper(level.String()) + "]")
}

// customColorLevelEncoder 自定义彩色Level编码器，使用[]包起来
func customColorLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	// 使用zapcore的彩色级别编码器，然后用[]包起来
	var buf []byte
	switch level {
	case zapcore.DebugLevel:
		buf = append(buf, "\x1b[90m["...)
	case zapcore.InfoLevel:
		buf = append(buf, "\x1b[36m["...)
	case zapcore.WarnLevel:
		buf = append(buf, "\x1b[33m["...)
	case zapcore.ErrorLevel:
		buf = append(buf, "\x1b[31m["...)
	case zapcore.DPanicLevel:
		buf = append(buf, "\x1b[35m["...)
	case zapcore.PanicLevel:
		buf = append(buf, "\x1b[35m["...)
	case zapcore.FatalLevel:
		buf = append(buf, "\x1b[35m["...)
	}
	buf = append(buf, strings.ToUpper(level.String())...)
	buf = append(buf, "]\x1b[0m"...)
	enc.AppendString(string(buf))
}

var (
	// Logger 全局日志实例
	Logger *zap.Logger
	// SugarLogger 全局Sugar日志实例
	SugarLogger *zap.SugaredLogger
	// once 确保日志只初始化一次
	once sync.Once
)

// getLogFile 获取按日期命名的日志文件
func getLogFile() *os.File {
	// 日志目录
	logDir := "./tmp/logs"
	// 确保日志目录存在
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(err)
	}
	// 按日期命名日志文件
	logFile := filepath.Join(logDir, time.Now().Format("2006-01-02")+".log")
	// 打开或创建日志文件
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return file
}

// InitLogger 初始化日志
func InitLogger() {
	once.Do(func() {
		consoleEncoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    customColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		fileEncoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    customColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		logFile := getLogFile()

		core := zapcore.NewTee(
			zapcore.NewCore(
				zapcore.NewConsoleEncoder(consoleEncoderConfig),
				zapcore.AddSync(os.Stdout),
				zapcore.InfoLevel,
			),
			zapcore.NewCore(
				zapcore.NewConsoleEncoder(fileEncoderConfig),
				zapcore.AddSync(logFile),
				zapcore.InfoLevel,
			),
		)

		Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
		SugarLogger = Logger.Sugar()
	})
}

// GetLogger 获取Logger实例
func GetLogger() *zap.Logger {
	if Logger == nil {
		InitLogger()
	}
	return Logger
}

// GetSugarLogger 获取SugarLogger实例
func GetSugarLogger() *zap.SugaredLogger {
	if SugarLogger == nil {
		InitLogger()
	}
	return SugarLogger
}

// Debug 输出Debug级别日志
func Debug(msg string, fields ...zapcore.Field) {
	GetLogger().Debug(msg, fields...)
}

// Info 输出Info级别日志
func Info(msg string, fields ...zapcore.Field) {
	GetLogger().Info(msg, fields...)
}

// Warn 输出Warn级别日志
func Warn(msg string, fields ...zapcore.Field) {
	GetLogger().Warn(msg, fields...)
}

// Error 输出Error级别日志
func Error(msg string, fields ...zapcore.Field) {
	GetLogger().Error(msg, fields...)
}

// Fatal 输出Fatal级别日志
func Fatal(msg string, fields ...zapcore.Field) {
	GetLogger().Fatal(msg, fields...)
}

// Debugf 输出格式化Debug级别日志
func Debugf(template string, args ...interface{}) {
	GetSugarLogger().Debugf(template, args...)
}

// Infof 输出格式化Info级别日志
func Infof(template string, args ...interface{}) {
	GetSugarLogger().Infof(template, args...)
}

// Warnf 输出格式化Warn级别日志
func Warnf(template string, args ...interface{}) {
	GetSugarLogger().Warnf(template, args...)
}

// Errorf 输出格式化Error级别日志
func Errorf(template string, args ...interface{}) {
	GetSugarLogger().Errorf(template, args...)
}

// Fatalf 输出格式化Fatal级别日志
func Fatalf(template string, args ...interface{}) {
	GetSugarLogger().Fatalf(template, args...)
}

// GinLogger returns a gin middleware for logging
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		if query != "" {
			path = path + "?" + query
		}

		GetSugarLogger().Infof("[GIN] %s | %3d | %13v | %15s | %-7s %s",
			end.Format("2006/01/02 - 15:04:05"),
			c.Writer.Status(),
			latency,
			c.ClientIP(),
			c.Request.Method,
			path,
		)
	}
}
