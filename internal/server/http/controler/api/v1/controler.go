package v1

import (
	"log-server/internal/service"
)

type controller struct {
	Service *service.Service
}

func NewController(service *service.Service) *controller {
	return &controller{service}
}