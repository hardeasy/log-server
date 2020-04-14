package v1

import (
	"github.com/gin-gonic/gin"
	"log-server/internal/service"
	"net/http"
)

type controller struct {
	Service *service.Service
}

func NewController(service *service.Service) *controller {
	return &controller{service}
}

func (this *controller) Error(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"message": message,
	})
}

func (this *controller) Echo(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"data": data,
		"message": message,
	})
}
