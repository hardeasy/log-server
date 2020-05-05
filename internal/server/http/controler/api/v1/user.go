package v1

import (
	"github.com/gin-gonic/gin"
	"log-server/internal/dto"
	"log-server/internal/utils"
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

func (this *controller) GetUserList(c *gin.Context) {
	page,err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	if page < 1 {
		page = 1
	}

	pageSize,err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		pageSize = 30
	}

	d := &dto.GeneralListDto{
		Offset: (page - 1) * pageSize,
		Limit:  pageSize,
		Order:  "",
		Q: map[string]interface{}{
			"username": c.Query("username"),
			"email": c.Query("email"),
		},
	}
	result,sum := this.Service.GetUserList(d)
	list := []dto.UserInfo{}
	for _,item := range result {
		list = append(list, dto.UserInfo{
			Id:       item.Id,
			Username: item.Username,
			Email:    item.Email,
			CreatedAt: item.CreatedAt.Format(utils.DatetimeFormart),
			IsDisable: item.IsDisable,
		})
	}
	this.Echo(c, map[string]interface{}{
		"items": list,
		"total": sum,
	}, "")
}
