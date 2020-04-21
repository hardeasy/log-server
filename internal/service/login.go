package service

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"log-server/internal/dto"
	"log-server/internal/models"
	"log-server/internal/utils"
	"time"
)

const (
	tokenExpiresTime = time.Second * 3600
)

func (s *Service) Login(d *dto.LoginDto) (rToken string, rErr error) {
	var err error
	//repeat login check
	userTokenRedisKey := s.GetLoginUserRedisKey(d.Username)
	token, err2 := s.Dao.Redis.Get(userTokenRedisKey).Result()
	if err2 != redis.Nil {
		_, err2 := s.CheckLogin(token, true)
		if err2 == nil {
			rToken = token
			rErr = nil
			return
		}
	}

	user := s.Dao.UserDao.GetByUsername(s.Dao.Db, d.Username)
	if user == nil || user.Password != EncryPasswd(d.Password) {
		return "", errors.New("login error")
	}
	rToken, err = s.Auth(d.Username)
	if err != nil {
		return "", errors.New("service error")
	}
	rErr = nil
	return
}

func (s *Service) Logout(token string) (rErr error) {
	rErr = nil
	redisKey := s.GetLoginRedisKey(token)
	username,err  := s.Dao.Redis.Get(redisKey).Result()
	if err != nil {
		userTokenRedisKey := s.GetLoginUserRedisKey(username)
		s.Dao.Redis.Del(userTokenRedisKey)
	}
	s.Dao.Redis.Del(redisKey)
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
		userTokenRedisKey := s.GetLoginUserRedisKey(username)
		_, err := s.Dao.Redis.Get(redisKey).Result()
		if err == redis.Nil {
			s.Dao.Redis.Set(redisKey, username, tokenExpiresTime) //token -> username
			s.Dao.Redis.Set(userTokenRedisKey, token, tokenExpiresTime) //username -> token
			return token, nil
		}
	}
	return "", errors.New("auth error")
}

func (s *Service) CheckLogin(token string, renewTTL bool) (username string, err error){
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

func (s *Service) GetLoginUserRedisKey(username string) string {
	return fmt.Sprintf("user_%s", username)
}

func (s *Service) GetUserByName(username string) *models.User {
	return s.Dao.UserDao.GetByUsername(s.Dao.Db, username)
}
func (s *Service) GetUserById(id int) *models.User {
	return s.Dao.UserDao.GetByUseId(s.Dao.Db, id)
}