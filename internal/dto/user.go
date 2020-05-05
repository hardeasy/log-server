package dto

type UserInfo struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	CreatedAt string `json:"created_at"`
	IsDisable int `json:"is_disable"`
}

type AddUser struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EditUser struct {
	Id int `json:"id"`
	IsDisable int `json:"is_disable"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password"`
}

type DeleteUser struct {
	Id int `json:"id" binding:"required"`
}
