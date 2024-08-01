package services

import (
	"blog/common"
	"blog/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BlockService struct {
}

func (bs *BlockService) Insert(block models.Block) *gorm.DB {
	return common.DB.Save(&block)
}

func (bs *BlockService) InsertMap(m map[string]interface{}) *gorm.DB {
	block := bs.MapToBlock(m)

	if block.Status == 2 {
		//发现孤块，保存预警信息
		ms := SendMsg{}
		ms.SaveMsgByType(block.Node, "节点"+block.Node+"出现孤块", fmt.Sprintf("节点%s出现孤块，高度为%s", block.Node, m["height"].(string)), models.OrphanBlock)
	}

	return bs.Insert(block)
}

func (bs *BlockService) MapToBlock(m map[string]interface{}) models.Block {
	return models.Block{
		Height:     m["height"].(int64),
		Node:       m["miner"].(string),
		Message:    m["cid"].(string),
		Status:     1,
		CreateTime: time.Now(),
	}
}
