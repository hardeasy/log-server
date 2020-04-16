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
