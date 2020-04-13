package middleware

import (
	"github.com/gin-gonic/gin"
)

func (this *MiddleWare) AuthUserFilter(c *gin.Context) {
	if c.Request.URL.Path == "/api/v1/login"{
		c.Next()
		return
	}
	token := c.GetHeader("X-TOKEN")
	_, err := this.Service.CheckLogin(token, true)
	if err != nil{
		c.JSON(200, map[string]interface{}{
			"code": 401,
			"message": "need auth",
		})
		c.Abort()
		return
	}
	//
	c.Next()
}

func (this *MiddleWare) PushAuthFilter(c *gin.Context) {
	token := c.GetHeader("X-TOKEN")
	if token != this.Service.Cfg.PushServer.Token {
		c.JSON(200, map[string]interface{}{
			"code": 401,
			"message": "need auth",
		})
		c.Abort()
		return
	}
	c.Next()
}
