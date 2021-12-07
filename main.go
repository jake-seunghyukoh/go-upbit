package main

import (
	"fmt"

	"github.com/ohshyuk5/go-upbit/trader"
)

func main() {
	fmt.Println("Running a Trader")
	trader.Start("KRW-BTC")
}
