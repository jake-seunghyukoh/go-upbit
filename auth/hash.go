package auth

import (
	"crypto/sha512"
	"fmt"
)

func HashQuery(queryString string) string {
	sha := sha512.New()
	sha.Write([]byte(queryString))
	hash := sha.Sum(nil)
	return fmt.Sprintf("%x", hash)
}
