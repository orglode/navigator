package http

import (
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {

	api := r.Group("/api/")
	{
		api.GET("info", GetWxInfo)
	}

	crm := r.Group("/crm")
	{
		crm.GET("/ping")
	}

}
