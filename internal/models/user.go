package models

import "time"

type User struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Username string
	Password string
	Email string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsOpen int
}

func (User) TableName() string {
	return "tbl_admin_user"
}
