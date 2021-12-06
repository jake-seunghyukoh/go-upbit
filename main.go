package main

import (
	"fmt"
	"log"
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
	fmt.Println(token)

}
