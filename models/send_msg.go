package models

import (
	"time"
)

type SendType int

const (
	SectorsError SendType = 101 // 扇区错误
	HeightDelay  SendType = 102 // 高度延迟
	LuckyLow     SendType = 103 // 幸运值过低
	OrphanBlock  SendType = 104 // 孤块
)

type Msg struct {
	ID         uint       `gorm:"primarykey"`
	Title      string     `gorm:"type:varchar(255)"`
	Node       string     `gorm:"type:varchar(30)"`
	Content    string     `gorm:"type:varchar(255)"`
	CreateTime time.Time  `gorm:"type:datetime"`
	SendTime   *time.Time `gorm:"type:datetime"`
	Type       SendType   `gorm:"type:int"`
	SendStatus int        `gorm:"type:int"`
}

func (table *Msg) TableName() string {
	return "send_msg"
}
