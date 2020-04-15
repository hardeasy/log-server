package dto

type PushLogDto struct {
	Time int `form:"time" json:"time" binding:"required"`
	Level string `form:"level" json:"level" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	Appcode string `form:"appcode" json:"appcode"`
}

type Log struct {
	Id string `json:"id"`
	Level string `json:"level"`
	Time string `json:"time"`
	Content string `json:"content"`
	Appcode string `json:"appcode"`
}