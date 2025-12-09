package http

import (
	"errors"
	apiErr "navigator/api/error"
	"navigator/conf"
	"navigator/model"
	"navigator/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/orglode/hades/logger"
	"github.com/orglode/hades/trace"
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
	// 错误中间件
	router.Use(ErrorHandlerMiddleware())

	// 初始化路由
	initRouter(router)

	// 启动服务
	router.Run(conf.Server.Addr)
}

// 错误处理中间件
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 如果有错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			var appErr *apiErr.AppError
			if errors.As(err, &appErr) {
				responseError(c, appErr.Code, appErr.Message)
			} else {
				responseError(c, model.SystemErr, err.Error())
			}
			return
		}
	}
}

// responseSuccess 成功响应
func responseSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, model.HttpResponse{
		Code:    model.SuccessCode,
		Message: "success",
		Data:    data,
	})
}

// ResponseError 错误响应
func responseError(c *gin.Context, code int, message ...string) {
	c.JSON(http.StatusOK, model.HttpResponse{
		Code:    code,
		Message: message[0],
		Data:    nil,
	})
}
