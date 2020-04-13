package v1

import (
	"github.com/gin-gonic/gin"
	"log-server/internal/dto"
)

func (this *controller) Login(c *gin.Context) {
	d := &dto.LoginDto{
		Username: "admin",
		Password: "123123",
	}

	token, err := this.Service.Login(d)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"code": -1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, map[string]interface{}{
		"code": 0,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}
