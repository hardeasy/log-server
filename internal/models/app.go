package models

import "time"

type App struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string
	Code string
	AccessToken string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (App) TableName() string {
	return "tbl_app"
}
