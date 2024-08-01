package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Nodes struct {
	ID                      uint            `gorm:"primarykey"`
	Node                    string          `gorm:"type:varchar(255)"`
	MsigNode                string          `gorm:"type:varchar(255)"`
	Address                 string          `gorm:"type:varchar(255)"`
	MsgCount                int             `gorm:"type:int"`
	SectorType              string          `gorm:"type:varchar(50)"`
	CreateTime              time.Time       `gorm:"type:datetime"`
	AvailableBalance        decimal.Decimal `gorm:"type:decimal(20, 8)"`
	Balance                 decimal.Decimal `gorm:"type:decimal(20, 8)"`
	SectorPledgeBalance     decimal.Decimal `gorm:"type:decimal(20, 8)"`
	VestingFunds            decimal.Decimal `gorm:"type:decimal(20, 8)"`
	Height                  uint            `gorm:"type:int"`
	Status                  int             `gorm:"type:int"`
	Type                    int             `gorm:"type:int"`
	LastTime                time.Time       `gorm:"type:datetime"`
	LastHandTime            *time.Time      `gorm:"type:datetime"`
	RewardValue             decimal.Decimal `gorm:"type:decimal(20, 8)"`
	WeightedBlocks          int             `gorm:"type:int"`
	QualityAdjPower         decimal.Decimal `gorm:"type:decimal(20, 4)" json:"qualityAdjPower"`
	QualityPower            decimal.Decimal `gorm:"type:decimal(20, 4)" json:"qualityPower"`
	RawPower                decimal.Decimal `gorm:"type:decimal(20, 4)" json:"rawPower"`
	PowerUnit               string          `gorm:"type:varchar(50)" comment:"算力单位"`
	PowerPoint              decimal.Decimal `gorm:"type:decimal(10,3)" comment:"算力占比"`
	PowerGrade              string          `gorm:"type:varchar(50)" comment:"算力排名"`
	SectorSize              string          `gorm:"type:varchar(50)" comment:"扇区大小"`
	SectorStatus            string          `gorm:"type:varchar(255)"`
	SectorTotal             int             `gorm:"type:int"`
	SectorEffective         int             `gorm:"type:int"`
	SectorError             int             `gorm:"type:int"`
	SectorRecovering        int             `gorm:"type:int"`
	ControlAddress          string          `gorm:"type:varchar(255)"`
	ControlBalance          decimal.Decimal `gorm:"type:decimal(20,8)"`
	HasTransfer             decimal.Decimal `gorm:"type:decimal(20,8)"`
	MiningEfficiency        decimal.Decimal `gorm:"type:decimal(20,8)"`
	BlocksMined24h          int             `gorm:"type:int" comment:"24h报块数量"`
	WeightedBlocksMined24h  int             `gorm:"type:int" comment:"24h出块份数"`
	TotalRewards24h         decimal.Decimal `gorm:"type:decimal(20, 8)" comment:"24h出块奖励金额"`
	LuckyValue24h           decimal.Decimal `gorm:"type:decimal(20, 8)" comment:"24hLucky值"`
	QualityAdjPowerDelta24h decimal.Decimal `gorm:"type:decimal(20, 4)" comment:"24h算力增量"`
	MiningEfficiency7d      decimal.Decimal `gorm:"type:decimal(20,8)"`
	BlocksMined7d           int             `gorm:"type:int" comment:"7d报块数量"`
	WeightedBlocksMined7d   int             `gorm:"type:int" comment:"7d出块份数"`
	TotalRewards7d          decimal.Decimal `gorm:"type:decimal(20, 8)" comment:"7d出块奖励金额"`
	LuckyValue7d            decimal.Decimal `gorm:"type:decimal(20, 8)" comment:"7dLucky值"`
	QualityAdjPowerDelta7d  decimal.Decimal `gorm:"type:decimal(20, 4)" comment:"7d算力增量"`
	MiningEfficiency30d     decimal.Decimal `gorm:"type:decimal(20,8)"`
	BlocksMined30d          int             `gorm:"type:int" comment:"月报块数量"`
	WeightedBlocksMined30d  int             `gorm:"type:int" comment:"月出块份数"`
	TotalRewards30d         decimal.Decimal `gorm:"type:decimal(20, 8)" comment:"月出块奖励金额"`
	LuckyValue30d           decimal.Decimal `gorm:"type:decimal(20, 8)" comment:"月Lucky值"`
	QualityAdjPowerDelta30d decimal.Decimal `gorm:"type:decimal(20, 4)" comment:"月算力增量"`
	LastDistributeTime      time.Time       `gorm:"type:datetime" comment:"最后一次分币时间"`
	ReceiveAmount           decimal.Decimal `gorm:"type:decimal(20, 8)" comment:"节点接收数量"`
	BurnAmount              decimal.Decimal `gorm:"type:decimal(20, 8)" comment:"节点销毁数量"`
	SendAmount              decimal.Decimal `gorm:"type:decimal(20, 8)" comment:"节点发送数量"`
	TimeTag                 int64           `gorm:"type:bigint" comment:"时间标签"`
	TransferCount           int64           `gorm:"type:int" comment:"转账交易数"`
	RealCount               int64           `gorm:"type:int" comment:"实际交易数量"`
	DeptId                  int             `json:"deptId"`
	SyncStatus              string          `json:"syncStatus" gorm:"type:varchar(50);comment:同步状态"`
}

func (table *Nodes) TableName() string {
	return "fil_nodes"
}
