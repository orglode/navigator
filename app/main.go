package main

import (
	"navigator/conf"
	"navigator/server/http"
	"navigator/service"
	"time"

	logger "github.com/orglode/hades/logger_v2"
)

func main() {
	// 初始化配置
	cfg := conf.Init()

	// 1. 初始化日志
	if err := logger.InitLogger(logger.Config{
		LogDir:       cfg.Server.LogsPath,
		MaxAge:       30 * 24 * time.Hour,
		RotationTime: 24 * time.Hour,
		Level:        "info",
		JSONFormat:   true,
	}); err != nil {
		panic("logger init error")
	}

	defer logger.Close()

	// 初始化service
	svc := service.NewService(cfg)

	// 初始化http控件
	http.Init(svc, cfg)
}
