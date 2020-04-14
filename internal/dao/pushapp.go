package dao

import (
	"github.com/jinzhu/gorm"
	"log-server/internal/models"
)

type PushappDao struct {
	
}

func (this *PushappDao) GetByCode(db *gorm.DB, code string) *models.PushApp {
	pushapp := &models.PushApp{}
	err := db.Where("code = ?", code).Find(pushapp).Error
	if err != nil {
		return nil
	}
	return pushapp
}

func (this *PushappDao) GetByAccessToken(db *gorm.DB, accessToken string) *models.PushApp {
	pushapp := &models.PushApp{}
	err := db.Where("access_token = ?", accessToken).Find(pushapp).Error
	if err != nil {
		return nil
	}
	return pushapp
}
