package main

import (
	"fmt"

	"github.com/ohshyuk5/go-upbit/exchange"
	"github.com/ohshyuk5/go-upbit/market"
)

func main() {
	fmt.Println("Requesting accounts")
	accounts := exchange.GetAccounts()

	for _, account := range accounts {
		fmt.Println(account)
	}

	fmt.Println("")
	fmt.Println("Requesting order chances")
	chance := exchange.GetChance("KRW-BTC")

	fmt.Println(chance)

	fmt.Println("")
	fmt.Println("Requesting Market List")
	markets := market.GetAllMarkets()

	fmt.Println(markets[0], "...")

	fmt.Println("")
	fmt.Println("Requesting BTC ticker")
	tickers := market.GetTicker("KRW-BTC")

	fmt.Println(tickers)
}
