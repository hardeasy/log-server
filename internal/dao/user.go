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

func (this *UserDao) GetByUseId(db *gorm.DB, id int) *models.User {
	user := &models.User{}
	err := db.Where("id = ?", id).Find(user).Error
	if err != nil {
		return nil
	}
	return user
}

func (this *UserDao) GetByUserEmail(db *gorm.DB, email string) *models.User {
	user := &models.User{}
	err := db.Where("email = ?", email).Find(user).Error
	if err != nil {
		return nil
	}
	return user
}

func (this *UserDao) Save (db *gorm.DB, user *models.User) error {
	return db.Save(user).Error
}
