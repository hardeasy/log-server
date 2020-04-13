package v1

import (
	"github.com/gin-gonic/gin"
	"log-server/internal/dto"
	"net/http"
)

func (this *controller) Push(c *gin.Context) {
	var data dto.PushLogDto
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"errmsg": "请求格式出错",
		})
		return
	}
	err := this.Service.Push(data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"errmsg": "调用失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"errmsg": "",
	})
}
