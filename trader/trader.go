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
