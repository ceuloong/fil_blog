package services

import (
	"blog/common"
	"blog/models"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func SaveNode(nodes []models.Nodes) *gorm.DB {
	return common.DB.CreateInBatches(&nodes, 100)
}

// FindAllNode 根据返回的数组判断是否已存在
func FindAllNode(node string) []models.Nodes {
	var db = common.DB
	var nodes []models.Nodes
	//Where("msig_node = ?", node).
	tx := db.Model(&models.Nodes{})
	if len(node) > 0 {
		tx.Where("status > 0 AND node = ?", node)
	} else {
		tx.Where("status > 0")
	}
	//tx.Where("balance > 0")
	tx.Order("id").Find(&nodes)

	return nodes
}

func UpdateRealCount(realCount int, node string) {
	sql := fmt.Sprintf("UPDATE fil_nodes SET real_count=%d where node='%s'", realCount, node)
	tx := common.DB.Exec(sql)
	if tx != nil {
		log.Printf("update %s real_count success:\n", node)
	}
}

func UpdateTransferCount(transferCount int64, realCount int, node string) {
	sql := fmt.Sprintf("UPDATE fil_nodes SET transfer_count=%d, real_count=%d where node='%s'", transferCount, realCount, node)
	tx := common.DB.Exec(sql)
	if tx != nil {
		log.Printf("update %s success:\n", node)
	}
}

func UpdateNode(node models.Nodes) {
	//sql := fmt.Sprintf(
	//	"UPDATE fil_nodes SET address='%s', msg_count=%d, create_time='%s', available_balance=%s, balance=%s, sector_pledge_balance='%s', vesting_funds='%s', height=%d, last_time='%s', reward_value='%s', quality_adj_power='%s', power_unit='%s', power_point='%s', power_grade='%s', sector_size='%s', sector_status='%s', control_address='%s', control_balance='%s', blocks_mined=%d, weighted_blocks_mined=%d, total_rewards24h=%s, lucky_value=%s, has_transfer=%s where node='%s'",
	//	node.Address, node.MsgCount, timestampToTime(node.CreateTime.Unix()), node.AvailableBalance, node.Balance, node.SectorPledgeBalance,
	//	node.VestingFunds, node.Height, timestampToTime(node.LastTime.Unix()), node.RewardValue, node.QualityAdjPower, node.PowerUnit, node.PowerPoint,
	//	node.PowerGrade, node.SectorSize, node.SectorStatus, node.ControlAddress, node.ControlBalance, node.BlocksMined, node.WeightedBlocksMined, node.TotalRewards24h, node.LuckyValue, node.HasTransfer, node.Node,
	//)
	tx := common.DB.Save(&node) //.Exec(sql)
	if tx != nil {
		log.Printf("update %s success:\n", node.Node)
	}
}

// UpdateNodeHeight 更新节点高度
func UpdateNodeHeight(height int64, node string) {
	sql := fmt.Sprintf("UPDATE fil_nodes SET height=%d where node='%s'", height, node)
	tx := common.DB.Exec(sql)
	if tx != nil {
		log.Printf("update %s height success:\n", node)
	}
}

// UpdateSyncStatus 更新同步状态
func UpdateSyncStatus(sync string, node string) {
	sql := fmt.Sprintf("UPDATE fil_nodes SET sync_status=%s where node='%s'", sync, node)
	tx := common.DB.Exec(sql)
	if tx != nil {
		log.Printf("update %s sync_status success:\n", node)
	}
}

func timestampToTime(t int64) string {
	timeTemplate1 := "2006-01-02 15:04:05"
	timeStr := time.Unix(t, 0).Format(timeTemplate1)
	return timeStr
}

type RewardValue struct {
	Amount decimal.Decimal
}

func SumReward(nodes models.Nodes, timeTag int64) float64 {
	var amount float64
	//sql := fmt.Sprintf("SElECT SUM(reward_value) as amount FROM lucky_block WHERE node = '%s' AND type = 'send' AND date > '%s'", nodes.Node, timestampToTime(nodes.LastDistributeTime.Unix()))
	//common.DB.Raw(sql, nodes.Node, timestampToTime(nodes.LastDistributeTime.Unix())).Pluck("SUM(reward_value) as amount", &totalAmount)
	//common.DB.Exec(sql).Pluck("COALESCE(SUM(reward_value), 0) as amount", &totalAmount)
	if timeTag > 0 {
		common.DB.Model(&models.LuckyBlock{}).Where("node = ? AND type = 'send' AND time_tag = ?", nodes.Node, timeTag).Pluck("COALESCE(SUM(reward_value), 0) as amount", &amount)
	} else {
		common.DB.Model(&models.LuckyBlock{}).Where("node = ? AND type = 'send' AND date > ?", nodes.Node, timestampToTime(nodes.LastDistributeTime.Unix())).Pluck("COALESCE(SUM(reward_value), 0) as amount", &amount)
	}

	return math.Abs(amount)
}

type Result []struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func SumValueByType(node string, timeTag int64) map[string]decimal.Decimal {
	var result Result
	var result2 Result
	// SELECT node, type, SUM(reward_value) as amount FROM lucky_block WHERE node = 'f0492009' GROUP BY type
	if timeTag > 0 {
		common.DB.Model(&models.LuckyBlock{}).Select("type, SUM(reward_value) as amount").Where("node = ? AND time_tag = ?", node, timeTag).Group("type").Find(&result)
		//common.DB.Model(&models.LuckyBlockBak{}).Select("type, SUM(reward_value) as amount").Where("node = ? AND time_tag = ?", node, timeTag).Group("type").Find(&result2)
	} else {
		common.DB.Model(&models.LuckyBlock{}).Select("type, SUM(reward_value) as amount").Where("node = ? OR pid_node = ?", node, node).Group("type").Find(&result)
		common.DB.Model(&models.LuckyBlockBak{}).Select("type, SUM(reward_value) as amount").Where("node = ?", node).Group("type").Find(&result2)
	}

	var m = make(map[string]decimal.Decimal)
	for _, rel := range result {
		m[rel.Type] = decimal.NewFromFloat(math.Abs(rel.Amount))
	}
	for _, rel := range result2 {
		m[rel.Type] = m[rel.Type].Add(decimal.NewFromFloat(math.Abs(rel.Amount)))
	}

	return m
}

func SumValueByTimeGroupType(node string, time time.Time) map[string]decimal.Decimal {
	var result Result
	//var result2 Result
	common.DB.Model(&models.LuckyBlock{}).Select("type, SUM(reward_value) as amount").Where("node = ? AND date > ?", node, time).Group("type").Find(&result)
	//common.DB.Model(&models.LuckyBlockBak{}).Select("type, SUM(reward_value) as amount").Where("node = ?", node).Group("type").Find(&result2)

	var m = make(map[string]decimal.Decimal)
	for _, rel := range result {
		m[rel.Type] = decimal.NewFromFloat(math.Abs(rel.Amount))
	}
	/*for _, rel := range result2 {
		m[rel.Type] = m[rel.Type].Add(decimal.NewFromFloat(math.Abs(rel.Amount)))
	}*/

	return m
}

func SumPidNodeByTimeGroupType(node string, time time.Time) map[string]decimal.Decimal {
	var result Result
	common.DB.Model(&models.LuckyBlock{}).Select("type, SUM(reward_value) as amount").Where("pid_node = ? AND date > ?", node, time).Group("type").Find(&result)
	if len(result) == 0 {
		//common.DB.Model(&models.LuckyBlockBak{}).Select("type, SUM(reward_value) as amount").Where("pid_node = ? AND date < ?", node, time).Group("type").Find(&result)
	}

	var m = make(map[string]decimal.Decimal)
	for _, rel := range result {
		m[rel.Type] = decimal.NewFromFloat(math.Abs(rel.Amount))
	}

	return m
}

// CountByNodeTime 根据返回的数组判断是否已存在
func CountByNodeTime(node string, time time.Time) int64 {
	var db = common.DB

	var count int64
	err := db.Model(&models.LuckyBlock{}).Where("node = ? AND type = 'reward' AND date > ?", node, time).Count(&count).Error
	if err != nil {
		log.Printf("FilNodesService GetPage error:%s \r\n", err)
		return 0
	}
	return count
}
