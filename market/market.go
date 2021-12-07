package market

import (
	"encoding/json"
	"log"

	"github.com/ohshyuk5/go-upbit/myHttp"
)

type market struct {
	Market      string `json:"market"`
	KoreanName  string `json:"korean_name"`
	EnglishName string `json:"english_name"`
}

func (m market) String() string {
	doc, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return string(doc)
}

func GetAllMarkets() []market {
	body := myHttp.Request("GET", MarketURL, "", "", "")

	var markets []market
	err := json.Unmarshal([]byte(body), &markets)
	if err != nil {
		log.Fatalf("Get All Markets -> %s", err)
	}

	return markets
}
