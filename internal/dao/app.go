package dao

import (
	"github.com/jinzhu/gorm"
	"log-server/internal/models"
)

type AppDao struct {
	
}

func (this *AppDao) GetById(db *gorm.DB, id int) *models.App {
	pushapp := &models.App{}
	err := db.Where("id = ?", id).Find(pushapp).Error
	if err != nil {
		return nil
	}
	return pushapp
}

func (this *AppDao) GetByCode(db *gorm.DB, code string) *models.App {
	pushapp := &models.App{}
	err := db.Where("code = ?", code).Find(pushapp).Error
	if err != nil {
		return nil
	}
	return pushapp
}

func (this *AppDao) GetByAccessToken(db *gorm.DB, accessToken string) *models.App {
	pushapp := &models.App{}
	err := db.Where("access_token = ?", accessToken).Find(pushapp).Error
	if err != nil {
		return nil
	}
	return pushapp
}

func (this *AppDao) GetAll(db *gorm.DB) []*models.App {
	list := []*models.App{}
	db.Find(&list)
	return list
}

func (this *AppDao) Add(db *gorm.DB, app *models.App) error {
	return db.Save(app).Error
}

func (this *AppDao) Edit(db *gorm.DB, app *models.App) error {
	return db.Save(app).Error
}

func (this *AppDao) GetAppMembers(db *gorm.DB, appId int) []int {
	list := []int{}
	models := []*models.AppMember{}
	db.Where("app_id = ?", appId).Order("id desc").Find(&models)
	for _, item := range models {
		list = append(list, item.UserId)
	}
	return list
}

func (this *AppDao) GetUserApps(db *gorm.DB, userId int) []int {
	list := []int{}
	models := []*models.AppMember{}
	db.Where("user_id = ?", userId).Find(&models)
	for _, item := range models {
		list = append(list, item.AppId)
	}
	return list
}

func (this *AppDao) AddAppMember(db *gorm.DB, appId int, userId int) error {
	item := &models.AppMember{
		AppId:  appId,
		UserId: userId,
	}
	db.Save(item)
	return nil
}

func (this *AppDao) DeleteAppMember(db *gorm.DB, appId int, userId int) error {
	db.Where("user_id = ? and app_id = ?", userId, appId).Delete(&models.AppMember{})
	return nil
}
