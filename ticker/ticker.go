package ticker

import (
	"blog/httputils"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"log"
)

type Ticker struct {
	Symbol string          `json:"symbol"`
	Price  decimal.Decimal `json:"price"`
}

// GetTicker 獲取節點的地址信息
func GetTicker() Ticker {
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", "FILUSDT")
	bodyText, err := httputils.RequestUrl(url)
	if err != nil {
		log.Fatal(err)
	}
	var result Ticker
	_ = json.Unmarshal(bodyText, &result) //byte to json

	return result
}
