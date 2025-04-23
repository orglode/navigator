package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"navigator/api/jwt"
)

func JwtTEstUser(c *gin.Context) {
	token, _ := jwt.GenerateToken(123)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": token,
	})

}

func GetWxInfo(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := svc.WxProgram(ctx)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}
