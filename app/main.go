package main

import (
	"github.com/orglode/hades/logger"
	"navigator/conf"
	"navigator/server/http"
	"navigator/service"
	"time"
)

func main() {
	// 初始化配置
	cfg := conf.Init()

	// 初始化日志
	config := logger.Config{
		LogDir:       "./logs",
		MaxAge:       30 * 24 * time.Hour,
		RotationTime: 24 * time.Hour,
		Level:        "debug",
		JSONFormat:   true,
	}
	if err := logger.InitLogger(config); err != nil {
		panic(err)
	}
	defer logger.Close()

	// 初始化service
	svc := service.NewService(cfg)

	// 初始化http控件
	http.Init(svc, cfg)
}
