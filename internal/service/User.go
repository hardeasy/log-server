package service

import (
	"errors"
	"log-server/internal/dto"
	"log-server/internal/models"
	"log-server/internal/utils"
	"time"
)

func (s *Service) GetUserList(d *dto.GeneralListDto) (rList []*models.User, rSum int) {
	rList, rSum= s.Dao.UserDao.GetList(s.Dao.Db, d)
	return
}

func (s *Service) AddUser(user dto.AddUser) error {
	repeat := s.Dao.UserDao.GetByUsername(s.Dao.Db, user.Username)
	if repeat != nil {
		return errors.New("账户重复")
	}
	repeat = s.Dao.UserDao.GetByUserEmail(s.Dao.Db, user.Email)
	if repeat != nil {
		return errors.New("Email重复")
	}

	userModel := &models.User{
		Username:  user.Username,
		Password:  EncryPasswd(user.Password),
		Email:     user.Email,
		CreatedAt: utils.ConvertTimezone(time.Now()),
		UpdatedAt: utils.ConvertTimezone(time.Now()),
		IsDisable:      2,
	}
	return s.Dao.UserDao.Save(s.Dao.Db, userModel)
}

func (s *Service) EditUser(user dto.EditUser) error {
	userModel := s.Dao.UserDao.GetByUseId(s.Dao.Db, user.Id)
	if userModel == nil {
		return errors.New("没有这个用户")
	}
	//email check
	repeatUser := s.Dao.UserDao.GetByUserEmail(s.Dao.Db, user.Email)
	if repeatUser != nil && repeatUser.Id != userModel.Id {
		return errors.New("email重复")
	}

	userModel.IsDisable = user.IsDisable
	userModel.Email = user.Email
	if len(user.Password) > 0 {
		userModel.Password = EncryPasswd(user.Password)
	}

	return s.Dao.UserDao.Save(s.Dao.Db, userModel)
}

func EncryPasswd(password string) string {
	return utils.GetMD5String(password)
}