package dto

type App struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	AccessToken string `json:"access_token"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AddApp struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type EditApp struct {
	Id int `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type AddAppMember struct {
	AppId int `json:"app_id"`
	UserIds []int `json:"user_ids" binding:"required"`
}

type DeleteAppMember struct {
	AppId int `json:"app_id"`
	UserId int `json:"user_id"`
}

type AppMember struct {
	UserId int `json:"user_id"`
	Username string `json:"username"`
	Email string `json:"email"`
}
