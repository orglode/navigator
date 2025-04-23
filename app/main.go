package main

import (
	"navigator/conf"
	"navigator/server/http"
	"navigator/service"
)

func main() {
	//初始化配置
	cfg := conf.Init()
	//初始化service
	svc := service.NewService(cfg)
	//初始化http控件
	http.Init(svc, cfg)
}
