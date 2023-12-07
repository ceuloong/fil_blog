package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type TronScan struct {
	//gorm.Model
	ID           uint            `gorm:"primarykey" json:"ID,omitempty"`
	TxId         string          `gorm:"type:varchar(255)" json:"txId,omitempty"`
	Height       int64           `gorm:"type:int;not null" json:"height,omitempty"`
	Date         time.Time       `gorm:"type:datetime" json:"date"`
	FromAddr     string          `gorm:"type:varchar(255)" json:"fromAddr,omitempty"`
	FromTag      string          `gorm:"type:varchar(50)" json:"fromTag,omitempty"`
	ToAddr       string          `gorm:"type:varchar(255)" json:"toAddr,omitempty"`
	ToTag        string          `gorm:"type:varchar(50)" json:"toTag,omitempty"`
	Contract     string          `gorm:"type:varchar(255)" json:"contract,omitempty"`
	Quantity     decimal.Decimal `gorm:"type:decimal(20, 8)" json:"quantity"`
	EventType    string          `gorm:"type:varchar(50)" json:"eventType,omitempty"`
	ContractType string          `gorm:"type:varchar(50)" json:"contractType,omitempty"`
	Confirmed    bool            `gorm:"type:tinyint" json:"confirmed,omitempty"`
	CreateTime   time.Time       `gorm:"type:datetime" json:"createTime"`
}

func (table *TronScan) TableName() string {
	return "tron"
}

//gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
//	return "tb_" + defaultTableName;
//}

//func Create() {
//	common.DB.AutoMigrate(&TronScan{})
//}
