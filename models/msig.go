package models

import "github.com/shopspring/decimal"

type Msig struct {
	ID            uint            `gorm:"primarykey"`
	Address       string          `gorm:"type:varchar(50)"`
	RobustAddress string          `gorm:"type:varchar(255)"`
	Balance       decimal.Decimal `gorm:"type:decimal(20, 8)"`
}

func (table *Msig) TableName() string {
	return "msig"
}
