package auth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var accessKey string = ""
var secretKey string = ""

func LoadKeys() (string, string) {
	if accessKey == "" || secretKey == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error while loading .env file")
		}
		accessKey = os.Getenv("ACCESS_KEY")
		secretKey = os.Getenv("SECRET_KEY")
	}
	return accessKey, secretKey
}
