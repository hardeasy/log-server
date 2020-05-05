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
		Q: map[string]interface{}{
			"level": c.Query("level"),
			"content": c.Query("content"),
		},
	}
	username,_ := c.Keys["username"].(string)

	appcode := c.Query("appcode")
	if len(appcode) > 0 {
		d.Q["appCode"] = []string{appcode}
	} else {
		if username != "admin" {
			userModel := this.Service.GetUserByName(username)
			userAppIds := this.Service.GetUserAppids(userModel.Id)
			appList := this.Service.GetAppByIds(userAppIds)
			allowAppcodes := []string{"_aaa"}
			for _,item := range appList {
				allowAppcodes = append(allowAppcodes, item.Code)
			}
			d.Q["appCode"] = allowAppcodes
		}
	}

	result,sum := this.Service.GetLogList(d)
	list := []dto.Log{}
	for _, item := range result {
		list = append(list, dto.Log{
			Id:      item.Id,
			Level:   item.Level,
			Time:    item.Time,
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
	d := this.Service.GetById(id)
	if d == nil {
		this.Error(c, "not found")
		return
	}
	data := dto.Log{
		Id:      d.Id,
		Level:   d.Level,
		Time:    d.Time,
		Content: d.Content,
		Appcode: d.Appcode,
	}

	this.Echo(c, data, "")
}

func (this *controller) Getindices(c *gin.Context) {
	list := this.Service.GetLogEsIndexList()
	this.Echo(c, list, "")
}

func (this *controller) DeleteIndex(c *gin.Context) {
	indexName := c.Param("index")
	username,_ := c.Keys["username"].(string)
	if username != "admin" {
		this.Error(c, "没权限")
		return
	}
	err := this.Service.DeleteEsIndex(indexName)
	if err != nil {
		this.Error(c, err.Error())
		return
	}
	this.Echo(c, nil, "")
}
