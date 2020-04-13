package service

import (
	"log-server/config"
	"log-server/internal/dao"
	"log-server/internal/dto"
)

type Service struct {
	Cfg *config.Config
	Dao *dao.Dao
	pushRun bool
	PushEsChan chan dto.PushLogDto
}

func NewService(cfg *config.Config, d *dao.Dao) *Service {
	serv :=  &Service{
		Cfg: cfg,
		Dao: d,
		pushRun: true,
		PushEsChan: make(chan dto.PushLogDto, 500),
	}
	go func() {
		serv.PushQueueListen()
	}()

	go func() {
		serv.PushLogToEs()
	}()

	return serv
}

func (s *Service) Stop() {
	s.pushRun = false
	return
}
