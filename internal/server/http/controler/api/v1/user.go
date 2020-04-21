package v1

import (
	"github.com/gin-gonic/gin"
	"log-server/internal/dto"
	"strconv"
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
	}
	this.Echo(c, user, "")
	return
}

func (this *controller) AddUser(c *gin.Context) {
	var data dto.AddUser
	if err := c.ShouldBind(&data); err != nil {
		this.Error(c, "请求格式出错")
		return
	}
	err := this.Service.AddUser(data)
	if err != nil {
		this.Error(c, err.Error())
		return
	}
	this.Echo(c, "", "")
	return
}

func (this *controller) EditUser(c *gin.Context) {
	var data dto.EditUser
	if err := c.ShouldBind(&data); err != nil {
		this.Error(c, "请求格式出错")
		return
	}
	data.Id,_ = strconv.Atoi(c.Param("id"))
	if c.Keys["username"] != "admin" {
		this.Error(c, "权限不够")
		return
	}

	err := this.Service.EditUser(data)
	if err != nil {
		this.Error(c, err.Error())
		return
	}
	this.Echo(c, "", "")
	return
}