package services

import (
	"blog/common"
	"blog/models"

	"gorm.io/gorm"
)

func SavePoolChart(poolChart *models.PoolChart) *gorm.DB {
	return common.DB.Create(&poolChart)
}
