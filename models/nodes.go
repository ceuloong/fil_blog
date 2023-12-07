package models

import (
	"github.com/shopspring/decimal"
	"time"
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
	RewardValue             decimal.Decimal `gorm:"type:decimal(20, 8)"`
	WeightedBlocks          int             `gorm:"type:int"`
	QualityAdjPower         decimal.Decimal `gorm:"type:decimal(20, 4)",有效算力`
	PowerUnit               string          `gorm:"type:varchar(50)",算力单位`
	PowerPoint              decimal.Decimal `gorm:"type:decimal(10,3)",算力占比`
	PowerGrade              string          `gorm:"type:varchar(50)",算力排名`
	SectorSize              string          `gorm:"type:varchar(50)",扇区大小`
	SectorStatus            string          `gorm:"type:varchar(255)"`
	SectorTotal             int             `gorm:"type:int"`
	SectorEffective         int             `gorm:"type:int"`
	SectorError             int             `gorm:"type:int"`
	SectorRecovering        int             `gorm:"type:int"`
	ControlAddress          string          `gorm:"type:varchar(255)"`
	ControlBalance          decimal.Decimal `gorm:"type:decimal(20,8)"`
	HasTransfer             decimal.Decimal `gorm:"type:decimal(20,8)"`
	BlocksMined24h          int             `gorm:"type:int" 24h报块数量`
	WeightedBlocksMined24h  int             `gorm:"type:int" 24h出块份数`
	TotalRewards24h         decimal.Decimal `gorm:"type:decimal(20, 8)" 24h出块奖励金额`
	LuckyValue24h           decimal.Decimal `gorm:"type:decimal(20, 8)" 24hLucky值`
	QualityAdjPowerDelta24h decimal.Decimal `gorm:"type:decimal(20, 4)" 24h算力增量`
	BlocksMined7d           int             `gorm:"type:int" 7d报块数量`
	WeightedBlocksMined7d   int             `gorm:"type:int" 7d出块份数`
	TotalRewards7d          decimal.Decimal `gorm:"type:decimal(20, 8)" 7d出块奖励金额`
	LuckyValue7d            decimal.Decimal `gorm:"type:decimal(20, 8)" 7dLucky值`
	QualityAdjPowerDelta7d  decimal.Decimal `gorm:"type:decimal(20, 4)" 7d算力增量`
	BlocksMined30d          int             `gorm:"type:int" 月报块数量`
	WeightedBlocksMined30d  int             `gorm:"type:int" 月出块份数`
	TotalRewards30d         decimal.Decimal `gorm:"type:decimal(20, 8)" 月出块奖励金额`
	LuckyValue30d           decimal.Decimal `gorm:"type:decimal(20, 8)" 月Lucky值`
	QualityAdjPowerDelta30d decimal.Decimal `gorm:"type:decimal(20, 4)" 月算力增量`
	LastDistributeTime      time.Time       `gorm:"type:datetime" 最后一次分币时间`
	ReceiveAmount           decimal.Decimal `gorm:"type:decimal(20, 8)" 节点接收数量`
	BurnAmount              decimal.Decimal `gorm:"type:decimal(20, 8)" 节点销毁数量`
	SendAmount              decimal.Decimal `gorm:"type:decimal(20, 8)" 节点发送数量`
	TimeTag                 int64           `gorm:"type:bigint" 时间标签`
	TransferCount           int64           `gorm:"type:int" 转账交易数`
}

func (table *Nodes) TableName() string {
	return "fil_nodes"
}
