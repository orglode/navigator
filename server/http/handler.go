package http

import (
	"navigator/api/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func testError(c *gin.Context) {
	data, err := svc.TestError(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}
	responseSuccess(c, data)
}

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
