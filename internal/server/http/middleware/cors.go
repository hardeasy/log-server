package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (this *MiddleWare) CorsFilter(c *gin.Context) {
	method := c.Request.Method

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, X-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")

	c.Header("Access-Control-Allow-Credentials", "true")

	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.Next()
}
