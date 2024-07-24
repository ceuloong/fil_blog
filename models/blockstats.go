package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type BlockStats struct {
	ID                    uint            `json:"id" gorm:"primarykey"`
	Node                  string          `json:"node" gorm:"type:varchar(50);comment:Node"`
	StatsType             string          `json:"statsType" gorm:"type:varchar(50);comment:StatsType"`
	DurationHour          int             `json:"durationHour" gorm:"type:int;comment:DurationHour"`
	BlocksGrowth          int             `json:"blocksGrowth" gorm:"type:int;comment:BlocksGrowth"`
	BlocksRewardGrowthFil decimal.Decimal `json:"blocksRewardGrowthFil" gorm:"type:decimal(20,8);comment:BlocksRewardGrowthFil"`
	HeightTimeStr         string          `json:"heightTimeStr" gorm:"type:varchar(50);comment:HeightTimeStr"`
	HeightTime            time.Time       `json:"heightTime" gorm:"type:datetime;comment:HeightTime"`
	CreatedAt             time.Time       `json:"createdAt" gorm:"type:datetime;comment:CreatedTime"`
}

func (BlockStats) TableName() string {
	return "block_stats"
}
