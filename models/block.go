package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type LuckyBlock struct {
	//gorm.Model
	ID          uint            `gorm:"primarykey"`
	Node        string          `gorm:"type:varchar(255)"`
	Height      int64           `gorm:"type:int;not null"`
	Date        time.Time       `gorm:"type:datetime"`
	NodeFrom    string          `gorm:"type:varchar(255)"`
	NodeTo      string          `gorm:"type:varchar(255)"`
	Message     string          `gorm:"type:varchar(255)"`
	RewardValue decimal.Decimal `gorm:"type:decimal(20, 8)"`
	Type        string          `gorm:"type:varchar(50)"`
	CreateTime  time.Time       `gorm:"type:datetime"`
	TimeTag     int64           `gorm:"type:bigint"`
	Category    string          `gorm:"type:varchar(50)"`
	PidNode     string          `gorm:"type:varchar(255)"`
}

func (table *LuckyBlock) TableName() string {
	return "lucky_block"
}

//gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
//	return "tb_" + defaultTableName;
//}

//func Create() {
//	database.DB.AutoMigrate(&LuckyBlock{})
//}
