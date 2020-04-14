package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (this *MiddleWare) AuthUserFilter(c *gin.Context) {
	if c.Request.URL.Path == "/api/v1/login"{
		c.Next()
		return
	}
	token := c.GetHeader("X-TOKEN")
	username, err := this.Service.CheckLogin(token, true)
	if err != nil{
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 401,
			"message": "need auth",
		})
		c.Abort()
		return
	}
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys["username"] = username
	//
	c.Next()
}

func (this *MiddleWare) PushAuthFilter(c *gin.Context) {
	token := c.GetHeader("X-TOKEN")
	//查询
	pushApp := this.Service.GetPushAppbyAccessToken(token)
	if pushApp == nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 401,
			"message": "need auth",
		})
		c.Abort()
		return
	}
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys["appCode"] = pushApp.Code
	c.Next()
}
