package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log-server/internal/dto"
	"log-server/internal/models"
	"log-server/internal/utils"
	"time"
)

func (s *Service) GetAppbyCode(code string) *models.App {
	return s.Dao.AppDao.GetByCode(s.Dao.Db, code)
}

func (s *Service) GetAppbyId(id int) *models.App {
	return s.Dao.AppDao.GetById(s.Dao.Db, id)
}

func (s *Service) GetAppbyAccessToken(accessToken string) *models.App {
	var app *models.App
	cacheKey := fmt.Sprintf("pushapp_%s", accessToken) //accesstoken -> pushapp 映射
	data, err := s.Dao.Redis.Get(cacheKey).Result()
	if err == nil {
		json.Unmarshal([]byte(data), &app)
		//fmt.Println("cache hit", app)
		return app
	}
	app = s.Dao.AppDao.GetByAccessToken(s.Dao.Db, accessToken)
	go func() {
		expireTime := time.Second * 3600
		if app == nil { // nil set redis ""
			s.Dao.Redis.Set(cacheKey, "", expireTime)
			return
		}
		data, _ := json.Marshal(app)
		s.Dao.Redis.Set(cacheKey, data, expireTime)
	}()
	return app
}

func (s Service) GetAppAll() []*models.App {
	return s.Dao.AppDao.GetAll(s.Dao.Db)
}

func (s Service) AddApp(app dto.AddApp) error {
	//check appcode
	repeatModel := s.Dao.AppDao.GetByCode(s.Dao.Db, app.Code)
	if repeatModel != nil {
		return errors.New("编码重复")
	}
	if len(app.AccessToken) == 0 {
		app.AccessToken = utils.RandString(32)
	}
	model := &models.App{
		Name:        app.Name,
		Code:        app.Code,
		AccessToken: app.AccessToken,
		CreatedAt:   utils.ConvertTimezone(time.Now()),
		UpdatedAt:   utils.ConvertTimezone(time.Now()),
	}
	return s.Dao.AppDao.Add(s.Dao.Db, model)
}

func (s Service) EditApp(app dto.EditApp) error {
	model := s.GetAppbyId(app.Id)
	if model == nil {
		return errors.New("没有这个APP")
	}
	//check appcode
	repeatModel := s.Dao.AppDao.GetByCode(s.Dao.Db, app.Code)
	if repeatModel != nil && repeatModel.Id != model.Id {
		return errors.New("编码重复")
	}
	model.Code = app.Code
	model.AccessToken = app.AccessToken
	model.Name = app.Name
	model.UpdatedAt = utils.ConvertTimezone(time.Now())
	return s.Dao.AppDao.Edit(s.Dao.Db, model)
}