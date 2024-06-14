package client

import (
	"blog/models"
	"blog/services"
	"log"
	"time"
)

type Monitor struct {
}

func (m Monitor) UpdateNodes(nodeParam string, timeTag int64) {
	nodes := services.FindAllNode(nodeParam)
	var hasPowerCount int
	var noPowerCount int

	var savePool bool

	for _, n := range nodes {
		savePool = true

		n.TimeTag = timeTag
		services.UpdateNode(n)
		// 保存图表数据
		services.SaveNodesChart(n)
		time.Sleep(15 * time.Second)

		if n.QualityAdjPower.IsZero() {
			noPowerCount++
		} else {
			hasPowerCount++
		}

	}

	// 保存矿池图表数据
	if savePool {
		log.Printf("一共更新的 %d 个节点，其中有算力的节点 %d 个, 算力为0的节点 %d 个。\n", len(nodes), hasPowerCount, noPowerCount)
	} else {
		log.Printf("没有需要更新的节点。")
	}
}

func (m Monitor) SavePoolChart() {
	nodes := services.FindAllNode("")

	var pChartMap = make(map[int64]models.PoolChart)
	poolChart := pChartMap[0]

	for _, n := range nodes {
		poolChart.Balance = poolChart.Balance.Add(n.Balance)
		poolChart.AvailableBalance = poolChart.AvailableBalance.Add(n.AvailableBalance)
		poolChart.SectorPledgeBalance = poolChart.SectorPledgeBalance.Add(n.SectorPledgeBalance)
		poolChart.VestingFunds = poolChart.VestingFunds.Add(n.VestingFunds)
		poolChart.QualityAdjPower = poolChart.QualityAdjPower.Add(n.QualityAdjPower)
		poolChart.PowerPoint = poolChart.PowerPoint.Add(n.PowerPoint)
		poolChart.ControlBalance = poolChart.ControlBalance.Add(n.ControlBalance)
		poolChart.RewardValue = poolChart.RewardValue.Add(n.RewardValue)

		//nodesChart := services.GetNodesChart(n)

	}
	poolChart.LastTime = time.Now()
	poolChart.PowerUnit = "PiB"
	services.SavePoolChart(&poolChart)
}
