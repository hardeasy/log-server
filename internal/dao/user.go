package dao

import (
	"github.com/jinzhu/gorm"
	"log-server/internal/models"
)

type UserDao struct {

}

func (this *UserDao) GetByUsername(db *gorm.DB, username string) *models.User {
	user := &models.User{}
	err := db.Where("username = ?", username).Find(user).Error
	if err != nil {
		return nil
	}
	return user
}
