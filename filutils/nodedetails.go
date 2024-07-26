package filutils

import (
	"blog/models"
	"blog/services"
	"blog/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type Account struct {
	Data struct {
		Height                 uint            `json:"height,omitempty"`
		Miner                  string          `json:"miner,omitempty"`
		RobustAddress          string          `json:"robustAddress,omitempty"`
		LastTime               int64           `json:"lastTime,omitempty"`
		ActorType              string          `json:"actorType,omitempty"`
		Balance                decimal.Decimal `json:"balance"`
		Available              decimal.Decimal `json:"available"`
		SectorsPledge          decimal.Decimal `json:"sectorsPledge"`
		LockedFunds            decimal.Decimal `json:"lockedFunds"`
		BlockReward            decimal.Decimal `json:"blockReward"`
		Blocks                 int             `json:"blocks,omitempty"`
		WinCount               int             `json:"winCount,omitempty"`
		MsgCount               int             `json:"msgCount,omitempty"`
		QualityPower           decimal.Decimal `json:"qualityPower"`
		RawPower               decimal.Decimal `json:"rawPower"`
		QualityPowerPercent    decimal.Decimal `json:"qualityPowerPercent"`
		QualityPowerPercentStr string          `json:"qualityPowerPercentStr,omitempty"`
		SectorSizeStr          string          `json:"sectorSizeStr,omitempty"`
		Owner                  string          `json:"owner,omitempty"`
		Worker                 string          `json:"worker,omitempty"`
		AllSectorCount         int             `json:"allSectorCount,omitempty"`
		LiveCount              int             `json:"liveCount,omitempty"`
		ActiveCount            int             `json:"activeCount,omitempty"`
		FaultCount             int             `json:"faultCount,omitempty"`
		RecoveryCount          int             `json:"recoveryCount,omitempty"`
		TerminatedCount        int             `json:"terminatedCount,omitempty"`
		CreateTime             string          `json:"createTime"`
		PowerRank              int             `json:"powerRank,omitempty"`
		BalanceStr             string          `json:"balanceStr,omitempty"`
		QualityPowerStr        string          `json:"qualityPowerStr,omitempty"`
		RawPowerStr            string          `json:"rawPowerStr,omitempty"`
		AvailableStr           string          `json:"availableStr,omitempty"`
		SectorsPledgeStr       string          `json:"sectorsPledgeStr,omitempty"`
		LockedFundsStr         string          `json:"lockedFundsStr,omitempty"`
		BlockRewardStr         string          `json:"blockRewardStr,omitempty"`
	} `json:"data"`
}

