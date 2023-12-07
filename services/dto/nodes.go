package dto

import (
	"blog/models"
	"github.com/shopspring/decimal"
	"time"
)

type NodesUpdateReq struct {
	Id                  uint            `json:"id"`
	Node                string          `json:"node"`
	MsigNode            string          `gorm:"type:varchar(255)"`
	Address             string          `gorm:"type:varchar(255)"`
	MsgCount            int             `gorm:"type:int"`
	SectorType          string          `gorm:"type:varchar(50)"`
	CreateTime          time.Time       `gorm:"type:datetime"`
	AvailableBalance    decimal.Decimal `gorm:"type:decimal(20, 8)"`
	Balance             decimal.Decimal `gorm:"type:decimal(20, 8)"`
	SectorPledgeBalance decimal.Decimal `gorm:"type:decimal(20, 8)"`
	VestingFunds        decimal.Decimal `gorm:"type:decimal(20, 8)"`
	Height              uint            `gorm:"type:int"`
	Status              int             `gorm:"type:int"`
	Type                int             `gorm:"type:int"`
	LastTime            time.Time       `gorm:"type:datetime"`
	RewardValue         decimal.Decimal `gorm:"type:decimal(20, 8)"`
	QualityAdjPower     decimal.Decimal `gorm:"type:decimal(20, 4)",有效算力`
	PowerUnit           string          `gorm:"type:varchar(50)",算力单位`
	PowerPoint          decimal.Decimal `gorm:"type:decimal(10,3)",算力占比`
	PowerGrade          string          `gorm:"type:varchar(50)",算力排名`
	SectorSize          string          `gorm:"type:varchar(50)",扇区大小`
	SectorStatus        string          `gorm:"type:varchar(255)"`
	ControlAddress      string          `gorm:"type:varchar(255)"`
	ControlBalance      decimal.Decimal `gorm:"type:decimal(20,8)"`
	HasTransfer         decimal.Decimal `gorm:"type:decimal(20,8)"`
	BlocksMined         int             `gorm:"type:int" 24h报块数量`
	WeightedBlocksMined int             `gorm:"type:int" 24h出块份数`
	TotalRewards24h     decimal.Decimal `gorm:"type:decimal(20, 8)" 24h出块奖励金额`
	LuckyValue          decimal.Decimal `gorm:"type:decimal(20, 8)" 24hLucky值`
	LastDistributeTime  time.Time       `gorm:"type:datetime" 最后一次分币时间`
}

func (s *NodesUpdateReq) Generate(model *models.Nodes) {
	if s.Id == 0 {
		model.ID = 0
	}
	model.Node = s.Node
	model.MsigNode = s.MsigNode
	model.Address = s.Address
	model.HasTransfer = s.HasTransfer
	model.LastDistributeTime = s.LastDistributeTime
}
