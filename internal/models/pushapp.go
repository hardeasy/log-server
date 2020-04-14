package models

import "time"

type PushApp struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string
	Code string
	AccessToken string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (PushApp) TableName() string {
	return "tbl_push_app"
}
