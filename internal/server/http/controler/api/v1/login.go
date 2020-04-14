package v1

import (
	"github.com/gin-gonic/gin"
	"log-server/internal/dto"
)

func (this *controller) Login(c *gin.Context) {
	d := &dto.LoginDto{}

	if err := c.ShouldBind(d); err != nil {
		this.Error(c, "请求格式错误")
		return
	}

	token, err := this.Service.Login(d)
	if err != nil {
		this.Error(c, err.Error())
		return
	}

	this.Echo(c, map[string]interface{}{
		"token": token,
	}, "")
}
