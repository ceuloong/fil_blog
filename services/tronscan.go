package services

import (
	"blog/common"
	"blog/models"
	"gorm.io/gorm"
)

func InsertTron(tron []models.TronScan) *gorm.DB {
	return common.DB.CreateInBatches(&tron, 100)
}

// FindAllTronByLevel 根据返回的数组判断是否已存在
func FindAllTronByLevel(level int) []models.TronScan {
	var db = common.DB
	var trons []models.TronScan
	tx := db.Model(&models.TronScan{}).Select("to_addr, COALESCE(SUM(quantity), 0) as quantity")
	tx.Where("level = ?", level)
	if level > 1 {
		tx.Where("date > ? AND (to_tag is NULL OR to_tag = '') AND quantity > 50000", "2022-01-19 06:00:00")
	}
	tx.Group("to_addr").Find(&trons)
	//tx.Order("id").Find(&trons)

	return trons
}

// FindAllTronByTxId 根据返回的数组判断是否已存在
func FindAllTronByTxId() map[string]string {
	var db = common.DB
	var trons []models.TronScan
	tx := db.Model(&models.TronScan{}).Select("tx_id")
	tx.Group("tx_id").Find(&trons)
	//tx.Order("id").Find(&trons)

	var m = make(map[string]string)
	for _, tron := range trons {
		m[tron.TxId] = tron.TxId
	}

	return m
}
