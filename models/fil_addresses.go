package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type FilAddresses struct {
	//gorm.Model
	ID               uint            `gorm:"primarykey"`
	Node             string          `gorm:"type:varchar(255)"`
	AccountId        string          `gorm:"type:varchar(50)"`
	Address          string          `gorm:"type:varchar(255)"`
	Message          string          `gorm:"type:varchar(255)"`
	Balance          decimal.Decimal `gorm:"type:decimal(20, 8)"`
	Type             string          `gorm:"type:varchar(50)"`
	CreateTime       time.Time       `gorm:"type:datetime" 地址创建时间`
	CreatedTime      time.Time       `gorm:"type:datetime" 记录创建时间`
	AccountType      string          `gorm:"type:varchar(50)"`
	LastTransferTime time.Time       `gorm:"type:datetime"`
	Nonce            int64           `gorm:"type:bigint"`
	ReceiveAmount    decimal.Decimal `gorm:"type:decimal(20, 8)" 节点接收数量`
	BurnAmount       decimal.Decimal `gorm:"type:decimal(20, 8)" 节点销毁数量`
	SendAmount       decimal.Decimal `gorm:"type:decimal(20, 8)" 节点发送数量`
	TransferCount    int64           `gorm:"type:int" 转账交易数`
	TimeTag          int64           `gorm:"type:bigint" 时间标签`
}

func (table *FilAddresses) TableName() string {
	return "fil_addresses"
}
