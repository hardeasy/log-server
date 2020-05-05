package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log-server/internal/dto"
	"log-server/internal/models"
)

type UserDao struct {

}

func (this *UserDao) GetList(db *gorm.DB, d *dto.GeneralListDto) (rList []*models.User, rSum int) {
	rList = []*models.User{}
	rSum = 0

	if username,ok := d.Q["username"].(string); ok && len(username) >0 {
		db = db.Where("username like ?", fmt.Sprintf("%%%s%%", username))
	}

	if email,ok := d.Q["email"].(string); ok && len(email) >0 {
		db = db.Where("email like ?", fmt.Sprintf("%%%s%%", email))
	}

	orderBy := "id desc"
	if len(d.Order) > 0 {
		orderBy = d.Order
	}
	db.Model(&models.User{}).Count(&rSum)
	db.Order(orderBy).Offset(d.Offset).Limit(d.Limit).Find(&rList)
	return
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
