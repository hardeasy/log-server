package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"log-server/internal/dto"
	"log-server/internal/utils"
	"time"
)

const (
	tokenExpiresTime = time.Second * 3600
)

func (s *Service) Login(d *dto.LoginDto) (rToken string, rErr error) {
	user := s.Dao.UserDao.GetByUsername(s.Dao.Db, d.Username)
	if user == nil || user.Password != s.EncryPasswd(d.Password) {
		return "", errors.New("login error")
	}
	var err error
	rToken, err = s.Auth(d.Username)
	if err != nil {
		return "", errors.New("service error")
	}
	rErr = nil
	return
}

func (s *Service) GenerateLoginToken() string {
	return utils.RandString(20)
}

func (s *Service) Auth(username string) (rToken string, rErr error) {
	var token string
	for i := 0; i < 5; i++ {
		token = s.GenerateLoginToken()
		redisKey := s.GetLoginRedisKey(token)
		_, err := s.Dao.Redis.Get(redisKey).Result()
		if err == redis.Nil {
			s.Dao.Redis.Set(redisKey, username, tokenExpiresTime)
			return token, nil
		}
	}
	return "", errors.New("auth error")
}

func (s *Service) EncryPasswd(passwd string) string {
	has := md5.Sum([]byte(passwd))
	return fmt.Sprintf("%x", has)
}

func (s *Service) CheckLogin(token string, renewTTL bool) (string, error){
	if len(token) == 0 {
		return "", errors.New("not found")
	}
	redisKey := s.GetLoginRedisKey(token)
	result, err := s.Dao.Redis.Get(redisKey).Result()
	if err != nil {
		return "", err
	}
	if len(result) > 0 && renewTTL {
		s.Dao.Redis.Expire(redisKey, tokenExpiresTime)
	}
	return result, nil
}

func (s *Service) GetLoginRedisKey(token string) string {
	return fmt.Sprintf("login_%s", token)
}
