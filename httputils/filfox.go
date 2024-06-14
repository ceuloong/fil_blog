package httputils

import (
	"blog/models"
	"blog/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Result 返回的结构体
type Result struct {
	TotalCount int64 `json:"totalCount"`
	Transfers  []struct {
		Height    int64  `json:"height"`
		Timestamp int64  `json:"timestamp"`
		From      string `json:"from"`
		To        string `json:"to"`
		Message   string `json:"message"`
		Value     string `json:"value"`
		Type      string `json:"type"`
	} `json:"transfers"`
	Types []string `json:"types"`
}

// UpdateDetail 结构体
type UpdateDetail struct { //
	services.LuckyBlock
	page int
}

type SpiderNode struct {
	Node     string
	PidNode  string
	TimeTag  int64
	Category string
	Page     int
}

var wg sync.WaitGroup

const pageSize = 100

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix)
	// 通过SetPrefix设置Logger结构体里的prefix属性
	log.SetPrefix("INFO:")
}

func Start(timeTag int64, nodeParam string) {
	// f17225euatkghukeyu6exm6j72fw5aneiiyapotbq  f1emjr5wnycomu7dgbt6aiaiir6rn53ggioiou56a
	//f01807413 f01822659 f01845913  f01874748  f01834260  f01899511
	//nodes := []string{"f01807413", "f01822659", "f01845913", "f01874748", "f01834260", "f01899511"}
	nodes := services.FindAllNode(nodeParam)

	pg := 5
	page := pageCount(len(nodes), pg)
	for p := 1; p <= page; p++ {
		pageNode(nodes, p, pg, timeTag)
		time.Sleep(5 * time.Second)
	}
}

func pageNode(nodes []models.Nodes, p int, pg int, timeTag int64) {
	s := services.LuckyBlock{}
	var newNodes []SpiderNode
	var updateNodes []SpiderNode
	//var m = make(map[string]int)

	index := pg * p
	if index > len(nodes) {
		index = len(nodes)
	}

	for i := (p - 1) * pg; i < index; i++ {

		//for i := 0; i < len(nodes); i++ {

		node := SpiderNode{
			Node:     nodes[i].Node,
			Category: "",
			TimeTag:  timeTag,
		}

		var count int64
		count = nodes[i].TransferCount
		if count == 0 {
			s.CountByNode(node.Node, &count)
		}

		//if count == 0 {
		//	newNodes = append(newNodes, node)
		//} else {
		//	if len(newNodes) > 0 {
		//		continue
		//	}
		total := getTotalNum(node.Node) // 获取全部记录数量
		services.UpdateRealCount(total, node.Node)
		page := pageCount(total-int(count), pageSize)
		if page <= 0 {
			continue
		}
		node.Page = page
		updateNodes = append(updateNodes, node)
		//}
	}

	// 抓取控制地址记录
	/*for i := 0; i < len(nodes); i++ {
		if nodes[i].Status > 1 && nodes[i].Status < 5 {
			node := nodes[i].ControlAddress
			var count int64
			s.CountByNode(node, &count)

			if count == 0 {
				newNodes = append(newNodes, node)
			} else {
				total := getTotalNum(node) // 获取全部记录数量
				page := pageCount(total-int(count), pageSize)
				if page == 0 {
					continue
				}
				m[node] = page
				updateNodes = append(updateNodes, node)
			}
		}
	}*/

	if len(newNodes) > 0 { // 新节点首次抓数据
		log.Printf("新节点首次抓数据: %d\n", len(newNodes))
		newNodeSave(newNodes)
	} else if len(updateNodes) > 0 { // 已经完整抓取过，实时更新就行
		log.Printf("更新节点: len(m)= %d\n", len(updateNodes))
		updateNodeNew(updateNodes)
	} else {
		log.Println("no data need to update. next page: ", p+1)
	}
}

// 已存在的节点更新数据
func updateNodeNew(nodes []SpiderNode) {
	start := time.Now()

	wg.Add(len(nodes))
	for i := 0; i < len(nodes); i++ {
		go func(i int) {
			defer wg.Done()
			GetHttpHtmlNew(nodes[i])
		}(i)
	}
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("updateNodeNew WaitGroupStart Time %s\n ", elapsed)
}

// 已存在的节点更新数据
/*func updateNode(m map[string]services.LuckyBlock) {
	for node, block := range m {
		var needToSave []models.LuckyBlock
		needToSave = getSpiders(node, block, 0, needToSave,)
		if len(needToSave) > 0 {
			log.Println("保存新数据：", len(needToSave))
			//services.Insert(needToSave)
			total := len(needToSave)
			pageCount := int(math.Ceil(float64(total) / float64(pageSize)))
			for i := 1; i <= pageCount; i++ {
				start, end := SlicePage(i, pageSize, total)
				blocks := needToSave[start:end]
				services.Insert(blocks)
			}
		}
		time.Sleep(1 * time.Second)
	}
}*/

