package services

import (
	"blog/common"
	"blog/models"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type BlockService struct {
}

func (bs *BlockService) Insert(block models.Block) *gorm.DB {
	return common.DB.Create(&block)
}

func (bs *BlockService) InsertMap(m map[string]interface{}, status int) *gorm.DB {
	block := bs.MapToBlock(m)
	block.Status = status
	if block.Status == 2 {
		//发现孤块，保存预警信息
		ms := SendMsg{}
		ms.SaveMsgByType(block.Node, "节点"+block.Node+"出现孤块", fmt.Sprintf("节点%s出现孤块，高度为%d", block.Node, block.Height), models.OrphanBlock)
	}

	return common.DB.Create(&block)
}

func (bs *BlockService) MapToBlock(m map[string]interface{}) models.Block {
	return models.Block{
		Height:      int64(m["height"].(float64)),
		Node:        m["miner"].(string),
		Message:     m["cid"].(string),
		RewardValue: decimal.Zero,
		CreateTime:  time.Now(),
		BlockTime:   time.Now(),
	}
}
