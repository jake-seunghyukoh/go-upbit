package exchange

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ohshyuk5/go-upbit/utils"
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
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := http.DefaultClient.Do(req)
	utils.HandleHttpError(res, err)

	defer res.Body.Close()

	var bodies []account
	err = json.NewDecoder(res.Body).Decode(&bodies)
	if err != nil {
		log.Fatalf("Get Accounts -> %s", err)
	}

	return bodies
}
