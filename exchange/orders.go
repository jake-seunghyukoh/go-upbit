package exchange

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ohshyuk5/go-upbit/auth"
	"github.com/ohshyuk5/go-upbit/myHttp"
)

type constraint struct {
	Currency  string `json:"currency"`
	PriceUnit string `json:"price_unit"`
	MinTotal  string `json:"min_total"`
}

type market struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	OrderTypes []string   `json:"order_types"`
	OrderSides []string   `json:"order_sides"`
	Bid        constraint `json:"bid"`
	Ask        constraint `json:"ask"`
	MaxTotal   string     `json:"max_total"`
	State      string     `json:"state"`
}

type chance struct {
	BidFee     string  `json:"bid_fee"`
	AskFee     string  `json:"ask_fee"`
	Market     market  `json:"market"`
	BidAccount account `json:"bid_account"`
	AskAccount account `json:"ask_account"`
}

func (c chance) String() string {
	doc, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(doc)
}

var baseURL string = "https://api.upbit.com/v1/orders/chance"

func GetChance(marketId string) chance {
	accessKey, secretKey := auth.LoadKeys()

	queryString := fmt.Sprintf("market=%s", marketId)
	queryHash := auth.HashQuery(queryString)

	token := auth.CreateJwtToken(accessKey, secretKey, queryHash)

	body := myHttp.Request("GET", baseURL, token, queryString, queryString)

	var chance chance
	err := json.Unmarshal([]byte(body), &chance)
	if err != nil {
		log.Fatalf("Get Chance -> %s", err)
	}
	return chance
}
