package models

type AppMember struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	AppId int
	UserId int
}

func (AppMember) TableName() string {
	return "tbl_app_member"
}
