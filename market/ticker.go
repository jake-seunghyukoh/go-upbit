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
	TradeTimestamp     float32 `json:"trade_timestamp"`
	OpeningPrice       float32 `json:"opening_price"`
	HighPrice          float32 `json:"high_price"`
	LowPrice           float32 `json:"low_price"`
	TradePrice         float32 `json:"trade_price"`
	PrevClosingPrice   float32 `json:"prev_closing_price"`
	Change             string  `json:"change"`
	ChangePrice        float32 `json:"change_price"`
	ChangeRate         float32 `json:"change_rate"`
	SignedChangePrice  float32 `json:"signed_change_price"`
	SignedChangeRate   float32 `json:"signed_change_rate"`
	TradeVolume        float32 `json:"trade_volume"`
	AccTradePrice      float32 `json:"acc_trade_price"`
	AccTradePrice24h   float32 `json:"acc_trade_price_24h"`
	AccTradeVolume     float32 `json:"acc_trade_volume"`
	AccTradeVolume24h  float32 `json:"acc_trade_volume_24h"`
	Highest52WeekPrice float32 `json:"highest_52_week_price"`
	Highest52WeekDate  string  `json:"highest_52_week_date"`
	Lowest52WeekPrice  float32 `json:"lowest_52_week_price"`
	Lowest52WeekDate   string  `json:"lowest_52_week_date"`
	Timestamp          float32 `json:"timestamp"`
}

func (p priceTick) String() string {
	doc, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	return string(doc)
}

func GetTicker(marketId string) []priceTick {
	body := myHttp.Request("GET", TickerURL, "", fmt.Sprintf("markets=%s", marketId), "")

	var priceTicks []priceTick
	err := json.Unmarshal([]byte(body), &priceTicks)
	if err != nil {
		log.Fatal(err)
	}
	return priceTicks
}
