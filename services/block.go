package services

import (
	"blog/common"
	"blog/models"
	"reflect"

	"gorm.io/gorm"
)

type BlockService struct {
}

func (bs *BlockService) Insert(block models.Block) *gorm.DB {
	return common.DB.Save(&block)
}

func (bs *BlockService) NeedToSave(lastBlock LuckyBlock, spiders []models.LuckyBlock) []models.LuckyBlock {
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
