package services

import (
	"blog/common"
	"blog/models"
	"log"
	"time"

	"gorm.io/gorm"
)

type SendMsg struct {
}

func (m *SendMsg) InsertMsg(msg models.Msg) *gorm.DB {
	return common.DB.Save(&msg)
}

func (m *SendMsg) UpdateSendStatus(msg models.Msg) {

	tx := common.DB.Save(&msg) //.Exec(sql)
	if tx != nil {
		log.Printf("update %s send_status success:\n", msg.Node)
	}
}

func (m *SendMsg) SaveMsgByType(node string, title string, content string, send_type models.SendType) *gorm.DB {
	ms := SendMsg{}
	msg := models.Msg{
		Node:       node,
		Title:      title,
		Content:    content,
		CreateTime: time.Now(),
		Type:       send_type,
		SendStatus: 0,
	}

	// switch send_type {
	// case models.SectorsError:
	// 	msg.Title = "扇区错误"
	// case models.HeightDelay:
	// 	msg.Title = "高度延迟"
	// case models.LuckyLow:
	// 	msg.Title = "幸运值过低"
	// case models.OrphanBlock:
	// 	msg.Title = "孤块"
	// 	msg.Content = fmt.Sprintf("节点%s出现孤块，高度为%s", node, content)
	// default:
	// 	msg.Title = "未知"
	// }

	return ms.InsertMsg(msg)
}
