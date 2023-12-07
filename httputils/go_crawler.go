package httputils

import (
	"blog/models"
	"blog/services"
	"fmt"
	"github.com/anaskhan96/soup"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetHttpContent(requestUrl string) string {
	//requestUrl := "https://www.liaoxuefeng.com/"
	// 发送Get请求
	rsp, err := http.Get(requestUrl)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	content := string(body)
	defer rsp.Body.Close()

	return content
}

// GetAllNodeByAddress 获取当前node下所有的节点名称
func GetAllNodeByAddress(node string) []models.Nodes {
	content := GetHttpContent("https://filfox.info/zh/address/" + node)
	// 下面主要是解析标签
	doc := soup.HTMLParse(content)
	subDocs := doc.FindAll("dl", "class", "flex")
	var nodes []models.Nodes
	for _, subDoc := range subDocs {
		link := subDoc.Find("dt")
		if link.Error != nil {
			//log.Println(link.Error)
			continue
		}
		if strings.TrimSpace(link.Text()) == "名下存储提供者" {
			fmt.Println(strings.TrimSpace(link.Text()))
			providers := subDoc.FindAll("p")
			for _, provide := range providers {
				nodeStr := provide.Find("a")
				if nodeStr.Error != nil {
					continue
				}
				oneNode := models.Nodes{
					Node:       strings.TrimSpace(nodeStr.Text()),
					MsigNode:   node,
					CreateTime: time.Now(),
					LastTime:   time.Now(),
				}
				nodes = append(nodes, oneNode)
			}
		}
	}
	return nodes
}

// GetNodeDetailByAddress 根据节点名称查询节点详细信息
func GetNodeDetailByAddress(nodes models.Nodes) models.Nodes {
	node := nodes.Node
	content := GetHttpContent("https://filfox.info/zh/address/" + node)
	// 下面主要是解析标签
	doc := soup.HTMLParse(content)
	subDocs := doc.FindAll("div", "class", "rounded-md")
	//var m = make(map[string]string)
	//var nodes models.Nodes
	var _thisLast int
	for _, subDoc := range subDocs {
		links := subDoc.FindAll("p")
		for _, link := range links {
			if link.Error != nil {
				//log.Println(link.Error)
				continue
			}
			var text = strings.TrimSpace(link.Text())
			if text == "账户余额" {
				balance := strings.TrimSpace(link.FindNextElementSibling().Text())
				nodes.Balance = DecimalValue(FormaterString(strings.Split(balance, " ")[0]))
				continue
			}
			// 可用余额页面抓取有问题，采用请求的方式获得
			/*if strings.Contains(text, "可用余额") {
				s := link.Children()
				for _, value := range s {
					if value.NodeValue == "span" {
						v := value.FindNextSibling().Text()
						split := strings.Split(strings.TrimSpace(strings.ReplaceAll(v, ":", "")), " ")
						nodes.AvailableBalance = DecimalValue(split[0])
					}
				}
				continue
			}*/
			if strings.Contains(text, "扇区抵押") {
				//m["扇区抵押"] = strings.SplitAfter(text, ":")[1]
				pledge := strings.TrimSpace(strings.SplitAfter(text, ":")[1])
				nodes.SectorPledgeBalance = DecimalValue(FormaterString(strings.Split(pledge, " ")[0]))
				continue
			}
			if strings.Contains(text, "提供存储服务锁仓") {
				//m["提供存储服务锁仓"] = strings.SplitAfter(text, ":")[1]
				funds := strings.TrimSpace(strings.SplitAfter(text, ":")[1])
				nodes.VestingFunds = DecimalValue(FormaterString(strings.Split(funds, " ")[0]))
				continue
			}
			if text == "有效算力" {
				_thisLast = 2
				continue
			}
			if _thisLast == 2 {
				//m["有效算力"] = text
				_thisLast = 0
				nodes.QualityAdjPower = DecimalValue(FormaterString(strings.Split(text, " ")[0]))
				nodes.PowerUnit = strings.TrimSpace(strings.Split(text, " ")[1])
				if nodes.PowerUnit == "TiB" {
					nodes.QualityAdjPower = nodes.QualityAdjPower.Div(DecimalValue("1000")).Round(2)
					nodes.PowerUnit = "PiB"
				} else if nodes.PowerUnit == "GiB" {
					nodes.QualityAdjPower = nodes.QualityAdjPower.Div(DecimalValue("1000000")).Round(2)
					nodes.PowerUnit = "PiB"
				}

				continue
			}
			if strings.Contains(text, "占比") {
				//m["算力占比"] = strings.SplitAfter(text, ":")[1]
				point := strings.Split(text, ":")[1]
				nodes.PowerPoint = DecimalValue(FormaterString(strings.Split(point, "%")[0]))
				continue
			}
			if strings.Contains(text, "排名") {
				//m["排名"] = strings.SplitAfter(text, ":")[1]
				nodes.PowerGrade = strings.TrimSpace(strings.SplitAfter(text, ":")[1])
				continue
			}
			if strings.Contains(text, "原值算力") {
				s := link.FindNextElementSibling().Children()
				for i, value := range s {
					if i == 2 {
						v := strings.TrimSpace(strings.ReplaceAll(value.NodeValue, ":", ""))
						num, _ := strconv.Atoi(v)
						nodes.WeightedBlocks = num
					}
				}

				continue
			}

			if strings.Contains(text, "累计出块奖励") {
				//m["累计出块奖励"] = strings.SplitAfter(text, ":")[1]
				reward := strings.TrimSpace(strings.SplitAfter(text, ":")[1])
				nodes.RewardValue = DecimalValue(FormaterString(strings.SplitAfter(reward, " ")[0]))
				continue
			}
			if strings.Contains(text, "扇区大小") {
				//m["扇区大小"] = strings.SplitAfter(text, ":")[1]
				nodes.SectorSize = strings.TrimSpace(strings.SplitAfter(text, ":")[1])
				continue
			}
			if strings.Contains(text, "扇区状态") {
				sectorStatus := link.FindNextElementSibling().FindAll("span")
				var stas string
				for _, sta := range sectorStatus {
					stas = stas + sta.Text()
				}
				for _, sta := range sectorStatus {
					split := strings.Split(strings.TrimSpace(strings.ReplaceAll(sta.Text(), ",", "")), " ")
					num, _ := strconv.Atoi(split[0])
					if split[1] == "全部" {
						nodes.SectorTotal = num
					} else if split[1] == "有效" {
						nodes.SectorEffective = num
					} else if split[1] == "错误" {
						nodes.SectorError = num
					} else {
						nodes.SectorRecovering = num
					}

				}
				//m["扇区状态"] = stas
				nodes.SectorStatus = strings.TrimSpace(stas)
				continue
			}
			if strings.Contains(text, "地址:") {
				//m["地址"] = link.FindNextElementSibling().Text()
				nodes.Address = strings.TrimSpace(link.FindNextElementSibling().Text())
				continue
			}
			if strings.Contains(text, "消息数") {
				m := link.FindNextElementSibling().Text()
				count, _ := strconv.Atoi(FormaterString(m))
				nodes.MsgCount = count
				continue
			}
			if strings.Contains(text, "类型:") {
				//m["类型"] = link.FindNextElementSibling().Text()
				nodes.SectorType = strings.TrimSpace(link.FindNextElementSibling().Text())
				continue
			}
			if strings.Contains(text, "创建时间") {
				//m["创建时间"] = link.FindNextElementSibling().Text()
				nodes.CreateTime = StringToTime(strings.TrimSpace(link.FindNextElementSibling().Text()))
				continue
			}
		}
	}
	return nodes
}

// GetControlDetailByAddress 根据地址查询详细信息
func GetControlDetailByAddress(addr string) Account {
	content := GetHttpContent("https://filfox.info/zh/address/" + addr)
	// 下面主要是解析标签
	doc := soup.HTMLParse(content)
	subDocs := doc.FindAll("div", "class", "rounded-md")
	var account Account
	for _, subDoc := range subDocs {
		links := subDoc.FindAll("dl")
		for _, link := range links {
			if link.Error != nil {
				//log.Println(link.Error)
				continue
			}
			text := strings.TrimSpace(link.Find("dt").Text())
			if text == "余额" {
				balance := strings.TrimSpace(link.FindNextElementSibling().Text())
				account.ControlBalance = DecimalValue(FormaterString(strings.Split(balance, " ")[0]))
				break
			}
		}
		break
	}
	return account
}

func SaveNodes(node string) {
	nodes := GetAllNodeByAddress(node)
	if len(nodes) > 0 {
		services.SaveNode(nodes)
	}
}

func UpdateNodes(nodeParam string, timeTag int64) {
	nodes := services.FindAllNode(nodeParam)
	var hasPowerCount int
	var noPowerCount int

	var savePool bool

	for _, oneNode := range nodes {
		if TimeAddMinutes(oneNode.LastTime, 30).Compare(time.Now()) > 0 {
			continue
		}
		savePool = true

		// 获取节点详细数据
		log.Printf("获取节点%s信息\n", oneNode.Node)
		n := GetNodeDetailByAddress(oneNode)
		//n.Node = oneNode.Node
		time.Sleep(5 * time.Second)

		// 获取节点可用余额数据
		log.Printf("获取节点%s可用余额\n", oneNode.Node)
		account := BalanceStats(oneNode.Node)
		n.AvailableBalance = account.AvailableBalance
		n.Height = account.Height
		n.LastTime = time.Now() //TimestampToTime(account.LastTime)
		//fmt.Printf("nodes=%+v\n", n)
		time.Sleep(5 * time.Second)

		//if !oneNode.Balance.IsZero() {
		log.Printf("获取节点%s的24hminer状态\n", oneNode.Node)
		miningDetail := MiningStats(oneNode.Node)
		n.BlocksMined24h = miningDetail.BlocksMined
		n.WeightedBlocksMined24h = miningDetail.WeightedBlocksMined
		n.TotalRewards24h = DecimalDiv18Value(miningDetail.TotalRewards)
		n.LuckyValue24h = decimal.NewFromFloat(miningDetail.LuckyValue)
		n.QualityAdjPowerDelta24h = DecimalDiv1024x4Value(miningDetail.QualityAdjPowerDelta)
		time.Sleep(5 * time.Second)

		/*log.Printf("获取节点%s的7dminer状态\n", oneNode.Node)
		miningDetail = MiningStatsCycle(oneNode.Node, "7d")
		n.BlocksMined7d = miningDetail.BlocksMined
		n.WeightedBlocksMined7d = miningDetail.WeightedBlocksMined
		n.TotalRewards7d = DecimalDiv18Value(miningDetail.TotalRewards)
		n.LuckyValue7d = decimal.NewFromFloat(miningDetail.LuckyValue)
		n.QualityAdjPowerDelta7d = DecimalDiv1024x4Value(miningDetail.QualityAdjPowerDelta)
		time.Sleep(5 * time.Second)

		log.Printf("获取节点%s的30dminer状态\n", oneNode.Node)
		miningDetail = MiningStatsCycle(oneNode.Node, "30d")
		n.BlocksMined30d = miningDetail.BlocksMined
		n.WeightedBlocksMined30d = miningDetail.WeightedBlocksMined
		n.TotalRewards30d = DecimalDiv18Value(miningDetail.TotalRewards)
		n.LuckyValue30d = decimal.NewFromFloat(miningDetail.LuckyValue)
		n.QualityAdjPowerDelta30d = DecimalDiv1024x4Value(miningDetail.QualityAdjPowerDelta)
		time.Sleep(5 * time.Second)*/
		//}

		// 获取节点的控制地址 独立方法获取 独立表保存，一个节点存在多个控制地址
		/*if len(n.ControlAddress) == 0 {
			log.Printf("获取节点%s控制地址\n", oneNode.Node)
			controlAddress := SpiderMessage(oneNode.Node)
			n.ControlAddress = controlAddress
			time.Sleep(5 * time.Second)
		}*/

		// 获取控制地址的余额
		//log.Printf("获取节点%s控制地址余额\n", oneNode.Node)
		//if len(n.ControlAddress) > 0 {
		//	account = BalanceControl(n.ControlAddress)
		//	n.ControlBalance = account.ControlBalance
		//}

		if n.Height > 0 {
			if n.LastDistributeTime.IsZero() {
				n.LastDistributeTime = time.Now()
			}
			transAmount := services.SumReward(n, timeTag)
			str := fmt.Sprintf("%f", transAmount)
			log.Printf("上次分币%s之后一共转出%s", n.LastDistributeTime, str)
			n.HasTransfer = n.HasTransfer.Add(DecimalValue(str))

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

			s := services.LuckyBlock{}
			var count int64
			s.CountByNodeTimeTag(oneNode.Node, timeTag, &count)
			n.TransferCount = n.TransferCount + count

			n.TimeTag = timeTag
			services.UpdateNode(n)
			// 保存图表数据
			services.SaveNodesChart(n)
			time.Sleep(15 * time.Second)
		}
		if n.QualityAdjPower.IsZero() {
			noPowerCount++
		} else {
			hasPowerCount++
		}

	}

	// 保存矿池图表数据
	if savePool {
		log.Printf("一共更新的 %d 个节点，其中有算力的节点 %d 个, 算力为0的节点 %d 个。\n", len(nodes), hasPowerCount, noPowerCount)
		savePoolChart()
	} else {
		log.Printf("没有需要更新的节点。")
	}
}

func savePoolChart() {
	nodes := services.FindAllNode("")
	poolChart := new(models.PoolChart)

	for _, n := range nodes {
		poolChart.Balance = poolChart.Balance.Add(n.Balance)
		poolChart.AvailableBalance = poolChart.AvailableBalance.Add(n.AvailableBalance)
		poolChart.SectorPledgeBalance = poolChart.SectorPledgeBalance.Add(n.SectorPledgeBalance)
		poolChart.VestingFunds = poolChart.VestingFunds.Add(n.VestingFunds)
		poolChart.QualityAdjPower = poolChart.QualityAdjPower.Add(n.QualityAdjPower)
		poolChart.PowerPoint = poolChart.PowerPoint.Add(n.PowerPoint)
		poolChart.ControlBalance = poolChart.ControlBalance.Add(n.ControlBalance)
		poolChart.RewardValue = poolChart.RewardValue.Add(n.RewardValue)
	}
	poolChart.LastTime = time.Now()
	poolChart.PowerUnit = "PiB"
	services.SavePoolChart(poolChart)
}
