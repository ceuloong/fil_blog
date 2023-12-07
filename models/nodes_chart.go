package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type NodesChart struct {
	ID                           uint            `gorm:"primarykey"`
	Node                         string          `gorm:"type:varchar(255)"`
	AvailableBalance             decimal.Decimal `gorm:"type:decimal(20, 8)"`
	LastAvailableBalance         decimal.Decimal `gorm:"type:decimal(20, 8)"`
	Balance                      decimal.Decimal `gorm:"type:decimal(20, 8)"`
	LastBalance                  decimal.Decimal `gorm:"type:decimal(20, 8)"`
	SectorPledgeBalance          decimal.Decimal `gorm:"type:decimal(20, 8)"`
	LastSectorPledgeBalance      decimal.Decimal `gorm:"type:decimal(20, 8)"`
	LastMonthSectorPledgeBalance decimal.Decimal `gorm:"type:decimal(20, 8)"`
	VestingFunds                 decimal.Decimal `gorm:"type:decimal(20, 8)"`
	LastVestingFunds             decimal.Decimal `gorm:"type:decimal(20, 8)"`
	Height                       uint            `gorm:"type:int"`
	LastTime                     time.Time       `gorm:"type:datetime"`
	RewardValue                  decimal.Decimal `gorm:"type:decimal(20, 8)"`
	LastRewardValue              decimal.Decimal `gorm:"type:decimal(20, 8)"`
	LastMonthRewardValue         decimal.Decimal `gorm:"type:decimal(20, 8)"`
	WeightedBlocks               int             `gorm:"type:int"`
	LastWeightedBlocks           int             `gorm:"type:int"`
	LastMonthWeightedBlocks      int             `gorm:"type:int"`
	QualityAdjPower              decimal.Decimal `gorm:"type:decimal(20, 4)",有效算力`
	LastQualityAdjPower          decimal.Decimal `gorm:"type:decimal(20, 4)",有效算力`
	LastMonthQualityAdjPower     decimal.Decimal `gorm:"type:decimal(20, 4)" 月初算力`
	PowerUnit                    string          `gorm:"type:varchar(50)",算力单位`
	PowerPoint                   decimal.Decimal `gorm:"type:decimal(10,3)",算力占比`
	ControlBalance               decimal.Decimal `gorm:"type:decimal(20,8)"`
	BlocksMined24h               int             `gorm:"type:int" 24h报块数量`
	TotalRewards24h              decimal.Decimal `gorm:"type:decimal(20, 8)" 24h出块奖励金额`
	LuckyValue24h                decimal.Decimal `gorm:"type:decimal(20, 8)" 24hLucky值`
	QualityAdjPowerDelta24h      decimal.Decimal `gorm:"type:decimal(20, 4)" 24h算力增量`
	ReceiveAmount                decimal.Decimal `gorm:"type:decimal(20, 8)" 节点接收数量`
	BurnAmount                   decimal.Decimal `gorm:"type:decimal(20, 8)" 节点销毁数量`
	SendAmount                   decimal.Decimal `gorm:"type:decimal(20, 8)" 节点发送数量`
	LastReceiveAmount            decimal.Decimal `gorm:"type:decimal(20, 8)" 前一天接收数量`
	LastBurnAmount               decimal.Decimal `gorm:"type:decimal(20, 4)" 前一天销毁数量`
	LastSendAmount               decimal.Decimal `gorm:"type:decimal(20, 4)" 前一天提现数量`
	LastMonthReceiveAmount       decimal.Decimal `gorm:"type:decimal(20, 8)" 上月末节点接收数量`
	LastMonthBurnAmount          decimal.Decimal `gorm:"type:decimal(20, 4)" 上月末销毁数量`
	LastMonthSendAmount          decimal.Decimal `gorm:"type:decimal(20, 4)" 上月末提现数量`
	TimeTag                      int64           `gorm:"type:bigint" 时间标签`
}

func (table *NodesChart) TableName() string {
	return "nodes_chart"
}
