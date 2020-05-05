package v1

import (
	"github.com/gin-gonic/gin"
	"log-server/internal/dto"
	"log-server/internal/models"
	"log-server/internal/utils"
	"strconv"
)

func (this *controller) GetAppAll(c *gin.Context) {
	var d []*models.App
	username,_ := c.Keys["username"].(string)
	if username == "admin" {
		d = this.Service.GetAppAll()
	} else {
		userModel := this.Service.GetUserByName(username)
		userAppIds := this.Service.GetUserAppids(userModel.Id)
		d = this.Service.GetAppByIds(userAppIds)
	}
	list := []dto.App{}
	if len(d) == 0 {
		this.Echo(c, list, "")
		return
	}
	for _, item := range d {list = append(list, dto.App{
			Id:   item.Id,
			Name: item.Name,
			Code: item.Code,
			AccessToken: item.AccessToken,
			CreatedAt: utils.ConvertTimezone(item.CreatedAt).Format(utils.DatetimeFormart),
			UpdatedAt: utils.ConvertTimezone(item.UpdatedAt).Format(utils.DatetimeFormart),
		})
	}
	this.Echo(c, list, "")
	return
}

func (this *controller) AddApp(c *gin.Context) {
	var data dto.AddApp
	if err := c.ShouldBind(&data); err != nil {
		this.Error(c, err.Error())
		return
	}
	err := this.Service.AddApp(data)
	if err != nil {
		this.Error(c, err.Error())
		return
	}
	this.Echo(c, nil, "")
	return
}

func (this *controller) EditApp(c *gin.Context) {
	var data dto.EditApp
	if err := c.ShouldBind(&data); err != nil {
		this.Error(c, err.Error())
		return
	}
	err := this.Service.EditApp(data)
	if err != nil {
		this.Error(c, err.Error())
		return
	}
	this.Echo(c, nil, "")
	return
}

func (this *controller) GetAppMembers(c *gin.Context) {
	appid,_ := strconv.Atoi(c.Param("appid"))
	memberIds := this.Service.GetAppMembers(appid)
	list := []dto.AppMember{}
	for _, userId := range memberIds {
		userItem := this.Service.GetUserById(userId)
		if userItem != nil {
			list = append(list, dto.AppMember{
				UserId:   userId,
				Username: userItem.Username,
				Email:   userItem.Email,
			})
		}
	}
	this.Echo(c, list, "")
}

func (this *controller) AddAppMember(c *gin.Context) {
	var data dto.AddAppMember
	if err := c.ShouldBind(&data); err != nil {
		this.Error(c, err.Error())
		return
	}
	appid,_ := strconv.Atoi(c.Param("appid"))
	data.AppId = appid
	err := this.Service.AddAppMember(data)
	if err != nil {
		this.Error(c, err.Error())
		return
	}
	this.Echo(c, nil, "")
	return
}

func (this *controller) DeleteAppMember(c *gin.Context) {
	data := dto.DeleteAppMember{}

	appid,_ := strconv.Atoi(c.Param("appid"))
	data.AppId = appid

	userid,_ := strconv.Atoi(c.Param("userid"))
	data.UserId = userid

	err := this.Service.DeleteAppMember(data)
	if err != nil {
		this.Error(c, err.Error())
		return
	}
	this.Echo(c, nil, "")
	return
}
