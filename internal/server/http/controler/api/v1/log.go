package v1

import (
	"github.com/gin-gonic/gin"
	"log-server/internal/dto"
	"log-server/internal/utils"
	"strconv"
	"time"
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
			"appCode": c.Query("appcode"),
		},
	}

	result,sum := this.Service.GetLogList(d)
	list := []dto.Log{}
	for _, item := range result {
		list = append(list, dto.Log{
			Id:      item.Id,
			Level:   item.Level,
			Time:    time.Unix(int64(item.Time), 0).Format(utils.DatetimeFormart),
			Content: item.Content,
			Appcode: item.Appcode,
		})
	}

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