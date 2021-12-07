package main

import (
	"fmt"

	"github.com/ohshyuk5/go-upbit/exchange"
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
}
