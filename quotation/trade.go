package quotation

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ohshyuk5/go-upbit/myHttp"
)

var history TradeHistory = TradeHistory{}

type TradeTick struct {
	TradeDateUTC     string  `json:"trade_date_utc"`
	TradeTimeUTC     string  `json:"trade_time_utc"`
	Timestamp        int64   `json:"timestamp"`
	TradePrice       float64 `json:"trade_price"`
	TradeVolume      float64 `json:"trade_volume"`
	PrevClosingPrice float64 `json:"prev_closing_price"`
	ChangePrice      float64 `json:"change_price"`
	AskBid           string  `json:"ask_bid"`
	SequentialId     int64   `json:"sequential_id"`
}

func (t TradeTick) String() string {
	return fmt.Sprintf("%s at %sT%s price: %f volume: %f", t.AskBid, t.TradeDateUTC, t.TradeTimeUTC, t.TradePrice, t.TradeVolume)
}

// Get recent trade history
func GetTradeTicks(marketId string) []TradeTick {
	queryString := fmt.Sprintf("market=%s&count=1", marketId)

	body := myHttp.Request("GET", TradeTicksURL, "", queryString, "")

	var trades []TradeTick

	err := json.Unmarshal([]byte(body), &trades)
	if err != nil {
		log.Fatal(err)
	}

	return trades
}

type TradeHistory struct {
	trades []TradeTick
}

func GetHistory() *TradeHistory {
	return &history
}

func (h *TradeHistory) Append(tradeTick TradeTick) {
	h.trades = append(h.trades, tradeTick)
}
