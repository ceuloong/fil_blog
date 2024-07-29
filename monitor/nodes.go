package client

import (
	"blog/models"
	"blog/services"
	"fmt"
	"log"
	"strings"
	"time"
)

type Monitor struct {
}

type Block struct {
	Height  int64  `json:"height"`
	Miner   string `json:"miner"`
	Message string `json:"cid"`
}

type MinerStatus struct {
	StartTime    string
	Chain        string
	Miner        string
	Power        string
	Raw          string
	Balance      string
	Pledge       string
	Vesting      string
	Available    string
	Beneficiary  string
	SectorsTotal string
}

func ReadToBean(content string) MinerStatus {
	list := strings.Split(content, "\\n")
	m := map[string]string{}
	for i, v := range list {
		fmt.Printf("index: %d, value: %s\n", i, v)
		v = strings.TrimSpace(v)
		kv := strings.Split(v, ":")
		if _, ok := m[kv[0]]; !ok && len(kv) > 1 {
			m[kv[0]] = strings.TrimSpace(kv[1])
		}
	}
	for k, v := range m {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
	ms := MinerStatus{
		StartTime:    m["StartTime"],
		Chain:        m["Chain"],
		Miner:        m["Miner"],
		Power:        strings.Split(m["Power"], "/")[0],
		Raw:          strings.Split(m["Raw"], "/")[0],
		Balance:      strings.Split(m["Miner Balance"], " ")[0],
		Pledge:       strings.Split(m["Pledge"], " ")[0],
		Vesting:      strings.Split(m["Vesting"], " ")[0],
		Available:    strings.Split(m["Available"], " ")[0],
		Beneficiary:  m["Beneficiary"],
		SectorsTotal: m["Total"],
	}
	fmt.Printf("MinerStatus: %v\n", ms)
	return ms
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
