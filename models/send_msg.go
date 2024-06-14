package models

import (
	"time"
)

type Msg struct {
	ID         uint       `gorm:"primarykey"`
	Title      string     `gorm:"type:varchar(255)"`
	Node       string     `gorm:"type:varchar(30)"`
	Content    string     `gorm:"type:varchar(255)"`
	CreateTime time.Time  `gorm:"type:datetime"`
	SendTime   *time.Time `gorm:"type:datetime"`
	SendStatus int        `gorm:"type:int"`
}

func (table *Msg) TableName() string {
	return "send_msg"
}
