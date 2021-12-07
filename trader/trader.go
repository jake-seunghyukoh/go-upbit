package trader

import (
	"fmt"
	"time"

	"github.com/ohshyuk5/go-upbit/quotation"
)

var marketId string

func getCurrentPrice(marketId string) quotation.TradeTick {
	tradeTick := quotation.GetTradeTicks(marketId)[0]

	return tradeTick
}

func goodToBuy(marketId string) {
	candles := quotation.GetMinuteCandles(marketId, 5, 12)
	for _, candle := range candles {
		fmt.Println(candle)
	}
}

func initialize() {
}

func Start(aMarketId string) {
	initialize()

	marketId = aMarketId

	for {
		time.Sleep(time.Millisecond * 100)
		tradeTick := getCurrentPrice(marketId)
		fmt.Println(tradeTick)
	}
}
