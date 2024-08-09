package models

import (
	"time"

	"github.com/shopspring/decimal"
)

/**
*  报块表
 * @Description: block表
*/
type Block struct {
	Height      int64           `gorm:"primarykey;type:int;not null" json:"height"`
	Node        string          `gorm:"type:varchar(255)" json:"node"`
	BlockTime   time.Time       `gorm:"type:datetime" json:"block_time"`
	NodeFrom    string          `gorm:"type:varchar(255)" json:"node_from"`
	NodeTo      string          `gorm:"type:varchar(255)" json:"node_to"`
	Message     string          `gorm:"type:varchar(255)" json:"message"`
	RewardValue decimal.Decimal `gorm:"type:decimal(20, 8)" json:"reward_value"`
	MsgCount    int             `gorm:"type:int" json:"msg_count"`
	BlockSize   int             `gorm:"type:int" json:"block_size"`
	Status      int             `gorm:"type:varchar(50)" json:"status"`
	CreateTime  time.Time       `gorm:"type:datetime" json:"create_time"`
}

func (table *Block) TableName() string {
	return "block"
}
