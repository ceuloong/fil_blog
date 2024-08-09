package dto

import (
	"blog/models"
	"blog/utils"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
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
	QualityAdjPower     decimal.Decimal `gorm:"type:decimal(20, 4)" commnet:"有效算力"`
	PowerUnit           string          `gorm:"type:varchar(50)" commnet:"算力单位"`
	PowerPoint          decimal.Decimal `gorm:"type:decimal(10,3)" commnet:"算力占比"`
	PowerGrade          string          `gorm:"type:varchar(50)" commnet:"算力排名"`
	SectorSize          string          `gorm:"type:varchar(50)" commnet:"扇区大小"`
	SectorStatus        string          `gorm:"type:varchar(255)"`
	ControlAddress      string          `gorm:"type:varchar(255)"`
	ControlBalance      decimal.Decimal `gorm:"type:decimal(20,8)"`
	HasTransfer         decimal.Decimal `gorm:"type:decimal(20,8)"`
	BlocksMined         int             `gorm:"type:int" commnet:"24h报块数量"`
	WeightedBlocksMined int             `gorm:"type:int" commnet:"24h出块份数"`
	TotalRewards24h     decimal.Decimal `gorm:"type:decimal(20, 8)" commnet:"24h出块奖励金额"`
	LuckyValue          decimal.Decimal `gorm:"type:decimal(20, 8)" commnet:"24hLucky值"`
	LastDistributeTime  time.Time       `gorm:"type:datetime" commnet:"最后一次分币时间"`
}

func (ms *MinerStatus) Generate(model *models.Nodes) {
	model.SyncStatus = ms.Chain
	model.QualityAdjPower = utils.DecimalValue(ms.Power)
	model.RawPower = utils.DecimalValue(ms.Raw)
	model.Balance = utils.DecimalValue(ms.Balance)
	model.AvailableBalance = utils.DecimalValue(ms.Available)
	model.SectorPledgeBalance = utils.DecimalValue(ms.Pledge)
	model.VestingFunds = utils.DecimalValue(ms.Vesting)
	model.Beneficiary = ms.Beneficiary
	total, _ := strconv.Atoi(ms.SectorsTotal)
	model.SectorTotal = total
}

type MinerStatus struct {
	StartTime    string
	Chain        string
	Miner        string
	Power        string
	Raw          string
	Balance      string
	Pledge       string
	Vesting      string
	Available    string
	Beneficiary  string
	SectorsTotal string
	SectorSize   string
}
