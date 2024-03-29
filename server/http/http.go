package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/orglode/navigator/conf"
	"github.com/orglode/navigator/model"
	"github.com/orglode/navigator/service"
	"io"
	"os"
	"time"
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
	//gin日志模式
	router.Use(gin.LoggerWithConfig(initGinLog()))
	router.Use(gin.Recovery())
	//初始化路由
	initRouter(router)
	router.Run(conf.Server.Addr)
}

// 初始化gin日志库
func initGinLog() gin.LoggerConfig {
	date := time.Now().Format("20060102")
	f, _ := os.OpenFile("./log/access_"+date+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	var logConf = gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("客户端IP:%s,请求时间:[%s],请求方式:%s,请求地址:%s,响应时间:%s\n",
				param.ClientIP,
				param.TimeStamp.Format("2006年01月02日 15:03:04"),
				param.Method,
				param.Path,
				param.Latency,
			)
		},
		Output: io.MultiWriter(os.Stdout, f),
	}
	return logConf
}
