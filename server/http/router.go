package http

import (
	"navigator/api/jwt"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {

	r.GET(apiPrefix+"test", Test)

	v1 := r.Group(apiPrefix+"v1/", jwt.AuthMiddleware())
	{
		v1.GET("/heath", Test)
	}

}
