package models

import "time"

type Log struct {
	Id string `json:"id"`
	Level string `json:"level"`
	Time time.Time `json:"time"`
	Content string `json:"content"`
	Appcode string `json:"appcode"`
}
