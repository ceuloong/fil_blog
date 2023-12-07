package httpcorrect

import (
	"encoding/json"
	"fmt"
	"log"
)

// MiningDetail 返回的结构体
type MiningDetail struct {
	BlocksMined          int     `json:"blocksMined"`
	WeightedBlocksMined  int     `json:"weightedBlocksMined"`
	TotalRewards         string  `json:"totalRewards"`
	NetworkTotalRewards  string  `json:"networkTotalRewards"`
	LuckyValue           float64 `json:"luckyValue"`
	QualityAdjPowerDelta string  `json:"qualityAdjPowerDelta"`
}

// MiningStats 获取节点状态数据
func MiningStats(node string) MiningDetail {
	return MiningStatsCycle(node, "24h")
}

// MiningStatsCycle 获取节点状态数据
func MiningStatsCycle(node string, cycle string) MiningDetail {
	url := fmt.Sprintf("https://filfox.info/api/v1/address/%s/mining-stats?duration=%s", node, cycle)
	bodyText, err := RequestUrl(url)
	if err != nil {
		log.Fatal(err)
	}
	var result MiningDetail
	_ = json.Unmarshal(bodyText, &result) //byte to json

	return result
}
