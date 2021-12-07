package quotation

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ohshyuk5/go-upbit/myHttp"
)

type candle struct {
	Market               string  `json:"market"`
	CandleDateTimeUTC    string  `json:"candle_date_time_utc"`
	CandleDateTimeKST    string  `json:"candle_date_time_kst"`
	OpeningPrice         float64 `json:"opening_price"`
	HighPrice            float64 `json:"high_price"`
	LowPrice             float64 `json:"low_price"`
	TradePrice           float64 `json:"trade_price"`
	Timestamp            int64   `json:"timestamp"`
	CandleAccTradePrice  float64 `json:"candle_acc_trade_price"`
	CandleAccTradeVolume float64 `json:"candle_acc_trade_volume"`
	Unit                 int     `json:"unit"`
}

func (c candle) String() string {
	return fmt.Sprintf(
		"Timestamp: %d, Opening: %.1f, Trade: %.1f Diff: %.1f, Volume: %f",
		c.Timestamp,
		c.OpeningPrice,
		c.TradePrice,
		c.TradePrice-c.OpeningPrice,
		c.CandleAccTradeVolume,
	)
}

func GetMinuteCandles(marketId string, unit, count int) []candle {
	queryString := fmt.Sprintf("market=%s&count=%d", marketId, count)
	body := myHttp.Request(
		"GET",
		fmt.Sprintf("%s/%d", CandleMinuteURL, unit),
		"",
		queryString,
		"",
	)

	var candles []candle
	err := json.Unmarshal([]byte(body), &candles)
	if err != nil {
		log.Fatal(err)
	}

	return candles
}
