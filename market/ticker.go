package market

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ohshyuk5/go-upbit/myHttp"
)

type priceTick struct {
	Market             string  `json:"market"`
	TradeDate          string  `json:"trade_date"`
	TradeTime          string  `json:"trade_time"`
	TradeDateKST       string  `json:"trade_date_kst"`
	TradeTimeKST       string  `json:"trade_time_kst"`
	TradeTimestamp     float64 `json:"trade_timestamp"`
	OpeningPrice       float64 `json:"opening_price"`
	HighPrice          float64 `json:"high_price"`
	LowPrice           float64 `json:"low_price"`
	TradePrice         float64 `json:"trade_price"`
	PrevClosingPrice   float64 `json:"prev_closing_price"`
	Change             string  `json:"change"`
	ChangePrice        float64 `json:"change_price"`
	ChangeRate         float64 `json:"change_rate"`
	SignedChangePrice  float64 `json:"signed_change_price"`
	SignedChangeRate   float64 `json:"signed_change_rate"`
	TradeVolume        float64 `json:"trade_volume"`
	AccTradePrice      float64 `json:"acc_trade_price"`
	AccTradePrice24h   float64 `json:"acc_trade_price_24h"`
	AccTradeVolume     float64 `json:"acc_trade_volume"`
	AccTradeVolume24h  float64 `json:"acc_trade_volume_24h"`
	Highest52WeekPrice float64 `json:"highest_52_week_price"`
	Highest52WeekDate  string  `json:"highest_52_week_date"`
	Lowest52WeekPrice  float64 `json:"lowest_52_week_price"`
	Lowest52WeekDate   string  `json:"lowest_52_week_date"`
	Timestamp          int64   `json:"timestamp"`
}

func (p priceTick) String() string {
	doc, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	return string(doc)
}

func GetTicker(marketId string) []priceTick {
	body := myHttp.Request(
		"GET",
		TickerURL,
		"",
		fmt.Sprintf("markets=%s", marketId),
		"",
	)

	var priceTicks []priceTick
	err := json.Unmarshal([]byte(body), &priceTicks)
	if err != nil {
		log.Fatal(err)
	}
	return priceTicks
}
