package models

type User struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Username string
	Password string
	Email string
	Phone string
	CreateAt int
	AuthKey string
	AccessToken string
	Open int
}

func (User) TableName() string {
	return "tbl_admin_user"
}
