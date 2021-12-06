package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/ohshyuk5/go-upbit/myJwt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file")
	}
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")

	token := myJwt.CreateJwtToken(accessKey, secretKey, "")

	url := "https://api.upbit.com/v1/accounts"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var data string

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	fmt.Println(data)

}
