package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/orglode/go-wake/model"
)

func test(c *gin.Context) {
	req := model.TestAReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		fmt.Println("asd")
	}
	data, _ := svc.Test(req)
	c.JSON(200, data)
	return
}
