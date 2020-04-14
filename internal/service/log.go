package service

import (
	"encoding/json"
	"log-server/internal/dto"
	"log-server/internal/models"
	"time"
)

func (s *Service) GetLogList(d *dto.GeneralListDto) (rList []*models.Log, rSum int) {
	rList, rSum= s.Dao.LogDao.GetListEs(s.Dao.Elastic, d)
	return
}

func (s *Service) GetById(appCode string, id string) *models.Log {
	return s.Dao.LogDao.GetById(s.Dao.Elastic, appCode, id)
}

func (s *Service) Push(d dto.PushLogDto) (rErr error) {
	data, _ := json.Marshal(d)
	s.Dao.Redis.RPush(s.Cfg.PushServer.QueueKey, string(data))
	rErr = nil
	return
}

func (s *Service) PushQueueListen() {
	for {
		if ! s.pushRun {
			close(s.PushEsChan)
			return
		}
		str, _ := s.Dao.Redis.LPop(s.Cfg.PushServer.QueueKey).Result()
		if len(str) == 0 {
			time.Sleep(time.Second * 2)
		}
		var pushDto  dto.PushLogDto
		if err := json.Unmarshal([]byte(str), &pushDto); err == nil {
			s.PushEsChan <- pushDto
		}
	}
}

func (s *Service) PushLogToEs() {
	for item := range s.PushEsChan {
		s.Dao.LogDao.AddPushLog(s.Dao.Elastic, item)
	}
}
