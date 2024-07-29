package services

import (
	"blog/common"
	"blog/models"
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
