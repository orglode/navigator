package http

import (
	"github.com/gin-gonic/gin"
	"navigator/api/jwt"
)

func initRouter(r *gin.Engine) {

	r.GET("login", JwtTEstUser)

	api := r.Group("/api/")
	{
		api.GET("info", GetWxInfo)
	}

	crm := r.Group("/crm", jwt.AuthMiddleware())
	{
		crm.GET("/ping")
	}

}
