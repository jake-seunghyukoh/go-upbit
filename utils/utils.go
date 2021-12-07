package utils

import (
	"fmt"
	"log"
	"net/http"
)

func HandleHttpError(response *http.Response, err error) {
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode >= 400 {
		log.Fatal(fmt.Sprintf("Http request error with code:%d", response.StatusCode))
	}
}
