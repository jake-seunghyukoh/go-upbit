package auth

import (
	"log"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func createPayload(accessKey string, queryHash string) jwt.Claims {
	if queryHash != "" {
		return jwt.MapClaims{"access_key": accessKey, "nonce": uuid.NewString(), "query_hash": queryHash}
	}
	return jwt.MapClaims{"access_key": accessKey, "nonce": uuid.NewString()}
}

func signPayload(secretKey string, payload jwt.Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatal("signPayload: ", err)
	}
	return signedToken
}

func CreateJwtToken(accessKey string, secretKey string, queryHash string) string {
	payload := createPayload(accessKey, queryHash)
	token := signPayload(secretKey, payload)

	return token
}
