package middleware

import "log-server/internal/service"

type MiddleWare struct {
	Service *service.Service
}

func NewMiddleWare(serve *service.Service) *MiddleWare {
	return &MiddleWare{Service:serve}
}
