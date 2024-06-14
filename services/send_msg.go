package services

import (
	"blog/common"
	"blog/models"
	"gorm.io/gorm"
	"log"
)

func InsertMsg(msg models.Msg) *gorm.DB {
	return common.DB.Save(&msg)
}

func UpdateSendStatus(msg models.Msg) {

	tx := common.DB.Save(&msg) //.Exec(sql)
	if tx != nil {
		log.Printf("update %s send_status success:\n", msg.Node)
	}
}
