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
			"url": c.Query("url"),
			"category": c.Query("category"),
			"content": c.Query("content"),
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