// NodeDetails 获取账户详情数组
func NodeDetails(node string) Account {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.filutils.com/api/v2/miner/%s", node)
	reqSpider, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	reqSpider.Header.Set("Accept", "*/*")
	reqSpider.Header.Set("Locale", "zh")
	reqSpider.Header.Set("Content-Type", "application/json;charset=UTF-8")
	reqSpider.Header.Set("sec-ch-ua", "\"Chromium\";v=\"126\", \"Google Chrome\";v=\"126\", \"Not:A-Brand\";v=\"99\"")
	reqSpider.Header.Set("sec-ch-ua-mobile", "?0")
	reqSpider.Header.Set("sec-ch-ua-platform", "macOS")
	reqSpider.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	respSpider, err := client.Do(reqSpider)
	if err != nil {
		log.Fatal(err)
	}
	defer respSpider.Body.Close()
	bodyText, err := ioutil.ReadAll(respSpider.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result Account
	_ = json.Unmarshal(bodyText, &result) //byte to json

	return result
}

type MineData struct {
	Data struct {
		Miner                 string          `json:"miner,omitempty"`
		QualityPowerGrowth    decimal.Decimal `json:"qualityPowerGrowth" comment:"24h算力增量"`
		MiningEfficiencyFloat float64         `json:"miningEfficiencyFloat" comment:"产出效率"`
		Blocks                int             `json:"blocks,omitempty" comment:"24h出块数量"`
		BlockReward           decimal.Decimal `json:"blockReward" comment:"24h出块奖励"`
		LuckyValue            float64         `json:"luckyValue" comment:"24hLucky值"`
	} `json:"data"`
}

// NodeDetails 获取账户详情数组
func MiningStats(node string, tp string) MineData {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.filutils.com/api/v2/miner/%s/mining-data", node)
	jsonStr := []byte(`{ "statsType": "` + tp + `"}`)
	reqSpider, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	reqSpider.Header.Set("Accept", "*/*")
	reqSpider.Header.Set("Locale", "zh")
	reqSpider.Header.Set("Content-Type", "application/json;charset=UTF-8")
	reqSpider.Header.Set("sec-ch-ua", "\"Chromium\";v=\"126\", \"Google Chrome\";v=\"126\", \"Not:A-Brand\";v=\"99\"")
	reqSpider.Header.Set("sec-ch-ua-mobile", "?0")
	reqSpider.Header.Set("sec-ch-ua-platform", "macOS")
	reqSpider.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	respSpider, err := client.Do(reqSpider)
	if err != nil {
		log.Fatal(err)
	}
	defer respSpider.Body.Close()
	bodyText, err := ioutil.ReadAll(respSpider.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result MineData
	_ = json.Unmarshal(bodyText, &result) //byte to json

	return result
}

// GetNodeDetailByAddress 根据节点名称查询节点详细信息
func GetNodeDetailByAddress(nodes models.Nodes) models.Nodes {
	node := nodes.Node
	result := NodeDetails(node)
	account := result.Data

	nodes.Balance = utils.DecimalValueFromLong(account.Balance)
	nodes.AvailableBalance = utils.DecimalValueFromLong(account.Available)
	nodes.SectorPledgeBalance = utils.DecimalValueFromLong(account.SectorsPledge)
	nodes.VestingFunds = utils.DecimalValueFromLong(account.LockedFunds)

	nodes.QualityAdjPower = utils.DecimalDiv1024xnValue(account.QualityPower.String(), 5)
	nodes.PowerUnit = "PiB"

	point := account.QualityPowerPercentStr
	nodes.PowerPoint = utils.DecimalValue(strings.Split(point, "%")[0])
	nodes.PowerGrade = strconv.Itoa(account.PowerRank)
	nodes.WeightedBlocks = account.WinCount
	nodes.RewardValue = account.BlockReward.RoundDown(4)
	nodes.SectorSize = account.SectorSizeStr

	sectorStatus := fmt.Sprintf("%d 全部, %d 有效, %d 错误, %d 恢复中", account.AllSectorCount, account.ActiveCount, account.FaultCount, account.RecoveryCount)
	nodes.SectorStatus = sectorStatus
	nodes.SectorTotal = account.AllSectorCount
	nodes.SectorEffective = account.ActiveCount
	nodes.SectorError = account.FaultCount
	nodes.SectorRecovering = account.RecoveryCount

	nodes.Address = account.RobustAddress
	nodes.MsgCount = account.MsgCount
	nodes.SectorType = account.ActorType
	nodes.CreateTime = utils.StringToTime(account.CreateTime)
	nodes.Height = account.Height
	nodes.MsigNode = account.Owner
	return nodes
}

func UpdateNodes(nodeParam string, timeTag int64) {
	nodes := services.FindAllNode(nodeParam)
	for _, oneNode := range nodes {
		if utils.TimeAddMinutes(oneNode.LastTime, 30).Compare(time.Now()) > 0 || oneNode.Balance.IsZero() {
			continue
		}
		var needMysql bool

		// 获取节点详细数据
		log.Printf("获取节点%s信息\n", oneNode.Node)
		n := GetNodeDetailByAddress(oneNode)

		n.LastTime = time.Now() //TimestampToTime(account.LastTime)
		time.Sleep(5 * time.Second)

		if !n.Balance.IsZero() || oneNode.Balance.IsZero() && oneNode.LuckyValue24h.GreaterThan(decimal.Zero) {
			needMysql = true

			log.Printf("获取节点%s的24hminer状态\n", oneNode.Node)
			md := MiningStats(oneNode.Node, "24h")
			n.MiningEfficiency = utils.DecimalValueFromFloat(md.Data.MiningEfficiencyFloat).Round(4)
			n.BlocksMined24h = md.Data.Blocks
			n.TotalRewards24h = md.Data.BlockReward.Round(4)
			n.LuckyValue24h = decimal.NewFromFloat(md.Data.LuckyValue)
			n.QualityAdjPowerDelta24h = utils.DecimalDiv1024x4Value(md.Data.QualityPowerGrowth.String())
			time.Sleep(5 * time.Second)

			log.Printf("获取节点%s的7dminer状态\n", oneNode.Node)
			md = MiningStats(oneNode.Node, "7d")
			n.MiningEfficiency7d = utils.DecimalValueFromFloat(md.Data.MiningEfficiencyFloat).Round(4)
			n.BlocksMined7d = md.Data.Blocks
			n.TotalRewards7d = md.Data.BlockReward.Round(4)
			n.LuckyValue7d = decimal.NewFromFloat(md.Data.LuckyValue)
			n.QualityAdjPowerDelta7d = utils.DecimalDiv1024x4Value(md.Data.QualityPowerGrowth.String())
			time.Sleep(5 * time.Second)

			log.Printf("获取节点%s的30dminer状态\n", oneNode.Node)
			md = MiningStats(oneNode.Node, "30d")
			n.MiningEfficiency30d = utils.DecimalValueFromFloat(md.Data.MiningEfficiencyFloat).Round(4)
			n.BlocksMined30d = md.Data.Blocks
			n.TotalRewards30d = md.Data.BlockReward.Round(4)
			n.LuckyValue30d = decimal.NewFromFloat(md.Data.LuckyValue)
			n.QualityAdjPowerDelta30d = utils.DecimalDiv1024x4Value(md.Data.QualityPowerGrowth.String())
			time.Sleep(5 * time.Second)
		}

		if n.LastDistributeTime.IsZero() {
			n.LastDistributeTime = time.Now()
		}
		if needMysql {
			transAmount := services.SumReward(n, timeTag)
			str := fmt.Sprintf("%f", transAmount)
			log.Printf("上次分币%s之后一共转出%s", n.LastDistributeTime, str)
			n.HasTransfer = n.HasTransfer.Add(utils.DecimalValue(str))

			// 获取节点的转入转出销毁数量
			mapA := services.SumValueByType(n.Node, timeTag)
			if value, ok := mapA["receive"]; ok {
				n.ReceiveAmount = n.ReceiveAmount.Add(value)
			}
			if value, ok := mapA["burn"]; ok {
				n.BurnAmount = n.BurnAmount.Add(value)
			}
			if value, ok := mapA["send"]; ok {
				n.SendAmount = n.SendAmount.Add(value)
			}

			//s := services.LuckyBlock{}
			//var count int64
			//s.CountByNodeTimeTag(oneNode.Node, timeTag, &count)
			n.TransferCount = oneNode.RealCount // 对齐交易的数量
		}
		n.TimeTag = timeTag
		services.UpdateNode(n)
		// 保存图表数据
		//services.SaveNodesChart(n)
		//blog抓取数据项目不再保存图表数据
		time.Sleep(10 * time.Second)
	}
}

/**
 * 保存矿池图表数据
 * 根据节点的所属部门分类保存各部门总算力
 */
func SavePoolChart() {
	now := time.Now()
	lastTime := utils.SetTime(now, now.Hour())
	poolChart := &models.PoolChart{
		LastTime:  lastTime,
		PowerUnit: "PiB",
		DeptId:    0,
	}

	nodes := services.FindAllNode("")

	deptPoolChart := make(map[int]*models.PoolChart)
	hasPowerCount := 0
	noPowerCount := 0

	for _, n := range nodes {
		if err := services.SaveNodesChart(n); err != nil {
			log.Printf("保存节点%s的图表数据失败：%s\n", n.Node, err)
		}

		updatePoolChart(poolChart, n) // 更新矿池图表数据

		if _, ok := deptPoolChart[n.DeptId]; !ok && n.DeptId > 0 {
			deptPoolChart[n.DeptId] = &models.PoolChart{
				LastTime:  lastTime,
				PowerUnit: "PiB",
				DeptId:    n.DeptId,
			}
		}
		if _, ok := deptPoolChart[n.DeptId]; ok {
			updatePoolChart(deptPoolChart[n.DeptId], n) // 更新部门矿池图表数据
		}

		if n.QualityAdjPower.IsZero() {
			noPowerCount++
		} else {
			hasPowerCount++
		}
	}

	services.SavePoolChart(poolChart) // 保存矿池图表数据

	log.Printf("一共更新的 %d 个节点，其中有算力的节点 %d 个, 算力为0的节点 %d 个。\n", len(nodes), hasPowerCount, noPowerCount)

	for k, v := range deptPoolChart {
		log.Printf("保存部门%d的矿池数据\n", k)
		services.SavePoolChart(v)
	}
}

func updatePoolChart(poolChart *models.PoolChart, node models.Nodes) {
	// 累加节点数据到矿池图表
	poolChart.Balance = poolChart.Balance.Add(node.Balance)
	poolChart.AvailableBalance = poolChart.AvailableBalance.Add(node.AvailableBalance)
	poolChart.SectorPledgeBalance = poolChart.SectorPledgeBalance.Add(node.SectorPledgeBalance)
	poolChart.VestingFunds = poolChart.VestingFunds.Add(node.VestingFunds)
	poolChart.QualityAdjPower = poolChart.QualityAdjPower.Add(node.QualityAdjPower)
	poolChart.PowerPoint = poolChart.PowerPoint.Add(node.PowerPoint)
	poolChart.ControlBalance = poolChart.ControlBalance.Add(node.ControlBalance)
	poolChart.RewardValue = poolChart.RewardValue.Add(node.RewardValue)
}

func HandUpdate(nodeParam string) {
	nodes := services.FindAllNode(nodeParam)

	var contents string

	for _, oneNode := range nodes {
		//if TimeAddMinutes(oneNode.LastHandTime, 30)
		if oneNode.LastHandTime != nil && oneNode.LastHandTime.Add(time.Minute*30).Compare(time.Now()) > 0 {
			continue
		}
		// 获取节点详细数据
		log.Printf("获取节点%s信息\n", oneNode.Node)
		n := GetNodeDetailByAddress(oneNode)
		//n.Node = oneNode.Node
		time.Sleep(5 * time.Second)

		// 获取节点可用余额数据
		log.Printf("获取节点%s可用余额\n", oneNode.Node)
		now := time.Now()
		n.LastHandTime = &now //TimestampToTime(account.LastTime)
		//fmt.Printf("nodes=%+v\n", n)
		time.Sleep(5 * time.Second)

		services.UpdateNode(n)

		if n.SectorError > oneNode.SectorError && (n.SectorError-oneNode.SectorError > 100) {
			contents += fmt.Sprintf("节点：%s 当前算力：%s %s，扇区状态：%s，错误扇区数量增加 %d。<br/>", n.Node, n.QualityAdjPower.String(), n.PowerUnit, n.SectorStatus, n.SectorError-oneNode.SectorError)
		}
	}
	if len(contents) > 0 {
		var msg models.Msg
		msg.Title = "扇区错误"
		msg.Content = contents
		msg.CreateTime = time.Now()
		msg.SendStatus = 0
		services.InsertMsg(msg)
	}

	log.Printf("手动更新节点余额成功。")
}

type BlockSta struct {
	Data struct {
		StatsType    string `json:"statsType"`
		DurationHour int    `json:"durationHour"`
		Blocks       []struct {
			HeightTimeStr         string          `json:"heightTimeStr"`
			BlocksGrowth          int             `json:"blocksGrowth"`
			BlocksRewardGrowthFil decimal.Decimal `json:"blocksRewardGrowthFil"`
		} `json:"blocks"`
	} `json:"data"`
}

// BlockStats 获取账户的小时报块数据
func BlockStats(node string, tp string) BlockSta {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.filutils.com/api/v2/miner/%s/blockstats", node)
	jsonStr := []byte(`{ "statsType": "` + tp + `"}`)
	reqSpider, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	reqSpider.Header.Set("Accept", "*/*")
	reqSpider.Header.Set("Locale", "zh")
	reqSpider.Header.Set("Content-Type", "application/json;charset=UTF-8")
	reqSpider.Header.Set("sec-ch-ua", "\"Chromium\";v=\"126\", \"Google Chrome\";v=\"126\", \"Not:A-Brand\";v=\"99\"")
	reqSpider.Header.Set("sec-ch-ua-mobile", "?0")
	reqSpider.Header.Set("sec-ch-ua-platform", "macOS")
	reqSpider.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	respSpider, err := client.Do(reqSpider)
	if err != nil {
		log.Fatal(err)
	}
	defer respSpider.Body.Close()
	bodyText, err := ioutil.ReadAll(respSpider.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result BlockSta
	_ = json.Unmarshal(bodyText, &result) //byte to json

	return result
}

func UpdateBlockStats(nodeParam string) {
	nodes := services.FindAllNode(nodeParam)

	for _, oneNode := range nodes {
		if oneNode.Balance.IsZero() {
			continue
		}
		// 获取节点报块数据
		log.Printf("获取节点%s小时报块信息\n", oneNode.Node)
		b := BlockStats(oneNode.Node, "24h")
		time.Sleep(5 * time.Second)

		blocks := SortBlockStats(b, oneNode.Node)
		lastBs := services.GetLastOne(oneNode.Node)
		if !reflect.DeepEqual(lastBs, models.BlockStats{}) {
			blocks = FilterHasSave(blocks, lastBs)
		}
		if len(blocks) > 0 {
			services.InsertStats(blocks)
		}
	}

	log.Printf("节点报块数据更新成功。")
}

func SortBlockStats(bs BlockSta, node string) []models.BlockStats {
	var mbs []models.BlockStats
	now := time.Now()
	hour := now.Hour() - 1 //过了当前小时才是前一小时
	var heightTime time.Time
	for _, b := range bs.Data.Blocks {
		heightHour, _ := strconv.Atoi(strings.Split(b.HeightTimeStr, ":")[0])
		if heightHour > hour { //时间超过了当前的时间，时间就是前一天
			lastDay := now.AddDate(0, 0, -1)
			heightTime = utils.SetTime(lastDay, heightHour)
		} else {
			heightTime = utils.SetTime(now, heightHour)
		}

		mbs = append(mbs, models.BlockStats{
			Node:                  node,
			StatsType:             bs.Data.StatsType,
			DurationHour:          bs.Data.DurationHour,
			BlocksGrowth:          b.BlocksGrowth,
			BlocksRewardGrowthFil: b.BlocksRewardGrowthFil,
			HeightTimeStr:         b.HeightTimeStr,
			HeightTime:            heightTime,
			CreatedAt:             now,
		})
	}
	return mbs
}

func FilterHasSave(blocks []models.BlockStats, last models.BlockStats) []models.BlockStats {
	var mbs []models.BlockStats
	for _, b := range blocks {
		if b.HeightTime.After(last.HeightTime) {
			mbs = append(mbs, b)
		}
	}
	return mbs
}
