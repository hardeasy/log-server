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
	AccessToken string `json:"access_token"`
}

type EditApp struct {
	Id int `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
	AccessToken string `json:"access_token"`
}
