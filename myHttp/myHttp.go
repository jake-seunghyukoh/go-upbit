package myhttp

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/ohshyuk5/go-upbit/utils"
)

func Request(method string, url string, token string, queryString string, body string) string {
	var req *http.Request
	var err error

	url = url + queryString

	if body != "" {
		payload := strings.NewReader(body)
		req, err = http.NewRequest(method, url, payload)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := http.DefaultClient.Do(req)
	utils.HandleHttpError(res, err)

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(resBody)
}
