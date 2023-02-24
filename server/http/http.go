package http

import (
	"github.com/gin-gonic/gin"
	"github.com/orglode/go-wake/conf"
	"github.com/orglode/go-wake/service"
)

var (
	svc *service.Service
)

func Init(s *service.Service, conf *conf.Config) {
	svc = s
	r := gin.Default()
	initRouter(r)
	r.Run(conf.HttpAddr)
}
