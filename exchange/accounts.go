package exchange

import (
	"encoding/json"
	"log"

	"github.com/ohshyuk5/go-upbit/auth"
	"github.com/ohshyuk5/go-upbit/myHttp"
)

type account struct {
	Currency            string `json:"currency"`
	Balance             string `json:"balance"`
	Locked              string `json:"locked"`
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
	UnitCurrency        string `json:"unit_currency"`
}

func GetAccounts() []account {
	accessKey, secretKey := auth.LoadKeys()

	token := auth.CreateJwtToken(accessKey, secretKey, "")

	body := myHttp.Request("GET", AccountsURL, token, "", "")

	var accounts []account
	err := json.Unmarshal([]byte(body), &accounts)
	if err != nil {
		log.Fatalf("Get Accounts -> %s", err)
	}

	return accounts
}
