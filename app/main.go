package main

import (
	"github.com/orglode/go-wake/conf"
	"github.com/orglode/go-wake/server/http"
	"github.com/orglode/go-wake/service"
)

func main() {
	//初始化配置
	cfg := conf.Init()
	// 初始化日志
	//初始化service
	svc := service.NewService(cfg)
	//初始化http控件
	http.Init(svc, cfg)
}
