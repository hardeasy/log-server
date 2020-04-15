package v1

import (
	"github.com/gin-gonic/gin"
	"log-server/internal/dto"
)

func (this *controller) GetPushAppAll(c *gin.Context) {
	d := this.Service.GetPushAppAll()
	if len(d) == 0 {
		this.Echo(c, nil, "")
		return
	}
	list := []dto.Pushapp{}
	for _, item := range d {
		list = append(list, dto.Pushapp{
			Id:   item.Id,
			Name: item.Name,
			Code: item.Code,
		})
	}
	this.Echo(c, list, "")
	return
}
