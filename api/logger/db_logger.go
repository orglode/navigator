package logger

import (
	"context"
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// 创建一个自定义的日志记录器
type DbLogger struct {
	SlowThreshold time.Duration
	LogLevel      logger.LogLevel
	Colorful      bool
	// 使用 GORM 的默认日志格式化器
	defaultLogger logger.Interface
}

func (c *DbLogger) LogMode(level logger.LogLevel) logger.Interface {
	c.LogLevel = level
	return c
}

func (c *DbLogger) Info(ctx context.Context, message string, data ...interface{}) {
	if c.LogLevel >= logger.Info {
		c.writeLog(ctx, "info", message, data...)
	}
}

func (c *DbLogger) Warn(ctx context.Context, message string, data ...interface{}) {
	if c.LogLevel >= logger.Warn {
		c.writeLog(ctx, "warn", message, data...)
	}
}

func (c *DbLogger) Error(ctx context.Context, message string, data ...interface{}) {
	if c.LogLevel >= logger.Error {
		c.writeLog(ctx, "error", message, data...)
	}
}

func (c *DbLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if c.LogLevel <= logger.Silent {
		return
	}

	sql, rowsAffected := fc()
	elapsed := time.Since(begin)
	switch {
	case elapsed > c.SlowThreshold && c.LogLevel >= logger.Warn:
		c.writeLog(ctx, "slow", sql, rowsAffected, elapsed, err)
	case c.LogLevel >= logger.Info:
		c.writeLog(ctx, "info", sql, rowsAffected, elapsed)
	}
}

func (c *DbLogger) writeLog(ctx context.Context, level string, message string, data ...interface{}) {
	_, file, line, _ := runtime.Caller(3)
	fileInfo := fmt.Sprintf("%s:%d", filepath.Base(file), line)
	// 格式化日志消息
	message += fmt.Sprintf("  row:%d time:%v", data...)
	var formattedLogMessage string
	formattedLogMessage = fmt.Sprintf("%s =>%s ", message, fileInfo)
	if c.Colorful {
		switch level {
		case "info":
			c.defaultLogger.Info(ctx, formattedLogMessage)
		case "warn":
			c.defaultLogger.Warn(ctx, formattedLogMessage)
		case "error":
			c.defaultLogger.Error(ctx, formattedLogMessage)
		case "slow":
			slowMessage := fmt.Sprintf("slow SQL: %s", message)
			for _, v := range data {
				slowMessage += " " + fmt.Sprintf("%v", v)
			}
			c.defaultLogger.Warn(ctx, slowMessage, data...)
		}
	} else {
		logMessage := fmt.Sprintf("[%s] %s", level, message)
		for _, v := range data {
			logMessage += " " + fmt.Sprintf("%v", v)
		}
		formattedLogMessage = logMessage
		fmt.Println(formattedLogMessage)
	}

	//formattedLogMessage += "\n"
	//filePath, err := GetLogFilePath()
	//if err != nil {
	//	return
	//}
	//
	//logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	return
	//}
	//defer logFile.Close()
	//logFile.WriteString(formattedLogMessage)
}

func NewDbLogger(fileLogger *log.Logger, slowThreshold time.Duration, logLevel logger.LogLevel, colorful bool) *DbLogger {
	defaultLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: slowThreshold,
			LogLevel:      logLevel,
			Colorful:      colorful,
		},
	)

	return &DbLogger{
		SlowThreshold: slowThreshold,
		LogLevel:      logLevel,
		Colorful:      colorful,
		defaultLogger: defaultLogger,
	}
}

func GetLogFilePath() (string, error) {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()

	// 创建年月文件夹
	yearMonthPath := filepath.Join("logs", fmt.Sprintf("%d-%02d", year, month))
	if err := os.MkdirAll(yearMonthPath, 0777); err != nil {
		return "", err
	}

	// 生成日志文件名
	logFileName := fmt.Sprintf("sql_%d%02d%02d.log", year, month, day)
	filePath := filepath.Join(yearMonthPath, logFileName)

	return filePath, nil
}
