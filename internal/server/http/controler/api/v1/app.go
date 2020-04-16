package v1

import (
	"github.com/gin-gonic/gin"
	"log-server/internal/dto"
	"log-server/internal/utils"
)

func (this *controller) GetAppAll(c *gin.Context) {
	d := this.Service.GetAppAll()
	if len(d) == 0 {
		this.Echo(c, nil, "")
		return
	}
	list := []dto.App{}
	for _, item := range d {
		list = append(list, dto.App{
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
