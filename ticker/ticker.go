package ticker

import (
	"blog/httputils"
	"blog/redis"
	"encoding/json"
	"log"
	"time"
)

type Ticker struct {
	Data struct {
		NewlyPrice    float64 `json:"newlyPrice"`
		PercentChange float64 `json:"percentChange"`
		FlowTotal     float64 `json:"flowTotal"`
	} `json:"data"`
}

// GetTicker 獲取節點的地址信息
func GetTicker() Ticker {
	url := `https://api.filutils.com/api/v2/network/filprice`
	bodyText, err := httputils.RequestUrl(url)
	if err != nil {
		log.Fatal(err)
	}
	var result Ticker
	_ = json.Unmarshal(bodyText, &result) //byte to json

	return result
}

func SetTickerToRedis() {
	ticker := GetTicker()

	redis.SetRedis("ticker", ticker.Data.NewlyPrice, time.Hour)
}
