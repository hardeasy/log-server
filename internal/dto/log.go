package dto

type PushLogDto struct {
	Time int `form:"time" json:"time"`
	Level string `form:"level" json:"level"`
	Data string `form:"data" json:"data"`
}
