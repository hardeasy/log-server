package v1

import (
	"github.com/gin-gonic/gin"
	"log-server/internal/dto"
	"strconv"
)

func (this *controller) GetLogList(c *gin.Context) {
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
		Q: map[string]string{
			"level": c.Query("level"),
			"content": c.Query("content"),
			"appCode": c.Param("appcode"),
		},
	}

	list,sum := this.Service.GetLogList(d)

	c.JSON(200, map[string]interface{}{
		"code": 0,
		"data": map[string]interface{}{
			"items": list,
			"total": sum,
		},
	})
}

func (this *controller) GetLogDetail(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		this.Error(c, "id not found")
		return
	}
	appcode := c.Param("appcode")
	if len(appcode) == 0 {
		this.Error(c, "appcode not found")
		return
	}
	d := this.Service.GetById(appcode, id)
	if d == nil {
		this.Error(c, "not found")
		return
	}
	this.Echo(c, d, "")
}