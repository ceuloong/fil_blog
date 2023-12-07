package services

import (
	"blog/common"
	"blog/models"
	"gorm.io/gorm"
	"time"
)

func InsertNodesChart(nodesChart models.NodesChart) *gorm.DB {
	return common.DB.Create(&nodesChart)
}

func UpdateNodesChart(nodesChart models.NodesChart) *gorm.DB {
	return common.DB.Save(&nodesChart)
}

func GetLastOneByTime(node models.Nodes, time time.Time) models.NodesChart {
	var lastOne models.NodesChart
	common.DB.Model(&models.NodesChart{}).Where("TO_DAYS(last_time) = TO_DAYS(?) AND node = ?", time, node.Node).Order("last_time DESC").First(&lastOne)
	return lastOne
}

func SaveNodesChart(nodes models.Nodes) {
	currentTime := time.Now()
	lastDay := currentTime.AddDate(0, 0, -1)
	lastOne := GetLastOneByTime(nodes, lastDay)

	lastMonthLastDay := currentTime.AddDate(0, 0, -currentTime.Day())
	lastMonthLastOne := GetLastOneByTime(nodes, lastMonthLastDay)

	nodesChart := models.NodesChart{
		Node:                         nodes.Node,
		AvailableBalance:             nodes.AvailableBalance,
		Balance:                      nodes.Balance,
		SectorPledgeBalance:          nodes.SectorPledgeBalance,
		VestingFunds:                 nodes.VestingFunds,
		Height:                       nodes.Height,
		LastTime:                     nodes.LastTime,
		RewardValue:                  nodes.RewardValue,
		WeightedBlocks:               nodes.WeightedBlocks,
		QualityAdjPower:              nodes.QualityAdjPower,
		PowerUnit:                    nodes.PowerUnit,
		PowerPoint:                   nodes.PowerPoint,
		ControlBalance:               nodes.ControlBalance,
		BlocksMined24h:               nodes.BlocksMined24h,
		TotalRewards24h:              nodes.TotalRewards24h,
		LuckyValue24h:                nodes.LuckyValue24h,
		QualityAdjPowerDelta24h:      nodes.QualityAdjPowerDelta24h,
		ReceiveAmount:                nodes.ReceiveAmount,
		BurnAmount:                   nodes.BurnAmount,
		SendAmount:                   nodes.SendAmount,
		LastAvailableBalance:         lastOne.AvailableBalance,
		LastBalance:                  lastOne.Balance,
		LastSectorPledgeBalance:      lastOne.SectorPledgeBalance,
		LastVestingFunds:             lastOne.VestingFunds,
		LastRewardValue:              lastOne.RewardValue,
		LastWeightedBlocks:           lastOne.WeightedBlocks,
		LastQualityAdjPower:          lastOne.QualityAdjPower,
		LastReceiveAmount:            lastOne.ReceiveAmount,
		LastBurnAmount:               lastOne.BurnAmount,
		LastSendAmount:               lastOne.SendAmount,
		LastMonthSectorPledgeBalance: lastMonthLastOne.SectorPledgeBalance,
		LastMonthRewardValue:         lastMonthLastOne.RewardValue,
		LastMonthWeightedBlocks:      lastMonthLastOne.WeightedBlocks,
		LastMonthQualityAdjPower:     lastMonthLastOne.QualityAdjPower,
		LastMonthReceiveAmount:       lastMonthLastOne.ReceiveAmount,
		LastMonthBurnAmount:          lastMonthLastOne.BurnAmount,
		LastMonthSendAmount:          lastMonthLastOne.SendAmount,
		TimeTag:                      nodes.TimeTag,
	}

	InsertNodesChart(nodesChart)
}
