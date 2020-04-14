package models

type Log struct {
	Id string `json:"id"`
	Level string `json:"level"`
	Time int `json:"time"`
	Content string `json:"content"`
	Appcode string `json:"appcode"`
}
