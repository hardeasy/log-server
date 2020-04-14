package v1

import (
	"github.com/gin-gonic/gin"
	"log-server/internal/dto"
)

func (this *controller) GetUserById(c *gin.Context) {
	username := c.Keys["username"].(string)
	d := this.Service.GetUserByName(username)
	if d == nil {
		this.Echo(c, nil, "")
		return
	}
	user := &dto.UserInfo{
		Id:       d.Id,
		Username: d.Username,
		Email:    d.Email,
		Phone:    d.Phone,
	}
	this.Echo(c, user, "")
	return
}