func getSpiders(node SpiderNode, block services.LuckyBlock, page int, needToSave []models.LuckyBlock) []models.LuckyBlock {
	time.Sleep(1 * time.Second)
	total, spiders := Spider(node, page, pageSize, 0)
	log.Printf("spiders.len: %d,total: %d", len(spiders), total)
	need := services.NeedToSave(block, spiders)
	if need == nil {
		return needToSave
	}

	needToSave = append(need, needToSave...)

	//if len(needToSave) == 20 {
	//	getSpiders(node, block, page+1)
	//} else {
	//	return needToSave
	//}
	return getSpiders(node, block, page+1, needToSave)
}

// 节点首次抓取数据
func newNodeSave(nodes []SpiderNode) {
	log.Println("开始抓取数据")
	start := time.Now()

	wg.Add(len(nodes))
	for i := 0; i < len(nodes); i++ {
		go func(i int) {
			defer wg.Done()
			GetHttpHtmlNew(nodes[i])
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("WaitGroupStart Time %s\n ", elapsed)
}

// GetHttpHtmlNew 首次抓取数据时执行
func GetHttpHtmlNew(node SpiderNode) {
	p := node.Page
	total := getTotalNum(node.Node) // 获取全部记录数量  8078
	// 首次抓取数据时100条每页
	var addPage int
	var errCount int
	var lastBlock services.LuckyBlock
	if p == 0 {
		p = pageCount(total, pageSize)
	} else {
		lastBlock = services.FindLastByNode(node.Node)
		log.Printf("获取节点%s保存的最后一条数据%T\n", node.Node, lastBlock)
	}
	needRemoveRep := true

	for page := p - 1; page >= 0; page-- {
		var totalBlock []models.LuckyBlock
		var totalCount int
		totalCount, totalBlock = Spider(node, page, pageSize, errCount) // 保存数据库
		if len(totalBlock) == 0 {
			for i := 0; i < 5; i++ {
				time.Sleep(5 * time.Second)
				totalCount, totalBlock = Spider(node, page, pageSize, errCount) // 保存数据库
				if len(totalBlock) > 0 {
					break
				}
				log.Printf("节点%s查询%d页时，数量为0，for i=%d\n", node, page, i)
			}
		}
		if len(totalBlock) > 0 {
			if needRemoveRep && lastBlock != (services.LuckyBlock{}) { //保存最后一页时去重
				totalBlock = services.NeedToSave(lastBlock, totalBlock)
				log.Printf("更新节点%s时去重，page:%d", node.Node, p)
			}

			if len(totalBlock) > 0 {
				needRemoveRep = false
				services.Insert(totalBlock)
			}
		}

		if totalCount > total {
			if page < totalCount/pageSize {
				addPage = page + 1 // 加1是再次查询当前页数据，总数量增加了一，页数的位置有变化
				_, newPage := Spider(node, addPage, pageSize, 0)
				log.Printf("totalCount:%d, total:%d,数量有变化，需要向前查一页addPage:%d, newPage长度:%d", totalCount, total, addPage, len(newPage))
				if len(newPage) > 0 {
					needSave := saveNewPage(newPage, totalCount-total)
					total = totalCount
					if len(needSave) > 0 {
						services.Insert(needSave)
					}
				}
			}
		}
		time.Sleep(2 * time.Second)
		log.Printf("保存%s第%d页数据成功\n", node.Node, page)
	}
}

// GetHttpHtml 首次抓取数据时执行
func GetHttpHtml(node SpiderNode) {
	GetHttpHtmlNew(node)
	/*total := getTotalNum(node) // 获取全部记录数量  8078
	// 首次抓取数据时100条每页
	pageSize := 100
	var addPage int
	var errCount int
	for page := total / pageSize; page >= 0; page-- {
		var totalBlock []models.LuckyBlock
		var totalCount int
		totalCount, totalBlock = Spider(node, page, pageSize, errCount) // 保存数据库
		if len(totalBlock) == 0 {
			for i := 0; i < 5; i++ {
				time.Sleep(5 * time.Second)
				totalCount, totalBlock = Spider(node, page, pageSize, errCount) // 保存数据库
				if len(totalBlock) > 0 {
					break
				}
				log.Printf("节点%s查询%d页时，数量为0，for i=%d\n", node, page, i)
			}
		}
		if len(totalBlock) > 0 {
			services.Insert(totalBlock)
		}

		if totalCount > total {
			if page < totalCount/pageSize {
				addPage = page + 2 // 加2是需要查当前页的上一页数据
				_, newPage := Spider(node, addPage, pageSize, 0)
				log.Printf("totalCount:%d, total:%d,数量有变化，需要向前查一页addPage:%d, newPage长度:%d", totalCount, total, addPage, len(newPage))
				if len(newPage) > 0 {
					needSave := saveNewPage(newPage, totalCount-total)
					total = totalCount
					if len(needSave) > 0 {
						services.Insert(needSave)
					}
				}
			}
		}

		time.Sleep(2 * time.Second)
		log.Printf("保存%s第%d页数据成功\n", node, page)
	}*/
}

// 获取全部的区块数量
func getTotalNum(node string) int {
	url := `https://filfox.info/api/v1/address/` + node + `/transfers?pageSize=20&page=0`
	//url := `https://blog.csdn.net/phoenix/web/v1/comment/list/121340420?page=1&size=100`
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Locale", "zh")
	req.Header.Set("sec-ch-ua", "\"Chromium\";v=\"115\", \"Google Chrome\";v=\"115\", \"Not:A-Brand\";v=\"99\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "macOS")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result Result
	_ = json.Unmarshal(bodyText, &result) //byte to json
	total := result.TotalCount
	fmt.Printf("node %s total: %d", node, total)
	return int(total)
}

// Spider 传入页数，一页一页爬取
func Spider(node SpiderNode, page int, pageSize int, errCount int) (int, []models.LuckyBlock) {
	var tmp []models.LuckyBlock
	p := strconv.Itoa(page)
	log.Printf("当前页page:%d, p:%s\n", page, p)
	client := &http.Client{}
	url := fmt.Sprintf("https://filfox.info/api/v1/address/%s/transfers?pageSize=%d&page=%d", node.Node, pageSize, page)
	reqSpider, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	reqSpider.Header.Set("Accept", "application/json, text/plain, */*")
	reqSpider.Header.Set("Locale", "zh")
	reqSpider.Header.Set("sec-ch-ua", "\"Chromium\";v=\"112\", \"Google Chrome\";v=\"112\", \"Not:A-Brand\";v=\"99\"")
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
	var result Result
	_ = json.Unmarshal(bodyText, &result) //byte to json
	num := len(result.Transfers)

	/*	if num == 0 && errCount < 10 {
		log.Printf("查询%s第%d页时为空，重新查询当前页，errCount:%d\n", node, page, errCount)
		time.Sleep(5 * time.Second)
		Spider(node, page, pageSize, errCount+1) // 可能超时没有返回数据，当前页数据再查询
	}*/

	transfers := result.Transfers
	total := result.TotalCount
	for i := num - 1; i >= 0; i-- {
		var luckBlock models.LuckyBlock
		luckBlock.Height = transfers[i].Height
		luckBlock.Node = node.Node
		luckBlock.Date = TimestampToTime(transfers[i].Timestamp)
		luckBlock.NodeFrom = transfers[i].From
		luckBlock.NodeTo = transfers[i].To
		luckBlock.RewardValue = DecimalDiv18Value(transfers[i].Value)
		luckBlock.Message = transfers[i].Message
		luckBlock.Type = transfers[i].Type
		luckBlock.CreateTime = time.Now()
		luckBlock.TimeTag = node.TimeTag
		luckBlock.Category = node.Category
		luckBlock.PidNode = node.PidNode
		tmp = append(tmp, luckBlock) //
	}

	return int(total), tmp
}

//func removeRepByMap(slc []LuckyBlock) []LuckyBlock { //去除重复的元素
//	var result []LuckyBlock          //存放返回的不重复切片
//	tempMap := map[LuckyBlock]byte{} // 存放不重复主键
//	for _, e := range slc {
//		l := len(tempMap)
//		tempMap[e] = 0         //当e存在于tempMap中时，再次添加是添加不进去的，因为key不允许重复
//		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
//			result = append(result, e) //当元素不重复时，将元素添加到切片result中
//		}
//	}
//	return result
//}

// new  old
/*func saveNewPage(tmp []models.LuckyBlock, totalBlock []models.LuckyBlock) []models.LuckyBlock {
	var needSave []models.LuckyBlock
	for i := 0; i < len(tmp); i++ {
		has := false
		for j := 0; j < len(totalBlock); j++ {
			if reflect.DeepEqual(tmp[i], totalBlock[j]) {
				has = true
				break
			}
		}
		if !has {
			needSave = append(needSave, tmp[i])
		}
	}
	return needSave
}*/

func saveNewPage(tmp []models.LuckyBlock, count int) []models.LuckyBlock {
	var needSave []models.LuckyBlock
	for i := len(tmp) - 1; i > len(tmp)-1-count; i-- {
		needSave = append(needSave, tmp[i])
	}
	return needSave
}

func SlicePage(page, pageSize, nums int) (sliceStart, sliceEnd int) {
	// 定义page和size的默认值
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 100
	}
	// 如果pageSize大于num（切片长度）, 那么sliceEnd直接返回num的值
	if pageSize > nums {
		return 0, nums
	}
	// 总页数计算，math.Ceil 返回不小于计算值的最小整数（的浮点值）
	pageCount := int(math.Ceil(float64(nums) / float64(pageSize)))
	if page > pageCount {
		return 0, 0
	}
	sliceStart = (page - 1) * pageSize
	sliceEnd = sliceStart + pageSize
	// 如果页总数比sliceEnd小，那么就把总数赋值给sliceEnd
	if sliceEnd > nums {
		sliceEnd = nums
	}
	return sliceStart, sliceEnd
}

func pageCount(total int, pageSize int) int {
	page := int(math.Ceil(float64(total) / float64(pageSize)))
	return page
}
