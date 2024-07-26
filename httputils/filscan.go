package httputils

import (
	"blog/models"
	"blog/services"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

// ResultTotal 返回的结构体
type ResultTotal struct {
	Result struct {
		AccountInfo struct {
			AccountMiner struct {
				WorkerAddress      string   `json:"worker_address"`
				ControllersAddress []string `json:"controllers_address"`
			} `json:"account_miner"`
			AccountBasic struct {
				AccountAddress   string `json:"account_address"`
				AccountBalance   string `json:"account_balance"`
				AccountId        string `json:"account_id"`
				AccountType      string `json:"account_type"`
				CreateTime       int64  `json:"create_time"`
				LastTransferTime int64  `json:"latest_transfer_time"`
				Nonce            int64  `json:"nonce"`
			} `json:"account_basic"`
		} `json:"account_info"`
	} `json:"result"`
}

// AccountAddress 返回的结构体
type AccountAddress struct {
	WorkerAddress      string   `json:"worker_address"`
	ControllersAddress []string `json:"controllers_address"`
}

type AccountBasic struct {
	AccountAddress   string          `json:"account_address"`
	AccountBalance   decimal.Decimal `json:"account_balance"`
	AccountId        string          `json:"account_id"`
	AccountType      string          `json:"account_type"`
	CreateTime       time.Time       `json:"create_time"`
	LastTransferTime time.Time       `json:"last_transfer_time"`
	Nonce            int64           `json:"nonce"`
}

// BalanceControlById 获取账户余额数组
func BalanceControlById(addr string) AccountBasic {
	client := &http.Client{}
	url := "https://api-v2.filscan.io/api/v1/AccountInfoByID"
	jsonStr := []byte(`{ "account_id": "` + addr + `"}`)
	reqSpider, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	reqSpider.Header.Set("Content-Type", "application/json;charset=UTF-8")
	reqSpider.Header.Set("Locale", "zh")
	reqSpider.Header.Set("sec-ch-ua", "\"Chromium\";v=\"112\", \"Google Chrome\";v=\"112\", \"Not:A-Brand\";v=\"99\"")
	reqSpider.Header.Set("sec-ch-ua-mobile", "?0")
	reqSpider.Header.Set("sec-ch-ua-platform", "macOS")
	reqSpider.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	respSpider, err := client.Do(reqSpider)
	if respSpider.StatusCode != 200 {
		log.Fatal(respSpider.Status)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer respSpider.Body.Close()
	bodyText, err := ioutil.ReadAll(respSpider.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("bodyText: %s\n", bodyText)
	var result ResultTotal
	_ = json.Unmarshal(bodyText, &result) //byte to json

	accountBasic := result.Result.AccountInfo.AccountBasic
	basic := AccountBasic{
		AccountBalance:   DecimalDiv18Value(accountBasic.AccountBalance),
		AccountId:        accountBasic.AccountId,
		AccountType:      accountBasic.AccountType,
		CreateTime:       TimestampToTime(accountBasic.CreateTime),
		LastTransferTime: TimestampToTime(accountBasic.LastTransferTime),
		Nonce:            accountBasic.Nonce,
	}

	return basic
}

// AccountInfoById 獲取節點的地址信息
func AccountInfoById(node string) AccountAddress {
	client := &http.Client{}
	url := "https://api-v2.filscan.io/api/v1/AccountInfoByID"
	jsonStr := []byte(`{ "account_id": "` + node + `"}`)
	reqSpider, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	reqSpider.Header.Set("Content-Type", "application/json;charset=UTF-8")
	//reqSpider.Header.Set("Accept", "application/json, text/plain, */*")
	reqSpider.Header.Set("Locale", "zh")
	reqSpider.Header.Set("sec-ch-ua", "\"Chromium\";v=\"112\", \"Google Chrome\";v=\"112\", \"Not:A-Brand\";v=\"99\"")
	reqSpider.Header.Set("sec-ch-ua-mobile", "?0")
	reqSpider.Header.Set("sec-ch-ua-platform", "macOS")
	reqSpider.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	respSpider, err := client.Do(reqSpider)
	if respSpider.StatusCode == 406 {
		log.Fatal(respSpider.Status)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer respSpider.Body.Close()
	bodyText, err := ioutil.ReadAll(respSpider.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result ResultTotal

	//fmt.Printf("bodyText: %s\n", bodyText)
	_ = json.Unmarshal(bodyText, &result) //byte to json

	accountAddress := AccountAddress{
		WorkerAddress:      result.Result.AccountInfo.AccountMiner.WorkerAddress,
		ControllersAddress: result.Result.AccountInfo.AccountMiner.ControllersAddress,
	}
	return accountAddress
}

func UpdateAddresses(msig string) {
	nodes := services.FindAllNode(msig)

	for _, oneNode := range nodes {
		node := oneNode.Node
		// 获取节点的控制地址 独立方法获取 独立表保存，一个节点存在多个控制地址
		if len(oneNode.ControlAddress) == 0 {
			var filAddresses []models.FilAddresses
			log.Printf("获取节点%s地址\n", node)
			account := AccountInfoById(node)
			oneNode.ControlAddress = account.WorkerAddress
			time.Sleep(5 * time.Second)

			address := models.FilAddresses{
				Node:             node,
				Address:          account.WorkerAddress,
				Type:             "worker",
				CreatedTime:      time.Now(),
				CreateTime:       time.Now(),
				LastTransferTime: time.Now(),
			}
			filAddresses = append(filAddresses, address)

			for _, controllersAddress := range account.ControllersAddress {
				address = models.FilAddresses{
					Node:             node,
					Address:          controllersAddress,
					Type:             "controller",
					CreatedTime:      time.Now(),
					CreateTime:       time.Now(),
					LastTransferTime: time.Now(),
				}

				filAddresses = append(filAddresses, address)
			}

			if len(filAddresses) > 0 {
				services.InsertAddress(filAddresses)

				if oneNode.CreateTime.IsZero() {
					oneNode.CreateTime = time.Now()
					oneNode.LastTime = time.Now()
				}
				services.UpdateNode(oneNode)
			}
		}
	}

}

func UpdateAddressesBalance(timeTag int64, addrParam string) {
	addresses := services.FindAllAddress(addrParam)
	s := services.LuckyBlock{}

	var nodeM = make(map[string]decimal.Decimal)
	for _, addr := range addresses {
		//if addr.LastTransferTime.Compare(GetLastMonthTime()) < 0 {
		//	continue
		//}

		realTimeTag := timeTag
		var count int64
		if addr.TransferCount == 0 {
			s.CountByNode(addr.AccountId, &count)
			if count > 0 {
				realTimeTag = 0
			}
		}

		var needRequestUrl bool = false
		// 获取地址的转入转出销毁数量
		mapA := services.SumValueByType(addr.AccountId, realTimeTag)
		//addr.ReceiveAmount = mapA["receive"]
		//addr.BurnAmount = mapA["burn-fee"]
		//addr.SendAmount = mapA["send"]

		var burnNew decimal.Decimal
		if value, ok := mapA["receive"]; ok {
			addr.ReceiveAmount = addr.ReceiveAmount.Add(value)
			needRequestUrl = true
		}
		if value, ok := mapA["burn-fee"]; ok {
			addr.BurnAmount = addr.BurnAmount.Add(value)
			burnNew = burnNew.Add(value)
			needRequestUrl = true
		}
		if value, ok := mapA["miner-fee"]; ok {
			addr.BurnAmount = addr.BurnAmount.Add(value)
			burnNew = burnNew.Add(value)
			needRequestUrl = true
		}
		if value, ok := mapA["send"]; ok {
			addr.SendAmount = addr.SendAmount.Add(value)
			needRequestUrl = true
		}

		if value, ok := nodeM[addr.Node]; ok {
			nodeM[addr.Node] = value.Add(burnNew)
		} else {
			nodeM[addr.Node] = burnNew
		}
		if needRequestUrl {
			account := BalanceControlById(addr.Address)
			time.Sleep(5 * time.Second)

			addr.Balance = account.AccountBalance
			if len(addr.AccountId) == 0 {
				addr.AccountId = account.AccountId
			}
			addr.CreateTime = account.CreateTime
			addr.AccountType = account.AccountType
			addr.LastTransferTime = account.LastTransferTime
			addr.Nonce = account.Nonce

			//s.CountByNodeTimeTag(addr.AccountId, timeTag, &count)
			addr.TransferCount = addr.RealCount

			addr.TimeTag = timeTag

			services.UpdateBalance(addr)
		}
	}

	if len(nodeM) > 0 {
		nodes := services.FindAllNode("")
		for _, n := range nodes {
			if value, ok := nodeM[n.Node]; ok {
				n.BurnAmount = n.BurnAmount.Add(value)
				services.UpdateNode(n)
			}
		}
	}

}
