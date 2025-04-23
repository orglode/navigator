package http

import (
	"github.com/gin-gonic/gin"
	"navigator/api/jwt"
	"net/http"
)

func JwtTEstUser(c *gin.Context) {
	token, _ := jwt.GenerateToken(123)
	responseSuccess(c, token)

}

func GetWxInfo(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := svc.WxProgram(ctx)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	responseSuccess(c, data)
}
