package services

import (
	"blog/common"
	"blog/models"

	"gorm.io/gorm"
)

func InsertStats(blockstats []models.BlockStats) *gorm.DB {
	return common.DB.CreateInBatches(&blockstats, 100)
}

// GetLastOne 根据返回的数组判断是否已存在
func GetLastOne(node string) models.BlockStats {
	var db = common.DB
	var lastBlock models.BlockStats
	db.Model(&models.BlockStats{}).Where("node = ?", node).Order("id desc").Limit(1).Find(&lastBlock)

	return lastBlock
}
