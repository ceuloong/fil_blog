package services

import (
	"blog/common"
	"blog/models"
	"gorm.io/gorm"
	"log"
	"reflect"
)

func Insert(block []models.LuckyBlock) *gorm.DB {
	return common.DB.CreateInBatches(&block, 100)
}

type LuckyBlock struct {
	Node     string
	Height   int64
	NodeFrom string
	NodeTo   string
	Message  string
	Type     string
}

// FindLastByNode 根据返回的数组判断是否已存在
func FindLastByNode(node string) LuckyBlock {
	var db = common.DB
	var lastBlock LuckyBlock
	db.Model(&models.LuckyBlock{}).Where("node = ?", node).Order("id desc").Limit(1).Find(&lastBlock)
	if lastBlock == (LuckyBlock{}) {
		db.Model(&models.LuckyBlockBak{}).Where("node = ?", node).Order("id desc").Limit(1).Find(&lastBlock)
	}

	return lastBlock
}

// CountByNode 根据返回的数组判断是否已存在
func (e *LuckyBlock) CountByNode(node string, count *int64) error {
	var db = common.DB

	err := db.Model(&models.LuckyBlock{}).Where("node = ?", node).Count(count).Error
	if err != nil {
		log.Printf("FilNodesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// CountByNodeBak 根据返回的数组判断是否已存在
func (e *LuckyBlock) CountByNodeBak(node string, count *int64) error {
	var db = common.DB

	err := db.Model(&models.LuckyBlockBak{}).Where("node = ?", node).Count(count).Error
	if err != nil {
		log.Printf("FilNodesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// CountByNodeTimeTag 根据返回的数组判断是否已存在
func (e *LuckyBlock) CountByNodeTimeTag(node string, timeTag int64, count *int64) error {
	var db = common.DB

	err := db.Model(&models.LuckyBlock{}).Where("node = ? AND time_tag = ?", node, timeTag).Count(count).Error
	if err != nil {
		log.Printf("FilNodesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

func NeedToSave(lastBlock LuckyBlock, spiders []models.LuckyBlock) []models.LuckyBlock {
	// 已保存的最新区块的高度大于当前抓取的最新记录，说明已保存过，返回
	if lastBlock.Height > spiders[len(spiders)-1].Height {
		return nil
	}

	if lastBlock.Height == spiders[len(spiders)-1].Height {
		tmp := spiderToTmp(spiders[len(spiders)-1])
		if reflect.DeepEqual(lastBlock, tmp) {
			return nil
		}
	}

	var needSaves []models.LuckyBlock
	needAddSameHeight := false
	// spiders 是按区块高度升级排列，所以正序比较，找出需要保存的记录
	for j := 0; j < len(spiders); j++ {
		if lastBlock.Height > spiders[j].Height {
			continue
		}

		if lastBlock.Height == spiders[j].Height {
			tmp := spiderToTmp(spiders[j])
			if reflect.DeepEqual(lastBlock, tmp) {
				needAddSameHeight = true
				continue
			}
			if needAddSameHeight {
				needSaves = append(needSaves, spiders[j])
			}
		} else if lastBlock.Height < spiders[j].Height {
			needSaves = append(needSaves, spiders[j])
		}
	}

	return needSaves
}

// 封装对象成临时
func spiderToTmp(spider models.LuckyBlock) LuckyBlock {
	var tmp LuckyBlock
	tmp.Node = spider.Node
	tmp.Height = spider.Height
	tmp.NodeFrom = spider.NodeFrom
	tmp.NodeTo = spider.NodeTo
	tmp.Message = spider.Message
	tmp.Type = spider.Type

	return tmp
}
