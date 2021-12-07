package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type errorMessage struct {
	Name    int    `json:"name"`
	Message string `json:"message"`
}

type errorBody struct {
	Error errorMessage `json:"error"`
}

func HandleHttpError(response *http.Response, err error) {
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode >= 400 {
		defer response.Body.Close()

		var errorBody errorBody
		errDecode := json.NewDecoder(response.Body).Decode(&errorBody)
		if errDecode != nil {
			log.Fatal("Http request decoding error")
		}

		log.Fatal(fmt.Sprintf("Http request error with code:%d\n%v", response.StatusCode, errorBody))
	}
}
