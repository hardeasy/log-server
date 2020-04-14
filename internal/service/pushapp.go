package service

import (
	"encoding/json"
	"fmt"
	"log-server/internal/models"
	"time"
)

func (s *Service) GetPushAppbyCode(code string) *models.PushApp {
	return s.Dao.PushappDao.GetByCode(s.Dao.Db, code)
}

func (s *Service) GetPushAppbyAccessToken(accessToken string) *models.PushApp {
	var pushapp *models.PushApp
	cacheKey := fmt.Sprintf("pushapp_%s", accessToken) //accesstoken -> pushapp 映射
	data, err := s.Dao.Redis.Get(cacheKey).Result()
	if err == nil {
		json.Unmarshal([]byte(data), &pushapp)
		//fmt.Println("cache hit", pushapp)
		return pushapp
	}
	pushapp = s.Dao.PushappDao.GetByAccessToken(s.Dao.Db, accessToken)
	go func() {
		expireTime := time.Second * 3600
		if pushapp == nil { // nil set redis ""
			s.Dao.Redis.Set(cacheKey, "", expireTime)
			return
		}
		data, _ := json.Marshal(pushapp)
		s.Dao.Redis.Set(cacheKey, data, expireTime)
	}()
	return pushapp
}
