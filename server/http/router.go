package http

import (
	"navigator/api/jwt"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {

	r.GET("login", JwtTEstUser)

	api := r.Group("/api/")
	{
		api.GET("info", GetWxInfo)
		api.GET("testError", testError)
	}

	crm := r.Group("/crm", jwt.AuthMiddleware())
	{
		crm.GET("/ping")
	}

}
