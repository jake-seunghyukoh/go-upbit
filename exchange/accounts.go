package exchange

import (
	"encoding/json"
	"log"

	myhttp "github.com/ohshyuk5/go-upbit/myHttp"
)

type account struct {
	Currency            string `json:"currency"`
	Balance             string `json:"balance"`
	Locked              string `json:"locked"`
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
	UnitCurrency        string `json:"unit_currency"`
}

var url string = "https://api.upbit.com/v1/accounts"

func GetAccounts(token string) []account {
	body := myhttp.Request("GET", url, token, "", "")

	var accounts []account
	err := json.Unmarshal([]byte(body), &accounts)
	if err != nil {
		log.Fatalf("Get Accounts -> %s", err)
	}

	return accounts
}
