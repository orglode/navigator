package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	err := fmt.Errorf("错误提交")
	standardOutput(c, err)
}
