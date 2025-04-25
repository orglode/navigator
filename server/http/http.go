package http

import (
	"github.com/gin-gonic/gin"
	"github.com/orglode/hades/logger"
	"github.com/orglode/hades/trace"
	"navigator/conf"
	"navigator/model"
	"navigator/service"
	"net/http"
)

var (
	svc *service.Service
)

func Init(s *service.Service, conf *conf.Config) {
	svc = s
	router := gin.New()

	//禁用调式终端颜色
	gin.DisableConsoleColor()

	//判断环境 是否开启debug模式
	if conf.Server.Env == model.EnvProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	// 注册trace中间件
	router.Use(trace.TraceIDMiddleware())

	// gin日志模式
	//router.Use(gin.LoggerWithConfig(initGinLog()))
	router.Use(logger.GinMiddleware())

	// gin恢复模式
	router.Use(gin.Recovery())

	// 初始化路由
	initRouter(router)

	// 启动服务
	router.Run(conf.Server.Addr)
}

func responseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

func responseError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code": 0,
		"msg":  message,
		"data": nil,
	})
}
