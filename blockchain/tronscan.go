package blockchain

import (
	"blog/httputils"
	"blog/models"
	"blog/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

type TronScanApi struct {
	Total     int64 `json:"total"`
	Transfers []struct {
		TxId      string `json:"transaction_id"`
		Timestamp int64  `json:"block_ts"`
		Height    int64  `json:"block"`
		From      string `json:"from_address"`
		FromTag   struct {
			Tag string `json:"from_address_tag"`
		} `json:"from_address_tag"`
		To    string `json:"to_address"`
		ToTag struct {
			Tag string `json:"to_address_tag"`
		} `json:"to_address_tag"`
		Contract     string `json:"contract_address"`
		Value        string `json:"quant"`
		EventType    string `json:"event_type"`
		ContractType string `json:"contract_type"`
		Confirmed    bool   `json:"confirmed"`
	} `json:"token_transfers"`
	AddressInfo []string `json:"normalAddressInfo"`
}

const owner string = "TZCTLnJ4xE1EbqoZmXmtkQZ3UbavV9cVTd"
const afterTime string = "2022-01-19 06:00:00"
const beforeTime string = "2022-08-19 06:00:00"

var m = make(map[string]string)

func StartTron(level int) {
	m = services.FindAllTronByTxId()
	addrs := services.FindAllTronByLevel(level)
	for _, addr := range addrs {
		//if i < 34 {
		//	continue
		//}
		GetHttpHtmlNew(addr.ToAddr)
		time.Sleep(5 * time.Second)
	}
}

// GetHttpHtmlNew 首次抓取数据时执行
func GetHttpHtmlNew(addr string) {
	total := getTotalNum(addr) // 获取全部记录数量  8078
	// 首次抓取数据时100条每页
	pageSize := 50
	p := pageCount(total, pageSize)

	if p > 100 {
		return
	}

	for page := p - 1; page >= 0; page-- {
		var totalBlock []models.TronScan
		total, totalBlock = Spider(addr, page, pageSize) // 保存数据库
		/*if len(totalBlock) == 0 {
			for i := 0; i < 5; i++ {
				time.Sleep(5 * time.Second)
				_, totalBlock = Spider(addr, page, pageSize) // 保存数据库
				if len(totalBlock) > 0 {
					break
				}
				log.Printf("节点%s查询%d页时，数量为0，for i=%d\n", addr, page, i)
			}
		}*/
		if len(totalBlock) > 0 {
			services.InsertTron(totalBlock)
		}

		time.Sleep(2 * time.Second)
		log.Printf("保存%s第%d页数据成功\n", addr, page)
		if total < 0 {
			break
		}
	}
}

// 获取全部的区块数量
func getTotalNum(addr string) int {
	url := fmt.Sprintf("https://apilist.tronscanapi.com/api/filter/trc20/transfers?limit=20&start=0&sort=-timestamp&count=true&filterTokenValue=0&relatedAddress=%s", addr)
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
	var result TronScanApi
	_ = json.Unmarshal(bodyText, &result) //byte to json
	total := result.Total
	fmt.Printf("node %s total: %d", addr, total)
	return int(total)
}

// Spider 传入页数，一页一页爬取
func Spider(addr string, page int, pageSize int) (int, []models.TronScan) {
	var tmp []models.TronScan
	p := strconv.Itoa(page)
	log.Printf("当前页page:%d, p:%s\n", page, p)
	client := &http.Client{}
	url := fmt.Sprintf("https://apilist.tronscanapi.com/api/filter/trc20/transfers?limit=%d&start=%d&sort=-timestamp&count=true&filterTokenValue=0&relatedAddress=%s", pageSize, page*pageSize, addr)
	println("url:" + url)
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
	var result TronScanApi
	_ = json.Unmarshal(bodyText, &result) //byte to json
	num := len(result.Transfers)

	transfers := result.Transfers
	total := result.Total
	for i := num - 1; i >= 0; i-- {
		txId := transfers[i].TxId

		if _, ok := m[txId]; ok {
			continue
		} else {
			m[txId] = txId
		}

		if transfers[i].Contract != "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t" {
			continue
		}

		var tron models.TronScan
		tron.Height = transfers[i].Height
		tron.TxId = transfers[i].TxId
		tron.Date = httputils.TimestampToTime(transfers[i].Timestamp / 1000)
		tron.FromAddr = transfers[i].From
		tron.ToAddr = transfers[i].To
		tron.FromTag = transfers[i].FromTag.Tag
		tron.ToTag = transfers[i].ToTag.Tag
		tron.Quantity = httputils.DecimalDivValue(transfers[i].Value, 6)
		tron.Contract = transfers[i].Contract
		tron.EventType = transfers[i].EventType
		tron.CreateTime = time.Now()
		tron.ContractType = transfers[i].ContractType
		tron.Confirmed = transfers[i].Confirmed

		if tron.Date.Before(httputils.StringToTime(afterTime)) {
			continue
		}

		if tron.FromAddr == addr && !tron.Quantity.IsZero() {
			tmp = append(tmp, tron)
		}

		if tron.Date.After(httputils.StringToTime(beforeTime)) {
			return -1, tmp
		}
	}

	return int(total), tmp
}

func pageCount(total int, pageSize int) int {
	page := int(math.Ceil(float64(total) / float64(pageSize)))
	return page
}
