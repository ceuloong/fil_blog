package httpcorrect

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// AccountBalance 返回的结构体
type AccountBalance []struct {
	Height              uint   `json:"height"`
	Timestamp           int64  `json:"timestamp"`
	Balance             string `json:"balance"`
	AvailableBalance    string `json:"availableBalance"`
	SectorPledgeBalance string `json:"sectorPledgeBalance"`
	VestingFunds        string `json:"vestingFunds"`
}

// MessagesResult 返回的结构体
type MessagesResult struct {
	TotalCount int64 `json:"totalCount"`
	Messages   []struct {
		Cid       string `json:"cid"`
		Height    uint   `json:"height"`
		Timestamp int64  `json:"timestamp"`
		From      string `json:"from"`
		To        string `json:"to"`
		Nonce     int    `json:"nonce"`
		Value     string `json:"value"`
		Method    string `json:"method"`
	} `json:"messages"`
	Methods []string `json:"methods"`
}

type Account struct {
	Height              uint
	LastTime            int64
	Balance             decimal.Decimal
	AvailableBalance    decimal.Decimal
	SectorPledgeBalance decimal.Decimal
	VestingFunds        decimal.Decimal
	ControlAddress      string
	ControlBalance      decimal.Decimal
}

// BalanceStats 获取账户余额数组
func BalanceStats(node string) Account {
	client := &http.Client{}
	url := fmt.Sprintf("https://filfox.info/api/v1/address/%s/balance-stats", node)
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
	var result AccountBalance
	_ = json.Unmarshal(bodyText, &result) //byte to json
	length := len(result)
	if length == 0 {
		fmt.Printf("Node:%s 查询余额为0\n", node)
		account := Account{
			Height:           0,
			LastTime:         time.Now().Unix(),
			Balance:          DecimalValue("0"),
			AvailableBalance: DecimalValue("0"),
		}
		return account
	}

	accountBalance := result[length-1]

	var account Account
	account.Height = accountBalance.Height
	account.LastTime = accountBalance.Timestamp
	account.Balance = DecimalDiv18Value(accountBalance.Balance)
	account.AvailableBalance = DecimalDiv18Value(accountBalance.AvailableBalance)
	account.SectorPledgeBalance = DecimalDiv18Value(accountBalance.SectorPledgeBalance)
	account.VestingFunds = DecimalDiv18Value(accountBalance.VestingFunds)

	return account
}

// BalanceControl 获取账户余额数组
func BalanceControl(addr string) Account {
	client := &http.Client{}
	url := fmt.Sprintf("https://filfox.info/api/v1/address/%s/balance-stats?duration=24h&samples=1", addr)
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
	var result AccountBalance
	_ = json.Unmarshal(bodyText, &result) //byte to json
	length := len(result)
	var account Account
	if length > 0 {
		accountBalance := result[length-1]

		account.ControlBalance = DecimalDiv18Value(accountBalance.Balance)
	}
	return account
}

// SpiderMessage 获取第一页第一条数据
func SpiderMessage(node string) string {
	client := &http.Client{}
	url := fmt.Sprintf("https://filfox.info/api/v1/address/%s/messages?pageSize=20&page=0&method=SubmitWindowedPoSt", node)
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
	var result MessagesResult
	_ = json.Unmarshal(bodyText, &result) //byte to json
	num := len(result.Messages)

	if num == 0 {
		return ""
	}
	message := result.Messages[0]

	return message.From
}
