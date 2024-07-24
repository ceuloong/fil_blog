package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type PoolChart struct {
	ID                  uint            `gorm:"primarykey"`
	AvailableBalance    decimal.Decimal `gorm:"type:decimal(20, 8)"`
	Balance             decimal.Decimal `gorm:"type:decimal(20, 8)"`
	SectorPledgeBalance decimal.Decimal `gorm:"type:decimal(20, 8)"`
	VestingFunds        decimal.Decimal `gorm:"type:decimal(20, 8)"`
	LastTime            time.Time       `gorm:"type:datetime"`
	RewardValue         decimal.Decimal `gorm:"type:decimal(20, 8)"`
	QualityAdjPower     decimal.Decimal `gorm:"type:decimal(20, 4)",有效算力`
	PowerUnit           string          `gorm:"type:varchar(50)",算力单位`
	PowerPoint          decimal.Decimal `gorm:"type:decimal(10,3)",算力占比`
	ControlBalance      decimal.Decimal `gorm:"type:decimal(20,8)"`
	DeptId              int             `gorm:"type:int"`
}

func (table *PoolChart) TableName() string {
	return "pool_chart"
}
